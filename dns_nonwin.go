// +build !windows

package netutils

import (
	"net"
)

// GetDNSForIP - returns all DNS server addresses associated with the given address
func GetDNSForIP(ip net.IP) ([]net.IP, error) {
	return nil, nil
}
