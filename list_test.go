package gomailman

import (
	"testing"
)

const (
	listID          = "something@localhost.com"
	listDescription = "A list about something"
)

func TestAddList(t *testing.T) {
	c, err := NewClient(baseURL, username, password)
	if err != nil {
		t.Error(err)
	}

	domain := &Domain{
		MailHost:    domainID,
		Description: domainDescription,
	}

	if err := c.AddDomain(domain); err != nil {
		t.Error(err)
	}

	if err := c.AddList(listID); err != nil {
		t.Error(err)
	}
}

func TestGetList(t *testing.T) {
}
