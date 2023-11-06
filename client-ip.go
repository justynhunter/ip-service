package main

import (
	"regexp"
	"strings"
)

type IpLookupStatus string

type IpLookupResponse struct {
	Status ResponseStatus
	IPs []string
}

func GetClientIp(ip string, ips []string) IpLookupResponse {
  re := regexp.MustCompile(`(^127\.)|(^10\.)|(^172\.1[6-9]\.)|(^172\.2[0-9]\.)|(^172\.3[0-1]\.)|(^192\.168\.)`)
  if re.MatchString(ip) {
  	if (len(ips) > 0) {
      for _, i := range ips {
        ip = ip + "\n" + i
      }
  	}
  }

	return IpLookupResponse {
		Status: Ok,
		IPs: strings.Split(ip, "\n"),
	}
}