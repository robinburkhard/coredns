package rhine

import (
"github.com/coredns/caddy"
"github.com/coredns/coredns/core/dnsserver"
"github.com/coredns/coredns/plugin"
)

func init() { plugin.Register("rhine", setup) }

func setup(c *caddy.Controller) error {
	c.Next() // 'whoami'
	if c.NextArg() {
		return plugin.Error("rhine", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Rhine{Next: next}
	})

	return nil
}