package gomailman

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Client struct {
	conn *Connection
}

func (c *Client) buildURL(parts ...string) string {
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

func NewClient(host, username, password string) (*Client, error) {
	conn, err := NewConnection(host, username, password)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) GetDomain(domainID string) (*Domain, error) {
	res, err := c.conn.do(http.MethodGet, c.buildURL("domains", domainID), http.NoBody)
	if err != nil {
		return nil, err
	}

	if res.StatusCode/100 != 2 {
		return nil, errors.New("Domain does not exist")
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	domain := new(Domain)
	err = json.Unmarshal(b, domain)
	if err != nil {
		return nil, err
	}

	return domain, res.Body.Close()
}

func (c *Client) AddDomain(domain *Domain) error {
	b, err := json.Marshal(domain)
	if err != nil {
		return err
	}

	res, err = c.conn.do(http.MethodPost, c.buildURL("domains"), bytes.NewReader(b))
	if err != nil {
		return err
	}

	if res.StatusCode/100 != 2 {
		return errors.New("Failed to create domain")
	}

	return res.Body.Close()
}

func (c *Client) DeleteDomain(domainID string) error {
	res, err := c.conn.do(http.MethodDelete, c.buildURL("domains", domainID), http.NoBody)
	if err != nil {
		return err
	}

	if res.StatusCode/100 != 2 {
		return errors.New("Failed to delete domain")
	}

	return res.Body.Close()
}
