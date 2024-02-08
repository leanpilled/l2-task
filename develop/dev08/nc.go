package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	var (
		host string
		port int
		udp  bool
	)

	flag.StringVar(&host, "host", "localhost", "Host to connect to")
	flag.IntVar(&port, "port", 8080, "Port to connect to")
	flag.BoolVar(&udp, "udp", false, "Use UDP instead of TCP")

	flag.Parse()

	address := fmt.Sprintf("%s:%d", host, port)
	fmt.Println(address)

	fmt.Printf("Connecting to %s via ", address)
	if udp {
		fmt.Println("UDP")
	} else {
		fmt.Println("TCP")
	}

	conn, err := createConnection(udp, address)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	go receiveData(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		conn.Write([]byte(message + "\n"))
	}
}

func createConnection(udp bool, address string) (net.Conn, error) {
	var conn net.Conn
	var err error
	if udp {
		conn, err = net.Dial("udp", address)
	} else {
		conn, err = net.Dial("tcp", address)
	}
	return conn, err
}

func receiveData(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("Received:", scanner.Text())
	}
}
