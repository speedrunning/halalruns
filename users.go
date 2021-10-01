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
	Max uint

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
	Top uint

	/* Return only PBs from games or romhacks in the specified series. This can be either the ID
	 * of a series, or the series abbreviation.
	 */
	Series string

	/* Return only PBs from the specified game. This can be either the ID of a game, or the
	 * games abbreviation.
	 */
	Game string
}

func paginateUsers(endpoint string, max uint) ([]User, error) {
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
	var i uint64
	var err error
	var mx sync.Mutex
	var wg sync.WaitGroup
	allData := []User{}
	count := uint64(math.Ceil(float64(max) / float64(maxRequestSize)))

	for i = 0; i < count; i++ {
		wg.Add(1)

		/* `i` is passed as a parameter to avoid a potential race condition */
		go func(n uint64) {
			defer wg.Done()
			n *= maxRequestSize

			var o struct {
				Data       []User `json:"data"`
				Pagination struct {
					Size int `json:"size"`
				} `json:"pagination"`
			}

			localErr := requestAndUnmarshall(endpoint+"&offset="+
				strconv.FormatUint(n, 10), &o)
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
	if uint(len(allData)) <= max {
		return allData, nil
	}
	return allData[:max], nil
}

/* FetchUser returns a single `User` struct based off of the provided `UserFilter`. This function
 * simply calls `halalruns.FetchUsers(uf)` and extracts only the first user.
 */
func FetchUser(uf UserFilter) (User, error) {
	users, err := FetchUsers(uf)
	if err != nil {
		return User{}, err
	}
	return users[0], nil
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
		endpoint += "&max=" + strconv.FormatUint(uint64(uf.Max), 10)
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
