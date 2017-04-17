package main

import "fmt"

type IPAddr [4]byte

// IPAddr型の値を「x.x.x.x」の形する
func (adder IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", adder[0], adder[1], adder[2], adder[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %s\n", name, ip)
	}
}
