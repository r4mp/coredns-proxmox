package proxmox

import (
	"testing"

	"github.com/coredns/caddy"
)

// TestSetup tests the various things that should be parsed by setup.
// Make sure you also test for parse errors.
func TestSetup(t *testing.T) {
	c := caddy.NewTestController("dns", `proxmox {
backend "https://jupiter.renner.uno:8006/api2/json/"
token_id "root@pam!cdns-dev"
token_secret "afe4c1a4-29a5-472a-8b8b-00c4c0b36b7d"
}`)
	if err := setup(c); err != nil {
		t.Fatalf("Expected no errors, but got: %v", err)
	}
}
