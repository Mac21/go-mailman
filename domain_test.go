package gomailman

import (
	"strings"
	"testing"
)

const (
	domainAlias       = "jimbo@localhost.com"
	domainID          = "localhost.com"
	domainDescription = "Test Domain"
)

func TestAddDomain(t *testing.T) {
	c, err := NewClient(baseURL, username, password)
	if err != nil {
		t.Error(err)
	}

	domain := &Domain{
		AliasDomain: domainAlias,
		Description: domainDescription,
		MailHost:    domainID,
	}

	if err := c.AddDomain(domain); err != nil {
		t.Error(err)
	}
}

func TestGetDomain(t *testing.T) {
	c, err := NewClient(baseURL, username, password)
	if err != nil {
		t.Error(err)
	}

	domain, err := c.GetDomain(domainID)
	if err != nil {
		t.Fatal(err)
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

func TestDeleteDomain(t *testing.T) {
	c, err := NewClient(baseURL, username, password)
	if err != nil {
		t.Error(err)
	}

	err = c.DeleteDomain(domainID)
	if err != nil {
		t.Error(err)
	}

	domain, err := c.GetDomain(domainID)
	if err != nil {
		e := err.Error()
		if !strings.Contains(e, "404 Not Found") {
			t.Error(e)
		}
	}

	if domain != nil {
		t.Error("Got non nil domain after delete")
	}
}

func TestDeleteMissingDomain(t *testing.T) {
	c, err := NewClient(baseURL, username, password)
	if err != nil {
		t.Error(err)
	}

	err = c.DeleteDomain("ldkjsf@localhost.com")
	if err != nil {
		if !strings.Contains(err.Error(), "404 Not Found") {
			t.Error(err)
		}
	}
}

func TestGetAllDomains(t *testing.T) {
	c, err := NewClient(baseURL, username, password)
	if err != nil {
		t.Error(err.Error())
	}

	mailHostsLoaded := map[string]bool{
		"swag.localhost.com":    false,
		"swagout.localhost.com": false,
		"swaggin.localhost.com": false,
		"swang.localhost.com":   false,
		"swig.localhost.com":    false,
		"swiggin.localhost.com": false,
		"swing.localhost.com":   false,
		"sing.localhost.com":    false,
		"swingin.localhost.com": false,
		"swagit.localhost.com":  false,
	}

	t.Cleanup(func() {
		for host, _ := range mailHostsLoaded {
			_ = c.DeleteDomain(host)
		}
	})

	placeHolder := &Domain{}
	for host, _ := range mailHostsLoaded {
		_ = c.DeleteDomain(host)
		placeHolder.MailHost = host
		err := c.AddDomain(placeHolder)
		if err != nil {
			t.Error(err)
		}
	}

	domains, err := c.GetAllDomains()
	if err != nil {
		t.Error(err)
	}

	for _, d := range domains {
		loaded, exists := mailHostsLoaded[d.MailHost]
		if !exists {
			t.Errorf("Unknown domain loaded: %s", d.MailHost)
		}

		if !loaded {
			mailHostsLoaded[d.MailHost] = true
		}
	}

	for host, loaded := range mailHostsLoaded {
		if !loaded {
			t.Errorf("Domain %s not loaded when expected", host)
		}
	}
}
