package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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

		if strings.Contains(s, "ERROR") {
			fmt.Println(s)
		}
	}
}
