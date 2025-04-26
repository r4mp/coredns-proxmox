package proxmox

import (
	"strings"

	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

type Config struct {
	backend     string
	tokenId     string
	tokenSecret string
	insecure    bool
	interfaces  string
}

// init registers this plugin.
func init() { plugin.Register("proxmox", setup) }
func setup(c *caddy.Controller) error {

	backend := ""
	tokenId := ""
	tokenSecret := ""
	insecure := ""
	var interfaces []string

	c.Next()
	if c.NextBlock() {
		for {
			switch c.Val() {
			case "backend":
				if !c.NextArg() {
					return plugin.Error("proxmox", c.ArgErr())
				}
				backend = c.Val()
				break
			case "token_id":
				if !c.NextArg() {
					return plugin.Error("proxmox", c.ArgErr())
				}
				tokenId = c.Val()
				break
			case "token_secret":
				if !c.NextArg() {
					return plugin.Error("proxmox", c.ArgErr())
				}
				tokenSecret = c.Val()
				break
			case "insecure":
				if !c.NextArg() {
					return plugin.Error("proxmox", c.ArgErr())
				}
				insecure = c.Val()
				break
			case "interfaces":
				if !c.NextArg() {
					return plugin.Error("proxmox", c.ArgErr())
				}
				interfaces = strings.Split(c.Val(), " ")
				break
			default:
				if c.Val() != "}" {
					return plugin.Error("proxmox", c.Err("unknown property"))
				}
			}
			if !c.Next() {
				break
			}
		}
	}

	if backend == "" || tokenId == "" || tokenSecret == "" || insecure == "" {
		return plugin.Error("proxmox", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Proxmox{
			Backend:     backend,
			TokenId:     tokenId,
			TokenSecret: tokenSecret,
			Insecure:    insecure,
			Interfaces:  interfaces,
			Next:        next,
		}
	})

	return nil
}
