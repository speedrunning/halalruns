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
		/* The color of the users name (only applies when Style == "solid") */
		Color struct {
			/* The light version of the color */
			Light string `json:"light"`
			/* The dark version of the color */
			Dark string `json:"dark"`
		} `json:"color"`
		/* The starting gradient color (only applies when Style == "gradient") */
		ColorFrom struct {
			/* The light version of the color */
			Light string `json:"light"`
			/* The dark version of the color */
			Dark string `json:"dark"`
		} `json:"color-from"`
		/* The ending gradient color (only applies when Style == "gradient") */
		ColorTo struct {
			/* The light version of the color */
			Light string `json:"light"`
			/* The dark version of the color */
			Dark string `json:"dark"`
		} `json:"color-to"`
	} `json:"name-style"`
	/* The users role */
	Role string `json:"role"`
	/* The time the user signed up to the site */
	Signup time.Time `json:"signup"`
	/* The users location */
	Location struct {
		/* The users country */
		Country struct {
			/* The countries ISO 3166-1 alpha-2 country code */
			Code  string `json:"code"`
			/* The countries name */
			Names struct {
				/* The countries name in english */
				International string `json:"international"`
				/* The countries name in japanese, this is deprecated. */
				Japanese string `json:"japanese"`
			} `json:"names"`
		} `json:"country"`
		/* The users region */
		Region struct {
			/* The regions code and the countries ISO 3166-1 alpha-2 country code */
			Code  string `json:"code"`
			/* The regions name */
			Names struct {
				/* The regions name in english */
				International string `json:"international"`
				/* The regions name in japanese, this is deprecated. */
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

type Run struct {
	ID       string `json:"id"`
	Weblink  string `json:"weblink"`
	Game     string `json:"game"`
	Level    string `json:"level"`
	Category string `json:"category"`
	Videos   struct {
		Text  string `json:"text"`
		Links []struct {
			URI string `json:"uri"`
		} `json:"links"`
	} `json:"videos"`
	Comment string `json:"comment"`
	Status  struct {
		Status     string    `json:"status"`
		Examiner   string    `json:"examiner"`
		VerifyDate time.Time `json:"verify-date"`
	} `json:"status"`
	Players []struct {
		Rel  string `json:"rel"`
		ID   string `json:"id"`
		Name string `json:"name"`
		URI  string `json:"uri"`
	} `json:"players"`
	Date      string    `json:"date"`
	Submitted time.Time `json:"submitted"`
	Times     struct {
		Primary          string  `json:"primary"`
		PrimaryT         float64 `json:"primary_t"`
		Realtime         string  `json:"realtime"`
		RealtimeT        float64 `json:"realtime_t"`
		RealtimeNoloads  string  `json:"realtime_noloads"`
		RealtimeNoloadsT float64 `json:"realtime_noloads_t"`
		Ingame           string  `json:"ingame"`
		IngameT          float64 `json:"ingame_t"`
	} `json:"times"`
	System struct {
		Platform string `json:"platform"`
		Emulated bool   `json:"emulated"`
		Region   string `json:"region"`
	} `json:"system"`
	Splits Link              `json:"splits"`
	Values map[string]string `json:"values"`
	Links  []Link            `json:"links"`
}

type Game struct {
	ID    string `json:"id"`
	Names struct {
		International string `json:"international"`
		Japanese      string `json:"japanese"`
		Twitch        string `json:"twitch"`
	} `json:"names"`
	Abbreviation string `json:"abbreviation"`
	Weblink      string `json:"weblink"`
	Released     uint   `json:"released"`
	ReleaseDate  string `json:"realease-date"`
	Ruleset      struct {
		ShowMilliseconds    bool     `json:"show-milliseconds"`
		RequireVerification bool     `json:"require-verification"`
		RequireVideo        bool     `json:"require-video"`
		RunTimes            []string `json:"run-times"`
		DefaultTime         string   `json:"default-time"`
		EmulatorsAllowed    bool     `json:"emulators-allowed"`
	} `json:"ruleset"`
	Romhack    bool              `json:"romhack"`
	Gametypes  []string          `json:"gametypes"`
	Platforms  []string          `json:"platforms"`
	Regions    []string          `json:"regions"`
	Genres     []string          `json:"genres"`
	Engines    []string          `json:"engines"`
	Developers []string          `json:"developers"`
	Publishers []string          `json:"publishers"`
	Moderators map[string]string `json:"moderators"`
	Created    time.Time         `json:"created"`
	Assets     struct {
		Logo struct {
			URI string `json:"uri"`
		} `json:"logo"`
		CoverTiny struct {
			URI string `json:"uri"`
		} `json:"cover-tiny"`
		CoverSmall struct {
			URI string `json:"uri"`
		} `json:"cover-small"`
		CoverMedium struct {
			URI string `json:"uri"`
		} `json:"cover-medium"`
		CoverLarge struct {
			URI string `json:"uri"`
		} `json:"cover-large"`
		Icon struct {
			URI string `json:"uri"`
		} `json:"icon"`
		Trophy1st struct {
			URI string `json:"uri"`
		} `json:"trophy-1st"`
		Trophy2nd struct {
			URI string `json:"uri"`
		} `json:"trophy-2nd"`
		Trophy3rd struct {
			URI string `json:"uri"`
		} `json:"trophy-3rd"`
		Trophy4th struct {
			URI string `json:"uri"`
		} `json:"trophy-4th"`
		Background struct {
			URI string `json:"uri"`
		} `json:"background"`
		Foreground struct {
			URI string `json:"uri"`
		} `json:"foreground"`
	} `json:"assets"`
	Links []Link `json:"links"`
}

type Category struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Weblink string `json:"weblink"`
	Type    string `json:"type"`
	Rules   string `json:"rules"`
	Players struct {
		Type  string `json:"type"`
		Value uint   `json:"value"`
	} `json:"players"`
	Miscellaneous bool   `json:"miscellaneous"`
	Links         []Link `json:"links"`
}

type Level struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Weblink string `json:"weblink"`
	Rules   string `json:"rules"`
	Links   []Link `json:"links"`
}

type Region struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Links []Link `json:"links"`
}

type Platform struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Released uint   `json:"released"`
	Links    []Link `json:"links"`
}

type PersonalBest struct {
	Place    uint
	Run      Run
	Game     Game
	Category Category
	Level    Level
	Region   Region
	Platform Platform
}
