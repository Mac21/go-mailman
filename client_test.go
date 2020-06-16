package gomailman

import (
	"testing"
)

func TestClientInit(t *testing.T) {
	_, err := NewClient(baseURL, username, password)
	if err != nil {
		t.Error(err.Error())
	}
}
