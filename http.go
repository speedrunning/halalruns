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
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	api       = "https://www.speedrun.com/api/v1"
	rateLimit = 420
	maxLoops  = 10
)

type httpError struct {
	Message string `json:"message"`
}

func requestAndUnmarshall(endpoint string, object interface{}, headers map[string]string) error {
	jsonBytes, err := request(endpoint, headers)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonBytes, object)
	if err != nil {
		return err
	}

	return nil
}

func request(endpoint string, headers map[string]string) ([]byte, error) {
	var resp *http.Response

	req, err := http.NewRequest("GET", api+endpoint, nil)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	for i := 0; true; i++ {
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode == rateLimit {
			if i == maxLoops {
				resp.Body.Close()
				return nil, errors.New("Request failed (too many rate limits)")
			}
			time.Sleep(2 * time.Second)
		} else if resp.StatusCode >= 400 {
			data, _ := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			message := httpError{}
			err = json.Unmarshal(data, &message)
			return nil, errors.New(message.Message)
		} else {
			data, _ := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			return data, nil
		}
	}

	/* NOTREACHED */
	return nil, nil
}
