package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	for {
		fmt.Println("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		commands := strings.Split(input, "|")
		var cmd *exec.Cmd

		for i, command := range commands {
			args := strings.Fields(command)
			switch args[0] {
			case "cd":
				if len(args) > 1 {
					err := os.Chdir(args[1])
					if err != nil {
						fmt.Println("Error changing directory:", err)
					}
				} else {
					fmt.Println("Usage: cd <directory>")
				}

			case "pwd":
				currentDir, err := os.Getwd()
				if err != nil {
					fmt.Println("Error getting current directory:", err)
				}
				fmt.Println(currentDir)

			case "echo":
				fmt.Println(strings.Join(args[1:], " "))

			case "kill":
				if len(args) > 1 {
					processID := args[1]
					cmd := exec.Command("kill", processID)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					err := cmd.Run()
					if err != nil {
						fmt.Println("Error killing process:", err)
					}
				} else {
					fmt.Println("Usage: kill <processID>")
				}

			case "ps":
				cmd = exec.Command("ps")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

			case "q":
				os.Exit(0)

			default:
				if i == 0 {
					cmd = exec.Command(args[0], args[1:]...)
				} else {
					cmd.Stdin, _ = cmd.StdoutPipe()
					cmd = exec.Command(args[0], args[1:]...)
				}
			}
		}

		if cmd != nil {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error executing command:", err)
			}
		}
	}
}
