package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/udhos/equalfile"
)

const problemsDir = "C:\\Users\\neste\\Downloads\\content-delivery\\"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
		problem string
		want string
	}{
		{
			problem: "1",
		},
		{
			problem: "2",
		},
		{
			problem: "3",
		},
		{
			problem: "4",
		},
		{
			problem: "12",
		},
		{
			problem: "13",
		},
		{
			problem: "14",
		},
		{
			problem: "15",
		},
		{
			problem: "20",
		},
		{
			problem: "21",
		},
		{
			problem: "22",
		},
		{
			problem: "23",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run(t, problemsDir + tt.problem, problemsDir + tt.problem + ".a")
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

	gotFile, err := os.CreateTemp(os.TempDir(), "content-delivery-")
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

