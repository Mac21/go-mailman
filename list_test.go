package gomailman

import (
	"strings"
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
	c, err := NewClient(baseURL, username, password)
	if err != nil {
		t.Error(err)
	}

	list, err := c.GetList(listID)
	if err != nil {
		t.Fatal(err)
	}

	if list.MailHost != domainID {
		t.Errorf("Unexpected list MailHost got %s expected %s", list.MailHost, domainID)
	}

	listsID := strings.Replace(listID, "@", ".", -1)
	if list.ID != listsID {
		t.Errorf("Unexpected list ID got %s expected %s", list.ID, listsID)
	}
}

func TestDeleteList(t *testing.T) {
	c, err := NewClient(baseURL, username, password)
	if err != nil {
		t.Error(err)
	}

	if err := c.DeleteList(listID); err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		c.DeleteDomain(domainID)
	})
}
