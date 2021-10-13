/*
 * BSD Zero Clause License
 *
 * Copyright (c) 2021 Thomas Voss
 *
 * Permission to use, copy, modify, and/or distribute this software for any
 * purpose with or without fee is hereby granted.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
 * REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
 * INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
 * LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR
 * OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
 * PERFORMANCE OF THIS SOFTWARE.
 */

package halalruns

import (
	"math"
	"strconv"
	"sync"
)

const maxRequestSize = 200

/* UserOrder is an enum for the different ways to order users */
type UserOrder uint

const (
	/* IntName sorts by international name */
	IntName UserOrder = iota + 1

	/* JapName sorts by japanese name which is deprecated, so don't use this */
	JapName

	/* Signup sorts by signup date */
	Signup

	/* Role sorts by the users role */
	Role
)

/* SortDirection is an enum for the direction in which users are sorted */
type SortDirection uint

const (
	/* Asc sorts in ascending order */
	Asc UserOrder = iota + 1

	/* Desc sorts in descending order */
	Desc
)

/* UserFilter is a struct of filters that can be passed to functions that search for users to
 * retrieve only the specific users you want.
 */
type UserFilter struct {
	/* The users user ID */
	ID string

	/* A case-sensitive exact match of the users name, urls, or social profiles */
	Lookup string

	/* A case-insensitive substring of the users name or urls */
	Name string

	/* The users Twitch username */
	Twitch string

	/* The users Hitbox username */
	Hitbox string

	/* The users Twitter username */
	Twitter string

	/* The users SpeedrunsLive username */
	Speedrunslive string

	/* The maximum number of users to return, by default this is 20. */
	Max int

	/* How to order the users being searched for, this can be either by international name,
	 * japanese name, signup date, or role.
	 */
	OrderBy UserOrder

	/* The direction with which to sort the users. This can be either ascending or
	 * descending.
	 */
	Direction SortDirection
}

/* PBFilter is a filter used by the `User.PersonalBests` method to filter the returned personal
 * bests
 */
type PBFilter struct {
	/* Return only PBs whos position is greater than or equal to this value. For example,
	 * `halalruns.PBFilter{Top: 3}` will return all of a users top-3 runs.
	 */
	Top int

	/* Return only PBs from games or romhacks in the specified series. This can be either the ID
	 * of a series, or the series abbreviation.
	 */
	Series string

	/* Return only PBs from the specified game. This can be either the ID of a game, or the
	 * games abbreviation.
	 */
	Game string

	/* Embed the specified fields into each personal best. The embeds need to be specified the
	 * same way they are specified in the GET request URI. Read the official documentation for
	 * more info.
	 */
	Embeds string
}

func paginateUsers(endpoint string, max int) ([]User, error) {
	/* If we can do this all in one request, just do that */
	if max <= maxRequestSize {
		var u struct {
			Data []User `json:"data"`
		}
		err := requestAndUnmarshall(endpoint, &u)
		if err != nil {
			return nil, err
		}
		return u.Data, nil
	}

	/* If we need multiple requests, we do them concurrently */
	var err error
	var mx sync.Mutex
	var wg sync.WaitGroup
	allData := []User{}
	count := int(math.Ceil(float64(max) / float64(maxRequestSize)))

	for i := 0; i < count; i++ {
		wg.Add(1)

		/* `i` is passed as a parameter to avoid a potential race condition */
		go func(n int) {
			defer wg.Done()
			n *= maxRequestSize

			var o struct {
				Data       []User `json:"data"`
				Pagination struct {
					Size int `json:"size"`
				} `json:"pagination"`
			}

			localErr := requestAndUnmarshall(endpoint+"&offset="+strconv.Itoa(n), &o)
			if localErr != nil {
				err = localErr
			}

			mx.Lock()
			allData = append(allData, o.Data...)
			mx.Unlock()
		}(i)
	}
	wg.Wait()

	/* If any error occured in any goroutine then fail */
	if err != nil {
		return nil, err
	}

	/* Return only the amount of users requested */
	if len(allData) <= max {
		return allData, nil
	}
	return allData[:max], nil
}

/* FetchUser returns a single `User` struct by searching for a user that matches the name `name`
 * exactly.
 */
func FetchUser(name string) (User, error) {
	user, err := FetchUsers(UserFilter{Lookup: name})
	if err != nil {
		return User{}, err
	}
	return user[0], nil
}

