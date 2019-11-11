package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		tokens := strings.Split(str, *sepChar)
		for i := 0; i < len(tokens); i++ {
			token := strings.TrimSpace(tokens[i])
			intValue, err := strconv.Atoi(token)
			if err != nil {
				println("token", token, "was not convertable to integer intValue")
			}
			fmt.Printf("%c", intValue)
		}

		fmt.Println()
	}
}
