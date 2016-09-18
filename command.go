package dnssd

import (
	"fmt"

	"github.com/miekg/dns"
)

type command struct {
	q *dns.Msg
	r interface{}
}

func (c *command) String() string {
	return fmt.Sprint("command{", c.q, "}")
}

func (cmd *command) match(q dns.Question, answer dns.RR) bool {
	if q.Qtype == answer.Header().Rrtype {
		if q.Name == answer.Header().Name {
			respond(cmd.r, answer)
			return true
		}
	}
	return false
}

func respond(r interface{}, rr dns.RR) {
	switch r := r.(type) {
	case QueryAnswered:
		r(nil, 0, 0, rr)
	default:
		panic(fmt.Sprint("Dont know what", r, " is"))
	}
}

func respondWithError(r interface{}, err error) {
	switch r := r.(type) {
	case QueryAnswered:
		r(err, 0, 0, nil)
	default:
		panic(fmt.Sprint("Dont know what", r, " is"))
	}
}
