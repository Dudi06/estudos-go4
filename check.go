package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination string, port string) string {
	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v não pode ser alcançado, \n Error: %v", destination, port)
	} else {
		status = fmt.Sprintf("[UP] %v é alcançavel,\n De: %v\n Para: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}
