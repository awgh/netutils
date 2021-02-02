// +build windows

package netutils

import (
	"net"
	"unsafe"

	"golang.org/x/sys/windows"
)

// GetDNSForIP - returns all DNS server addresses associated with the given address
func GetDNSForIP(ip net.IP) ([]net.IP, error) {

	l := uint32(20000)
	b := make([]byte, l)
	var retval []net.IP

	if err := windows.GetAdaptersAddresses(windows.AF_UNSPEC,
		windows.GAA_FLAG_INCLUDE_PREFIX, 0, (*windows.IpAdapterAddresses)(unsafe.Pointer(&b[0])), &l); err != nil {
		return nil, err
	}

	var addresses []*windows.IpAdapterAddresses
	for addr := (*windows.IpAdapterAddresses)(unsafe.Pointer(&b[0])); addr != nil; addr = addr.Next {
		addresses = append(addresses, addr)
	}
	for _, a := range addresses {
		for next := a.FirstUnicastAddress; next != nil; next = next.Next {
			if next.Address.IP().Equal(ip) {
				for dns := a.FirstDnsServerAddress; dns != nil; dns = dns.Next {
					retval = append(retval, dns.Address.IP())
				}
				break
			}
		}
	}
	return retval, nil
}