/* FetchUsers returns a slice of `User` structs based off of the provided `UserFilter` */
func FetchUsers(uf UserFilter) ([]User, error) {
	if uf.ID != "" {
		var u struct {
			Data User `json:"data"`
		}
		endpoint := "/users/" + uf.ID

		err := requestAndUnmarshall(endpoint, &u)
		if err != nil {
			return nil, err
		}
		return []User{u.Data}, nil
	}

	endpoint := "/users?"
	if uf.Lookup != "" {
		endpoint += "&lookup=" + uf.Lookup
	}
	if uf.Name != "" {
		endpoint += "&name=" + uf.Name
	}
	if uf.Twitch != "" {
		endpoint += "&twitch=" + uf.Twitch
	}
	if uf.Hitbox != "" {
		endpoint += "&hitbox=" + uf.Hitbox
	}
	if uf.Twitter != "" {
		endpoint += "&twitter=" + uf.Twitter
	}
	if uf.Speedrunslive != "" {
		endpoint += "&speedrunslive=" + uf.Speedrunslive
	}
	if uf.Max != 0 {
		endpoint += "&max=" + strconv.Itoa(uf.Max)
	}

	switch uf.OrderBy {
	case UserOrder(IntName):
		endpoint += "&orderby=name.int"
	case UserOrder(JapName):
		endpoint += "&orderby=name.jap"
	case UserOrder(Signup):
		endpoint += "&orderby=name.signup"
	case UserOrder(Role):
		endpoint += "&orderby=name.role"
	}

	switch uf.Direction {
	case SortDirection(Asc):
		endpoint += "&direction=asc"
	case SortDirection(Desc):
		endpoint += "&direction=desc"
	}

	users, err := paginateUsers(endpoint, uf.Max)
	if err != nil {
		return nil, err
	}
	return users, nil
}

/* PersonalBests fetches all of users personal bests which can be filtered via the given `PBFilter`
 * struct. If you do not want to apply any filters you can pass `nil` in place of the filter.
 */
func (u User) PersonalBests(pbf PBFilter) ([]PersonalBest, error) {
	endpoint := "/users/" + u.ID + "/personal-bests?"

	if pbf != (PBFilter{}) {
		if pbf.Game != "" {
			endpoint += "&game=" + pbf.Game
		}
		if pbf.Series != "" {
			endpoint += "&series=" + pbf.Series
		}
		if pbf.Top != 0 {
			endpoint += "&top=" + strconv.Itoa(pbf.Top)
		}
		if pbf.Embeds != "" {
			endpoint += "&embed=" + pbf.Embeds
		}
	}

	/* speedrun.com is pretty retarded, so lets make the interface less retarded for the user */
	var runs struct {
		Data []struct {
			Place uint `json:"place"`
			Run   Run  `json:"run"`
			Game  struct {
				Data Game `json:"data"`
			} `json:"game"`
			Category struct {
				Data Category `json:"data"`
			} `json:"category"`
			Level struct {
				Data Level `json:"data"`
			} `json:"level"`
			Region struct {
				Data Region `json:"data"`
			} `json:"region"`
			Platform struct {
				Data Platform `json:"data"`
			} `json:"platform"`
		} `json:"data"`
	}
	err := requestAndUnmarshall(endpoint, &runs)
	if err != nil {
		return nil, err
	}
	finalRuns := make([]PersonalBest, len(runs.Data))
	for i, r := range runs.Data {
		finalRuns[i] = PersonalBest{
			Place:    r.Place,
			Run:      r.Run,
			Game:     r.Game.Data,
			Category: r.Category.Data,
			Level:    r.Level.Data,
			Region:   r.Region.Data,
			Platform: r.Platform.Data,
		}
	}

	return finalRuns, nil
}

/* WorldRecords fetches all of users world records. This function does not accept a `PBFilter`. If
 * you want to filter specific WRs you are better off using `User.PersonalBests()`.
 */
func (u User) WorldRecords() ([]PersonalBest, error) {
	return u.PersonalBests(PBFilter{Top: 1})
}

/* Podiums fetches all of users podium runs. A podium run is any top-3 run. This function does not
 * accept a `PBFilter`. If you want to filter specific WRs you are better off using
 * `User.PersonalBests()`.
 */
func (u User) Podiums() ([]PersonalBest, error) {
	return u.PersonalBests(PBFilter{Top: 3})
}
