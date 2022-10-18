package client

import (
	"errors"
	"net/http"
)

func DoReq() error {
	resp, err := http.Get("http://localhost:8080/ping")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("bad response")
	}

	return nil
}
