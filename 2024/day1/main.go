package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(run(os.Args[1:]))
}

func run(args []string) string {
	if len(args) != 1 {
		return "Usage: go run main.go <input-file>"
	}
	fileName := args[0]
	contents, err := os.ReadFile(fileName)
	if err != nil {
		return err.Error()
	}
	return string(contents)
}
