package gomailman

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Domain struct {
	AliasDomain string `json:"alias_domain,omitempty"`
	Description string `json:"description,omitempty"`
	MailHost    string `json:"mail_host"`
}

func (c *Client) GetDomain(domainID string) (*Domain, error) {
	res, err := c.conn.do(http.MethodGet, buildURL("domains", domainID), http.NoBody)
	if err != nil {
		return nil, err
	}

	if err := parseResponseError(res); err != nil {
		return nil, err
	}

	domain := new(Domain)
	if err = json.NewDecoder(res.Body).Decode(domain); err != nil {
		return nil, err
	}

	return domain, res.Body.Close()
}

func (c *Client) GetAllDomains() ([]*Domain, error) {
	res, err := c.conn.do(http.MethodGet, buildURL("domains"), http.NoBody)
	if err != nil {
		return nil, err
	}

	if err := parseResponseError(res); err != nil {
		return nil, err
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

	res, err := c.conn.do(http.MethodPost, buildURL("domains"), bytes.NewReader(b))
	if err != nil {
		return err
	}

	if err := parseResponseError(res); err != nil {
		return err
	}

	return res.Body.Close()
}

func (c *Client) DeleteDomain(domainID string) error {
	res, err := c.conn.do(http.MethodDelete, buildURL("domains", domainID), http.NoBody)
	if err != nil {
		return err
	}

	if err := parseResponseError(res); err != nil {
		return err
	}

	return res.Body.Close()
}
