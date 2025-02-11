package main

import (
	"bufio"
	"fmt"
	"github.com/YA2IR/sha1"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Welcome to SHA1 REPL, write 'exit' to exit\n> ")
	h := sha1.NewSHA1()

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "exit" {
			break
		}
		result := h.Hash([]byte(line))
		fmt.Printf("SHA1: \033[32m%x\033[0m\n", result)
		fmt.Print("> ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}
}
