package validators

import "net"


func ValidateHost(host string) bool {
	host, _, err := net.SplitHostPort(host)
	if err != nil {
		return false
	}

	ip := net.ParseIP(host)
	if ip == nil {
		return false
	}

	return true
}