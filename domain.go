package gomailman

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ErrorDomainGet    = "Error getting domain"
	ErrorDomainDelete = "Error deleting domain"
	ErrorDomainAdd    = "Error adding domain"
)

type Domain struct {
	AliasDomain string `json:"alias_domain,omitempty"`
	Description string `json:"description,omitempty"`
	MailHost    string `json:"mail_host"`
}

func (c *Client) GetDomain(domainID string) (*Domain, error) {
	res, err := c.conn.do(http.MethodGet, c.buildURL("domains", domainID), http.NoBody)
	if err != nil {
		return nil, err
	}

	if res.StatusCode/100 != 2 {
		return nil, fmt.Errorf("%s: %s", ErrorDomainGet, res.Status)
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
	if domain == nil {
		return errors.New("Error adding nil domain")
	}

	b, err := json.Marshal(domain)
	if err != nil {
		return err
	}

	res, err := c.conn.do(http.MethodPost, c.buildURL("domains"), bytes.NewReader(b))
	if err != nil {
		return err
	}

	if res.StatusCode/100 != 2 {
		return fmt.Errorf("%s: %s", ErrorDomainAdd, res.Status)
	}

	return res.Body.Close()
}

func (c *Client) DeleteDomain(domainID string) error {
	res, err := c.conn.do(http.MethodDelete, c.buildURL("domains", domainID), http.NoBody)
	if err != nil {
		return err
	}

	if res.StatusCode/100 != 2 {
		return fmt.Errorf("%s: %s", ErrorDomainDelete, res.Status)
	}

	return res.Body.Close()
}
