package gomailman

import (
	"testing"
)

func TestClientInit(t *testing.T) {
	c, err := NewClient(baseURL, username, password)
	if err != nil {
		t.Error(err.Error())
	}
}
