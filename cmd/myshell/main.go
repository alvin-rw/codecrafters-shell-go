package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")

		cmd, args, err := readInput(rd)
		if err != nil {
			fmt.Printf("error when reading command %s", err)
			os.Exit(1)
		}

		switch cmd {
		case "exit":
			exitCommand(args[0])
		case "echo":
			echoCommand(args)
		case "type":
			typeCommand(args[0])
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}

func readInput(rd *bufio.Reader) (string, []string, error) {
	userInput, err := rd.ReadString('\n')
	if err != nil {
		return "", nil, err
	}

	userInput = strings.TrimSuffix(userInput, "\n")
	userInputSlice := strings.Split(userInput, " ")

	command := userInputSlice[0]

	var args []string
	if len(userInputSlice) > 1 {
		args = userInputSlice[1:]
	}

	return command, args, nil
}

func exitCommand(input string) error {
	statusCode, err := strconv.Atoi(input)
	if err != nil {
		return err
	}

	os.Exit(statusCode)
	return nil
}

func echoCommand(input []string) {
	out := strings.Join(input, " ")

	fmt.Printf("%s\n", out)
}

func typeCommand(input string) {
	builtinCommands := []string{
		"echo",
		"type",
		"exit",
	}

	for _, cmd := range builtinCommands {
		if input == cmd {
			fmt.Printf("%s is a shell builtin\n", input)
			return
		}
	}

	pathCommands := readPathEnv()
	for _, dir := range pathCommands {
		s := strings.Split(dir, "/")
		command := s[len(s)-1]

		if input == command {
			fmt.Printf("%s is %s\n", command, dir)
			return
		}
	}

	fmt.Printf("%s: not found\n", input)
}

func readPathEnv() []string {
	pathValue := os.Getenv("PATH")

	availableCommands := strings.Split(pathValue, ":")

	return availableCommands
}
