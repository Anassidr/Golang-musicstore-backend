package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// x is an empty interface that will be used to hold the parsed JSON data from the request body

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
