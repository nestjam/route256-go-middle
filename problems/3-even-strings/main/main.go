package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setCount int
	var count int
	var s string

	fmt.Fscan(in, &setCount)

	for i := 0; i < setCount; i++ {
		fmt.Fscan(in, &count)

		strings := make([]string, count)

		for j := 0; j < count; j++ {
			fmt.Fscan(in, &s)
			strings[j] = s
		}

		fmt.Fprint(out, findPairs(strings), "\n")
	}
}
