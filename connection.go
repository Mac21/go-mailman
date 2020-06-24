package gomailman

import (
	"errors"
	"io"
	"net/http"
	"net/url"
)

type Connection struct {
	conn     http.Client
	baseURL  string
	username string
	password string
}

func (c *Connection) do(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, c.baseURL+url, body)
	if err != nil {
		return nil, err
	}

	if req.ContentLength > 0 {
		req.Header.Add("Content-Type", "application/json")
	}

	req.SetBasicAuth(c.username, c.password)
	req.Close = true
	return c.conn.Do(req)
}

func NewConnection(baseURL, username, password string) (*Connection, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	if username == "" || password == "" {
		return nil, errors.New("username and password required to connect to mailman")
	}

	return &Connection{
		baseURL:  baseURL,
		username: username,
		password: password,
	}, nil
}
