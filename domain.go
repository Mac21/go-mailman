package gomailman

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrorDomainGet is returned in GetDomain and GetAllDomains by any non 200 response
	ErrorDomainGet = errors.New("go-mailman: Error getting domain")
	// ErrorDomainDelete is returned by DeleteDomain by any non 200 respose
	ErrorDomainDelete = errors.New("go-mailman: Error deleting domain")
	// ErrorDomainAdd is returned by AddDomain when any non 200 respose
	ErrorDomainAdd = errors.New("go-mailman: Error adding domain")
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
		re := &RequestError{}
		if err := json.NewDecoder(res.Body).Decode(re); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s %s", ErrorDomainGet, re)
	}

	domain := new(Domain)
	if err = json.NewDecoder(res.Body).Decode(domain); err != nil {
		return nil, err
	}

	return domain, res.Body.Close()
}

func (c *Client) GetAllDomains() ([]*Domain, error) {
	res, err := c.conn.do(http.MethodGet, c.buildURL("domains"), http.NoBody)
	if err != nil {
		return nil, err
	}

	if res.StatusCode/100 != 2 {
		re := &RequestError{}
		if err := json.NewDecoder(res.Body).Decode(re); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s %s", ErrorDomainGet, re)
	}

	pr := &PagedResult{}
	if err = json.NewDecoder(res.Body).Decode(pr); err != nil {
		return nil, err
	}

	domains := make([]*Domain, 0)
	if err = json.Unmarshal(pr.Entries, &domains); err != nil {
		return nil, err
	}

	return domains, res.Body.Close()
}

func (c *Client) AddDomain(domain *Domain) error {
	if domain == nil {
		return errors.New("go-mailman: Error adding nil domain")
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
		re := &RequestError{}
		if err := json.NewDecoder(res.Body).Decode(re); err != nil {
			return err
		}
		return fmt.Errorf("%s %s", ErrorDomainAdd, re)
	}

	return res.Body.Close()
}

func (c *Client) DeleteDomain(domainID string) error {
	res, err := c.conn.do(http.MethodDelete, c.buildURL("domains", domainID), http.NoBody)
	if err != nil {
		return err
	}

	if res.StatusCode/100 != 2 {
		re := &RequestError{}
		if err := json.NewDecoder(res.Body).Decode(re); err != nil {
			return err
		}

		return fmt.Errorf("%s %s", ErrorDomainDelete, re)
	}

	return res.Body.Close()
}
