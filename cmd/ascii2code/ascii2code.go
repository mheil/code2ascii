package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	sepChar := flag.String("sep", ",", "separator character")
	flag.Parse()

	args := flag.Args()

	var reader *bufio.Reader
	if len(args) > 0 {
		reader = bufio.NewReader(strings.NewReader(args[0] + "\n"))
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		var numChars = len(line) - 1
		for i := 0; i < numChars; i++ {
			fmt.Print(fmt.Sprint(line[i]))
			if i < numChars-1 {
				fmt.Print(*sepChar)
			}
		}

		fmt.Println()
	}
}
