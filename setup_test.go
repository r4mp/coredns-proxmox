package proxmox

import (
	"testing"

	"github.com/coredns/caddy"
)

// TestSetup tests the various things that should be parsed by setup.
// Make sure you also test for parse errors.
func TestSetup(t *testing.T) {
	c := caddy.NewTestController("dns", `proxmox {
backend "https://proxmox.example.com:8006/api2/json/"
token_id "coredns@pve!coredns"
token_secret "xyaaaa-b4cd-cde5-abc4-1234567"
insecure false
}`)
	if err := setup(c); err != nil {
		t.Fatalf("Expected no errors, but got: %v", err)
	}
}
