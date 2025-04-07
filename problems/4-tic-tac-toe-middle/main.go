package main

import (
	"bufio"
	"fmt"
	"os"
	"net/http"
	_ "net/http/pprof" // Import pprof for profiling
)

func main() {
	// Start a server for pprof endpoints
	go func() {
		fmt.Println("Starting pprof server at http://localhost:6060")
		http.ListenAndServe("localhost:6060", nil)
	}()

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setCount int
	var k, n, m int
	var line string

	fmt.Fscan(in, &setCount)

	for i := 0; i < setCount; i++ {
		fmt.Fscan(in, &k)
		fmt.Fscan(in, &n, &m)

		cells := make([][]cell, n)

		for j := 0; j < n; j++ {
			fmt.Fscan(in, &line)
			rowCells := make([]cell, m)

			for l := 0; l < len(line); l++ {
				if line[l] == 'X' {
					rowCells[l] = x
				} else if line[l] == 'O' {
					rowCells[l] = o
				} else if line[l] == '.' {
					rowCells[l] = e
				} else {
					panic("not supported")
				}
			}

			cells[j] = rowCells
		}

		board := newBoard(cells)

		if canCrossWin(&board, k) {
			fmt.Fprint(out, "YES\n")
		} else {
			fmt.Fprint(out, "NO\n")
		}
	}

	fmt.Println("Application running...")
	select {} // Block forever to keep the app running
}
