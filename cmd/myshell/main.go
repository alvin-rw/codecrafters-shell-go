package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("error when reading command %s", err)
			os.Exit(1)
		}
		cmd = strings.TrimSuffix(cmd, "\n")

		fmt.Printf("%s: command not found\n", cmd)
	}
}
