package main

import (
	"bufio"
	"fmt"
	"strings"
)

func getUserInput(question string, reader *bufio.Reader) (string, error) {
	fmt.Print(question);
	resp, error := reader.ReadString('\n')
	
	return strings.TrimSpace(resp), error
}
