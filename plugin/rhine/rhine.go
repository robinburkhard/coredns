// Package rhine implements a plugin that returns additional rhine certificate TXT RRs for zone authentication
package rhine

import (
	"context"
	"fmt"
	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
)

const name = "rhine"

// Rhine is a plugin that returns additional rhine certificate TXT RRs for zone authentication
type Rhine struct{
	Next plugin.Handler
}

// ServeDNS implements the plugin.Handler interface.
func (rh Rhine) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	fmt.Println("Hiii ")
	fmt.Println(r.Question)
	fmt.Println(r.Question[0].Name)
	r.Question = append(r.Question, dns.Question{
		Name:   "host1.rhine-ns.com",
		Qtype:  dns.TypeTXT,
		Qclass: 1,
	})

	fmt.Println(r.Question)

	//plugin.NextOrFailure(rh.Name(), rh.Next, ctx, w, r)
	r.Question[0] = dns.Question{
		Name:   "host1.rhine-ns.com",
		Qtype:  dns.TypeTXT,
		Qclass: 1,
	}
	plugin.NextOrFailure(rh.Name(), rh.Next, ctx, w, r)

	return 0, nil
}

// Name implements the Handler interface.
func (rh Rhine) Name() string { return name }
