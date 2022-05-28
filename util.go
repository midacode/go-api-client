package client

import (
	"encoding/json"
	"net/http"
)

func readJSON(r *http.Response, dst interface{}) error {
	// Decode the request body into the target destination.
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		return err
	}
	return nil
}

func ServerAuth(r *http.Request) (username string, password string, ok bool) {
	return r.BasicAuth()
}
