package gomailman

import (
	"bytes"
	"net/http"
	"testing"
)

const (
	BaseURL  = "localhost"
	Username = "user"
	Password = "changeme"
)

func ConnectionBasicAuthRejectedTest(t testing.T) {
	conn, err := NewConnection(BaseURL, Username, Password)
	if err != nil {
		t.Errorf("Connection failed to create: %s", err.Error())
	}

	_, err = conn.do(http.MethodGet, "system/version", bytes.NewBuffer(nil))
	if err != nil {
		t.Errorf("Connection request failed: %s", err.Error())
	}
}

func ConnectionNilUsernameTest(t testing.T) {
	_, err := NewConnection(BaseURL, "", Password)
	if err == nil {
		t.Error("Connection successfully created with blank username")
	}
}

func ConntionNilPasswordTest(t testing.T) {
	_, err := NewConnection(BaseURL, Username, "")
	if err == nil {
		t.Error("Connection successfully created with blank password")
	}
}

func ConntionNilUsernamePasswordTest(t testing.T) {
	_, err := NewConnection(BaseURL, "", "")
	if err == nil {
		t.Error("Connection successfully created with blank username and password")
	}
}
