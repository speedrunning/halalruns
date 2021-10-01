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
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	api       = "https://www.speedrun.com/api/v1"
	rateLimit = 420
)

func requestAndUnmarshall(endpoint string, object interface{}) error {
	jsonBytes, err := request(endpoint)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonBytes, object)
	if err != nil {
		return err
	}

	return nil
}

func request(endpoint string) ([]byte, error) {
	var resp *http.Response
	var err error

	for {
		resp, err = http.Get(api + endpoint)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != rateLimit {
			break
		}
		time.Sleep(2 * time.Second)
		resp.Body.Close()
	}

	data, _ := ioutil.ReadAll(resp.Body)
	return data, nil
}
