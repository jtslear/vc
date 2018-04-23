package main

import (
	"crypto/tls"
	"flag"
	"fmt"
)

func testTLS(host string) (valid bool, reason string) {
	conn, err := tls.Dial("tcp", host, &tls.Config{})
	if err != nil {
		return false, err.Error()
	}
	conn.Close()
	return true, ""
}

func main() {
	flag.Parse()
	for _, v := range flag.Args() {
		ok, msg := testTLS(v)
		if !ok {
			fmt.Printf("%s: notok: %s\n", v, msg)
		} else {
			fmt.Printf("%s: ok\n", v)
		}
	}
}
