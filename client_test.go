package gomailman

import (
	"fmt"
	"testing"
)

func TestClientAddDomain(t *testing.T) {
	c, err := NewClient(BaseURL, Username, Password)
	if err != nil {
		t.Error(err.Error())
	}

	domain := &Domain{
		AliasDomain: "",
		Description: "Test Domain",
		MailHost:    "test@localhost.com",
	}

	if err := c.AddDomain(domain); err != nil {
		t.Error(err.Error())
	}
}

func TestClientGetDomain(t *testing.T) {
	c, err := NewClient(BaseURL, Username, Password)
	if err != nil {
		t.Error(err.Error())
	}

	domain, err := c.GetDomain("test@localhost.com")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(domain)
}
