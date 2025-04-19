package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/udhos/equalfile"
)

const problemsDir = "C:\\Users\\neste\\Downloads\\even-strings\\"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
		problem string
		want string
	}{
		{
			problem: "1",
			want: "1.a",
		},
		{
			problem: "2",
			want: "2.a",
		},
		{
			problem: "3",
			want: "3.a",
		},
		{
			problem: "8",
			want: "8.a",
		},
		{
			problem: "9",
			want: "9.a",
		},
		{
			problem: "10",
			want: "10.a",
		},
		{
			problem: "21",
			want: "21.a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run(t, problemsDir + tt.problem, problemsDir + tt.want)
		})
	}
}

func run(t *testing.T, problemFilename, wantFilename string){
	t.Helper()
	
	problemFile, err := os.Open(problemFilename)
	if err != nil {
		require.Fail(t, err.Error())
	}
	defer problemFile.Close()

	gotFile, err := os.CreateTemp(os.TempDir(), "even-strings-")
	if err != nil {
		require.Fail(t, err.Error())
	}
	defer gotFile.Close()

	os.Stdin = problemFile
	os.Stdout = gotFile

	main()

	_, err = gotFile.Seek(0, 0)
	if err != nil {
		require.Fail(t, err.Error())
	}

	wantFile, err := os.Open(wantFilename)
	if err != nil {
		require.Fail(t, err.Error())
	}
	defer wantFile.Close()

	cmp := equalfile.New(nil, equalfile.Options{})
	equal, err := cmp.CompareReader(gotFile, wantFile)
	require.True(t, equal)
}

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
