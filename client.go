package gomailman

import (
	"bytes"
	"encoding/json"
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

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	domain := new(Domain)
	err = json.Unmarshal(b, domain)
	if err != nil {
		return nil, err
	}

	err = res.Body.Close()
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (c *Client) AddDomain(domain *Domain) error {
	b, err := json.Marshal(domain)
	if err != nil {
		return err
	}

	_, err = c.conn.do(http.MethodPost, c.buildURL("domains"), bytes.NewReader(b))
	if err != nil {
		return err
	}

	return nil
}
