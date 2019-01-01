package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
	var s []string
	for _, i := range ipAddr {
		s = append(s, strconv.Itoa(int(i)))
	}
	return fmt.Sprintf(strings.Join(s, "."))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
