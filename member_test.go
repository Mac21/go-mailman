package gomailman

import (
	"testing"
)

const (
	memberDisplayName = "Jimmy Johns"
	memberEmail       = "jimmyjohns@localhost.com"
)

func TestAddListMember(t *testing.T) {
	c, err := NewClient(baseURL, username, password)
	if err != nil {
		t.Error(err)
	}

	sub := &ListSubscriber{
		ListID:       listID,
		Subscriber:   memberEmail,
		DisplayName:  memberDisplayName,
		PreVerified:  true,
		PreConfirmed: true,
		PreApproved:  true,
	}

	if err = c.AddListMember(sub); err != nil {
		t.Error(err)
	}
}

func TestGetListMembers(t *testing.T) {
	c, err := NewClient(baseURL, username, password)
	if err != nil {
		t.Error(err)
	}

	members, err := c.GetListMembers(listID)
	if err != nil {
		t.Error(err)
	}

	if len(members) < 1 {
		t.Errorf("Error list %s expected more than one member got none", listID)
	}
}
