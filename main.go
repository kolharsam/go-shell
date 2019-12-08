package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// isError function is used to handle errors
func isError(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
	}
}

// executeCommand does like the name suggests - executes the command
func executeCommand(input string) error {
	input = strings.TrimSuffix(input, "\n")

	cmd := exec.Command(input)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("go-shell> $ ")
		input, err := reader.ReadString('\n')
		isError(err)

		err = executeCommand(input)
		isError(err)
	}
}
