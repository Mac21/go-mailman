package gomailman

import (
	"bytes"
	"net/http"
	"strings"
	"testing"
)

const (
	baseURL  = "http://localhost:8001/3.1"
	username = "restadmin"
	password = "restpass"
)

func TestConnectionBasicAuthRejected(t *testing.T) {
	conn, err := NewConnection(baseURL, "jimmer", password)
	if err != nil {
		t.Fatalf("Connection failed to create: %s", err)
	}

	res, err := conn.do(http.MethodGet, "system/versions", bytes.NewBuffer(nil))
	if err != nil {
		t.Fatalf("Connection request failed: %s", err)
	}

	b := bytes.NewBuffer(nil)
	b.ReadFrom(res.Body)

	result := b.String()
	if !strings.Contains(result, "401 Unauthorized") {
		t.Error("Basic authentication was sucessful expected 401 Unauthorized")
	}

}

func TestConnectionNilUsername(t *testing.T) {
	_, err := NewConnection(baseURL, "", password)
	if err == nil {
		t.Error("Connection successfully created with blank username")
	}
}

func TestConnectionNilPassword(t *testing.T) {
	_, err := NewConnection(baseURL, username, "")
	if err == nil {
		t.Error("Connection successfully created with blank password")
	}
}

func TestConnectionNilUsernamePassword(t *testing.T) {
	_, err := NewConnection(baseURL, "", "")
	if err == nil {
		t.Error("Connection successfully created with blank username and password")
	}
}
