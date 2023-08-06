package examples

import (
	"bufio"
	"fmt"
	"os"
)

func Interactive(
	process func(request string),
	prompt string,
) {
	for {
		fmt.Print(prompt)

		in := bufio.NewReader(os.Stdin)
		request, err := in.ReadString('\n')
		if err != nil {
			panic(err)
		}

		process(request)
	}
}
