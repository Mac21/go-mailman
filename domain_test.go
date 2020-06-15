package gomailman

import (
	"strings"
	"testing"
)

const (
	domainAlias       = "jimbo@localhost.com"
	domainID          = "test@localhost.com"
	domainDescription = "Test Domain"
)

func TestClientAddDomain(t *testing.T) {
	c, err := NewClient(baseURL, username, password)
	if err != nil {
		t.Error(err.Error())
	}

	domain := &Domain{
		AliasDomain: domainAlias,
		Description: domainDescription,
		MailHost:    domainID,
	}

	if err := c.AddDomain(domain); err != nil {
		t.Error(err.Error())
	}
}

func TestClientGetDomain(t *testing.T) {
	c, err := NewClient(baseURL, username, password)
	if err != nil {
		t.Error(err.Error())
	}

	domain, err := c.GetDomain(domainID)
	if err != nil {
		t.Error(err.Error())
	}

	if domain.AliasDomain != domainAlias {
		t.Errorf("Unexpected domain AliasDomain got %s expected %s", domain.AliasDomain, domainAlias)
	}

	if domain.MailHost != domainID {
		t.Errorf("Unexpected domain MailHost got %s expected %s", domain.MailHost, domainID)
	}

	if domain.Description != domainDescription {
		t.Errorf("Unexpected domain Description got %s expected %s", domain.Description, domainDescription)
	}
}

func TestClientDeleteDomain(t *testing.T) {
	c, err := NewClient(baseURL, username, password)
	if err != nil {
		t.Error(err.Error())
	}

	err = c.DeleteDomain(domainID)
	if err != nil {
		t.Error(err.Error())
	}

	domain, err := c.GetDomain(domainID)
	if err != nil {
		e := err.Error()
		if !strings.Contains(e, ErrorDomainGet) {
			t.Error(e)
		}
	}

	if domain != nil {
		t.Error("Got non nil domain after delete")
	}
}