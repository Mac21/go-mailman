package gomailman

import (
	"bytes"
	"net/http"
	"strings"
	"testing"
)

const (
	BaseURL  = "http://localhost:8001/3.1"
	Username = "restadmin"
	Password = "restpass"
)

func TestConnectionBasicAuthRejected(t *testing.T) {
	conn, err := NewConnection(BaseURL, "jimmer", Password)
	if err != nil {
		t.Errorf("Connection failed to create: %s", err.Error())
	}

	res, err := conn.do(http.MethodGet, "system/versions", bytes.NewBuffer(nil))
	if err != nil {
		t.Errorf("Connection request failed: %s", err.Error())
	}

	b := bytes.NewBuffer(nil)
	b.ReadFrom(res.Body)

	result := b.String()
	if !strings.Contains(result, "401 Unauthorized") {
		t.Error("Basic authentication was sucessful expected 401 Unauthorized")
	}

}

func TestConnectionNilUsername(t *testing.T) {
	_, err := NewConnection(BaseURL, "", Password)
	if err == nil {
		t.Error("Connection successfully created with blank username")
	}
}

func TestConnectionNilPassword(t *testing.T) {
	_, err := NewConnection(BaseURL, Username, "")
	if err == nil {
		t.Error("Connection successfully created with blank password")
	}
}

func TestConnectionNilUsernamePassword(t *testing.T) {
	_, err := NewConnection(BaseURL, "", "")
	if err == nil {
		t.Error("Connection successfully created with blank username and password")
	}
}
