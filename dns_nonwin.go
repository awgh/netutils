// +build !windows

package netutils

import (
	"bufio"
	"net"
	"os"
	"strings"
)

// GetDNSForIP - returns all DNS server addresses associated with the given address
func GetDNSForIP(_ net.IP) ([]net.IP, error) {

	// on non-windows, we ignore the ip parameter because routing is not insane

	file, err := os.Open("/etc/resolv.conf")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var retval []net.IP
	r := bufio.NewReader(file)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		if len(line) > 0 && (line[0] == ';' || line[0] == '#') {
			continue // skip comments
		}
		f := strings.Fields(string(line))
		if len(f) < 1 {
			continue
		}
		switch f[0] {
		case "nameserver":
			if len(f) > 1 {
				ip := net.ParseIP(f[1])
				if ip != nil {
					retval = append(retval, ip)
				}
			}
		}
	}
	return retval, nil
}
