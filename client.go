package gomailman

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	conn *Connection
}

func NewClient(host, username, password string) (*Client, error) {
	conn, err := NewConnection(host, username, password)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

func buildURL(parts ...string) string {
	var url string
	for indx, p := range parts {
		if indx == 0 {
			url += p
		} else {
			url += "/" + p
		}
	}

	return url
}

func parseResponseError(res *http.Response) error {
	if res.StatusCode/100 != 2 {
		re := RequestError{}
		if err := json.NewDecoder(res.Body).Decode(&re); err != nil {
			return fmt.Errorf("go-mailman: Error %d %s", res.StatusCode, err)
		}

		return re
	}

	return nil
}
