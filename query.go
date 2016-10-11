package dnssd

import (
	"context"

	"github.com/miekg/dns"
)

const MORE_COMING = 1
const RECORD_ADDED = 8

/* This is called when a query has been resolved.
flags may be MORE_COMING or RECORD_ADDED.
ifIndex is the interface the query was responden on.
rr is a resource record matching the query.
*/
type QueryAnswered func(flags Flags, ifIndex int, rr dns.RR)

/* Query an arbitrary record.
 */
func query(ctx context.Context, flags Flags, ifIndex int, serviceName string, rrtype, rrclass uint16, response QueryAnswered, errc ErrCallback) {
	ds := getDnssd()

	// send the query
	q := &dns.Question{serviceName, rrtype, rrclass}
	cb := &callback{ctx, response}
	ds.cmdCh <- func() {
		ds.runQuery(probe, ctx, q, cb)
	}
}

/*
Query an arbitrary record. ctx is the query context and can be used to cancel or timeout a query.
flags - Possible values are: MORE_COMING.
ifIndex - If non-zero, specifies the interface on which to issue the query (the index for a given interface is determined via the if_nametoindex() family of calls.) Passing 0 causes the name to be queried for on all interfaces. Passing -1 causes the name to be queried for only on the local host.
question - The question to query for.
response - This closure will get called when the query completes.
errc - This closure will be called when a query has an error.
*/
func Query(ctx context.Context, flags Flags, ifIndex int, serviceName string, rrtype, rrclass uint16, response QueryAnswered, errc ErrCallback) {
	query(ctx, flags, ifIndex, serviceName, rrtype, rrclass, response, errc)
}

// Instruct the daemon to verify the validity of a resource record that appears to be out of date.
func ReconfirmRecord(flags Flags, ifIndex int, rr *dns.RR) {
	panic("NYI")
}
