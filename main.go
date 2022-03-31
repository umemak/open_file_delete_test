package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	_, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	fmt.Println("file opened. hit any key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}
