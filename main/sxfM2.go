package main

import (
	"fmt"
	"strings"
	"strconv"
)
// ip checks if a given string is a valid IPv4 address.
func ip(ip string) bool {
	// Split the IP address into parts by '.'
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}

	for _, part := range parts {
		// Each part should be a number between 0 and 255
		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}

	return true
}

func main() {
	tests := []string{
		"192.168.0.1",
		"256.256.256.256",
		"::1",
		"abcd",
		"10.0.0.256",
		"2001:db8::ff00:42:8329",
	}

	for _, test := range tests {
		if ip(test) {
			fmt.Printf("%s is a valid IPv4 address\n", test)
		} else {
			fmt.Printf("%s is not a valid IPv4 address\n", test)
		}
	}
}
