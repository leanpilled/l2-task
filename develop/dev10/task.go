package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	var timeoutDuration string
	flag.StringVar(&timeoutDuration, "timeout", "10s", "Connection timeout duration")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: go-telnet --timeout=<timeout> <host> <port>")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	timeout, err := time.ParseDuration(timeoutDuration)
	if err != nil {
		fmt.Println("Error parsing timeout:", err)
		os.Exit(1)
	}

	address := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	go func() {
		io.Copy(os.Stdout, conn)
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		conn.Close()
		os.Exit(0)
	}()

	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil && err != io.EOF {
			fmt.Println("Error copying to connection:", err)
		}
	}()

	select {}
}
