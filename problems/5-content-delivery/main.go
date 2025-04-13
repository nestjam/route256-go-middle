package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "net/http"
	// _ "net/http/pprof" // Import pprof for profiling
)

func main() {
	// // Start a server for pprof endpoints
	// go func() {
	// 	fmt.Println("Starting pprof server at http://localhost:6060")
	// 	http.ListenAndServe("localhost:6060", nil)
	// }()

	var in *bufio.Reader
	//in = bufio.NewReader(os.Stdin)
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	file, _ := os.Open("C:/Users/neste/Downloads/content-delivery/15")
	defer file.Close() // Ensure the file is closed when the function exits
	in = bufio.NewReader(file)
	fmt.Fprintln(out)

	var setCount int
	var serverCount, imageCount int
	var serverThrouputs, imageWeights []int

	fmt.Fscan(in, &setCount)

	for i := 0; i < setCount; i++ {
		fmt.Fscan(in, &serverCount)
		serverThrouputs = fscanIntegers(in, serverCount)

		fmt.Fscan(in, &imageCount)
		imageWeights = fscanIntegers(in, imageCount)

		delta, imageStorages := distribute(imageWeights, serverThrouputs)

		fmt.Fprintln(out, delta)
		fprintIntegers(out, imageStorages)
	}

	// fmt.Println("Application running...")
	// select {} // Block forever to keep the app running
}

func fscanIntegers(in *bufio.Reader, count int) []int {
	values := make([]int, count)
	var value int
	for j := 0; j < count; j++ {
		fmt.Fscan(in, &value)
		values[j] = value
	}
	return values
}

func fprintIntegers(out *bufio.Writer, input []int) {
	for i := 0; i < len(input); i++ {
		if i == len(input)-1 {
			fmt.Fprintln(out, strconv.Itoa(input[i]))
		} else {
			fmt.Fprint(out, strconv.Itoa(input[i]), " ")
		}
	}
}
