package main

import "net"

type NsLookupResponse struct {
	Status ResponseStatus
	IPs []string
}

func NsLookup(url string) NsLookupResponse {
	ips, err := net.LookupIP(url)

	if err != nil {
		response := NsLookupResponse {
			Status: Error,
			IPs: nil,
		}
		return response
	}

	var out []string
	for _, ip := range ips {
		out = append(out, ip.String())
	}

	response := NsLookupResponse {
		Status: Ok,
		IPs: out,
	}

	return response
}