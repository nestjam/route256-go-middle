package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func BenchmarkSum(b *testing.B) {
	filename := "C:\\Users\\neste\\Downloads\\even-strings\\10"
	sets, err := readSetsFromFile(filename)
	if err != nil {
		panic(err)
	}

	b.ResetTimer() // Reset timer to exclude setup time

	for i := 0; i < b.N; i++ {
		findPairs(sets[0])
	}
}

func readSetsFromFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalSets int
	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to read total sets")
	}
	totalSets, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}

	var allSets [][]string
	for i := 0; i < totalSets; i++ {
		if !scanner.Scan() {
			return nil, fmt.Errorf("failed to read set size for set %d", i+1)
		}
		setSize, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		var currentSet []string
		for j := 0; j < setSize; j++ {
			if !scanner.Scan() {
				return nil, fmt.Errorf("failed to read line %d in set %d", j+1, i+1)
			}
			currentSet = append(currentSet, scanner.Text())
		}
		allSets = append(allSets, currentSet)
	}

	return allSets, nil
}
