package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		r := bufio.NewReader(os.Stdin)
		s, _ := r.ReadString('\n')
		s = strings.Replace(s, "\r\n", "", -1)
		msg, err := mcguide(s)
		if err != nil {
			msg = err.Error()
		}

		if msg != "" {
			fmt.Println(msg)
		}
	}
}
