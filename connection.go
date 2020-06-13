package gomailman

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

type Connection struct {
	conn     http.Client
	baseurl  string
	username string
	password string
}

func (c *Connection) do(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, c.baseurl+url, body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.username, c.password)
	return c.conn.Do(req)
}

func NewConnection(baseurl, username, password string) (*Connection, error) {
	if !strings.HasSuffix(baseurl, "/") {
		baseurl += "/"
	}

	if username == "" || password == "" {
		return nil, errors.New("username and password required to connect to mailman")
	}

	return &Connection{
		baseurl:  baseurl,
		username: username,
		password: password,
	}, nil
}
