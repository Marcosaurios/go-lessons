package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var ErrorType string

func main() {
	flag.StringVar(&ErrorType, "error", "ERROR", "Describes the log type to scan.")
	flag.Parse()

	f, err := os.Open("./logs.log")

	if err != nil {
		fmt.Println("Problem reading the file, %v", err)
		panic("Problem reading")
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')

		if err != nil {
			fmt.Errorf("There was a problem reading the line")
			break
		}

		if strings.Contains(s, ErrorType) {
			fmt.Println(s)
		}
	}
}
