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
			exit(args[0])
		case "echo":
			echo(args)
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

func exit(input string) error {
	statusCode, err := strconv.Atoi(input)
	if err != nil {
		return err
	}

	os.Exit(statusCode)
	return nil
}

func echo(input []string) {
	out := strings.Join(input, " ")

	fmt.Printf("%s\n", out)
}
