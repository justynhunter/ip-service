package main

import "net"

type NsLookupStatus string

const (
	Ok NsLookupStatus = "OK"
	Error NsLookupStatus = "ERROR"
)

type NsLookupResponse struct {
	Status NsLookupStatus
	IPs []string
}

func nsLookup(url string) NsLookupResponse {
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