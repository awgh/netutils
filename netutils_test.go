package netutils

import (
	"fmt"
	"testing"
)

func Test_NetUtils_DNS_1(t *testing.T) {

	outboundIP := GetOutboundIP()
	ips, err := GetDNSForIP(outboundIP)
	fmt.Printf("%+v %s\n", ips, err)

}
