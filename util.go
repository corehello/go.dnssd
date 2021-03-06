package dnssd

import (
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"

	"github.com/miekg/dns"
)

// Concatenate a three-part domain name (as provided to the response funcs) into a properly-escaped full domain name.
func ConstructFullName(serviceName, regType, domain string) string {
	return fmt.Sprintf("%s.%s.%s.", serviceName, regType, domain)
}

// domain names are "unpacked" using escape sequences and character
// escapes. Repack them to a proper UTF-8 string
func RepackToUTF8(unpacked string) string {
	us := []rune(unpacked)
	var rs []rune

	for ii := 0; ii < len(us); ii++ {
		if us[ii] == '\\' {
			ii++
			switch us[ii] {
			case 'r':
				rs = append(rs, '\r')
			case 't':
				rs = append(rs, '\t')
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				rb := make([]byte, 2)
				rb[0] = threeDigitsToInt(us[ii:])
				ii += 2
				if us[ii+1] == '\\' {
					rb[1] = threeDigitsToInt(us[ii+2:])
					r, _ := utf8.DecodeLastRune(rb)
					ii += 4
					rs = append(rs, rune(r))
				} else {
					rs = append(rs, rune(rb[0]))
				}
			default:
				rs = append(rs, us[ii])
			}
		} else {
			rs = append(rs, us[ii])
		}

	}
	return string(rs)
}

func threeDigitsToInt(us []rune) byte {
	cc := (int(us[0]) - 48) * 100
	cc += (int(us[1]) - 48) * 10
	cc += int(us[2]) - 48
	return byte(cc)
}

func getNextTime(t1, t2 time.Time) time.Time {
	if t1.IsZero() {
		return t2
	}
	if t2.IsZero() {
		return t1
	}
	if t1.After(t2) {
		return t1
	}
	return t2
}

/*
The determination of whether a given record answers a given question
is made using the standard DNS rules: the record name must match the
question name, the record rrtype must match the question qtype unless
the qtype is "ANY" (255) or the rrtype is "CNAME" (5), and the record
rrclass must match the question qclass unless the qclass is "ANY"
(255). As with Unicast DNS, generally only DNS class 1 ("Internet")
is used, but should client software use classes other than 1, the
matching rules described above MUST be used.
*/
func matchQuestionAndRR(q *dns.Question, rr dns.RR) bool {
	return (q.Qtype == dns.TypeANY || q.Qtype == rr.Header().Rrtype || rr.Header().Rrtype == dns.TypeCNAME) &&
		(q.Qclass == rr.Header().Class) &&
		(q.Name == rr.Header().Name)
}

func matchRRHeader(rr1, rr2 *dns.RR_Header) bool {
	return (rr1.Rrtype == rr2.Rrtype) &&
		(rr1.Class == rr2.Class) &&
		(rr1.Name == rr2.Name)
}

func matchRRs(rr1, rr2 dns.RR) bool {
	return rr1.String() == rr2.String()
}

func matchQuestions(q1, q2 *dns.Question) bool {
	return (q1.Qtype == q2.Qtype) &&
		(q1.Qclass == q2.Qclass) &&
		(q1.Name == q2.Name)
}

// Return a randomDuraion of +/- <variation> percent of the given duration.
func randomDuration(d time.Duration, variation int) time.Duration {
	ud := int64(d)
	vd := rand.Int63n(ud)
	vd *= int64(variation)
	vd /= 100
	return time.Duration(vd)

}
