package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner

	// TODO: write your code here

	fmt.Println(scanner.Text())
}
