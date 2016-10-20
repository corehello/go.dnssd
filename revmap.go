package dnssd

/*
4.  Reverse Address Mapping

   Like ".local.", the IPv4 and IPv6 reverse mapping domains are also
   defined to be link-local:

      Any DNS query for a name ending with "254.169.in-addr.arpa." MUST
      be sent to the mDNS IPv4 link-local multicast address 224.0.0.251
      or the mDNS IPv6 multicast address FF02::FB.  Since names under
      this domain correspond to IPv4 link-local addresses, it is logical
      that the local link is the best place to find information
      pertaining to those names.

      Likewise, any DNS query for a name within the reverse mapping
      domains for IPv6 link-local addresses ("8.e.f.ip6.arpa.",
      "9.e.f.ip6.arpa.", "a.e.f.ip6.arpa.", and "b.e.f.ip6.arpa.") MUST
      be sent to the mDNS IPv6 link-local multicast address FF02::FB or
      the mDNS IPv4 link-local multicast address 224.0.0.251.
*/
