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

import "time"

/* A link to another page */
type Link struct {
	/* The name of the page */
	Rel string `json:"rel"`
	/* The URI to the page */
	URI string `json:"uri"`
}

/* User struct representing a speedrun.com user */
type User struct {
	/* The users user ID */
	ID string `json:"id"`
	/* The users username */
	Names struct {
		/* The users international name, this is what you see on the site */
		International string `json:"international"`
		/* The users japanese name, this is deprecated */
		Japanese string `json:"japanese"`
	} `json:"names"`
	/* The users pronouns */
	Pronouns string `json:"pronouns"`
	/* A link to the users profile */
	Weblink string `json:"weblink"`
	/* The style and colors of the users name */
	NameStyle struct {
		/* The style type of the users name */
		Style string `json:"style"`
		/* The starting gradient color */
		ColorFrom struct {
			/* The light version of the color */
			Light string `json:"light"`
			/* The dark version of the color */
			Dark string `json:"dark"`
		} `json:"color-from"`
		/* The ending gradient color */
		ColorTo struct {
			/* The light version of the color */
			Light string `json:"light"`
			/* The dark version of the color */
			Dark string `json:"dark"`
		} `json:"color-to"`
	} `json:"name-style"`
	Role string `json:"role"`
	/* The time the user signed up to the site */
	Signup time.Time `json:"signup"`
	/* The users location */
	Location struct {
		/* The users country */
		Country struct {
			/* The countries ISO 3166-1 alpha-2 country code */
			Code  string `json:"code"`
			Names struct {
				/* The countries name */
				International string `json:"international"`
				/* The countries name in Japanese, this is deprecated. */
				Japanese string `json:"japanese"`
			} `json:"names"`
		} `json:"country"`
		/* The users region */
		Region struct {
			/* The regions code and the countries ISO 3166-1 alpha-2 country code */
			Code  string `json:"code"`
			Names struct {
				/* The regions name */
				International string `json:"international"`
				/* The regions name in Japanese, this is deprecated. */
				Japanese string `json:"japanese"`
			} `json:"names"`
		} `json:"region"`
	} `json:"location"`
	/* The users Twitch channel */
	Twitch struct {
		/* The URI */
		URI string `json:"uri"`
	} `json:"twitch"`
	/* The users Hitbox account */
	Hitbox struct {
		/* The URI */
		URI string `json:"uri"`
	} `json:"hitbox"`
	/* The users YouTube channel */
	Youtube struct {
		/* The URI */
		URI string `json:"uri"`
	} `json:"youtube"`
	/* The users Twitter account */
	Twitter struct {
		/* The URI */
		URI string `json:"uri"`
	} `json:"twitter"`
	/* The users SpeedRunsLive account */
	SpeedRunsLive struct {
		/* The URI */
		URI string `json:"uri"`
	} `json:"speedrunslive"`
	/* The users assets on their profile */
	Assets struct {
		/* The users icon next to their name */
		Icon struct {
			/* Link to the icon */
			URI string `json:"uri"`
		} `json:"icon"`
		/* The users profile picture */
		Image struct {
			/* Link to the profile picture */
			URI string `json:"uri"`
		} `json:"image"`
	} `json:"assets"`
	/* Various links related to the user */
	Links []Link `json:"links"`
}
