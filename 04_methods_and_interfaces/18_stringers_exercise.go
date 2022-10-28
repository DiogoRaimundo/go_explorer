package goTour04

import (
	"fmt"
)

// One of the most ubiquitous interfaces is Stringer defined by the fmt package.
// A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
type Stringer interface {
	String() string
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func RunExercise18() {
	// Stringer example
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
