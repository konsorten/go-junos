package junos

import (
	"testing"
)

func TestNewSession(t *testing.T) {
	auth := &AuthMethod{
		Credentials: []string{"read-only", "ACwv!trnn6UHxcSm"},
	}

	jnpr, err := NewSession("juniper4", auth)
	if err != nil {
		t.Fatal(err)
	}

	jnpr.Close()
}
