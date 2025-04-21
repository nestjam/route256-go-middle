package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	var in *bufio.Reader
	in = bufio.NewReader(os.Stdin)
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setCount int
	var statementCount int

	fmt.Fscan(in, &setCount)

	for i := 0; i < setCount; i++ {
		fmt.Fscan(in, &statementCount)
		fmt.Fscanln(in)

		statements := readStatements(in, statementCount)

		suspects := judge(statements)

		writeStatements(out, suspects, statements[0].action)
	}
}

func judge(statements []statement) []string {
	ratings := make(map[string]int)

	for i := 0; i < len(statements); i++ {
		s := statements[i]
		ratings[s.object] = 0
		ratings[s.subject] = 0
	}

	for i := 0; i < len(statements); i++ {
		s := statements[i]

		if s.isAboutSelf {
			if s.isPositiveRelated {
				ratings[s.subject] += 2
			} else {
				ratings[s.subject] -= 1
			}

			continue
		}

		if s.isPositiveRelated {
			ratings[s.object] += 1
		} else {
			ratings[s.object] -= 1
		}
	}

	suspects := make([]string, 0)
	max := math.MinInt

	for person, rating := range ratings {
		if rating > max {
			suspects = suspects[:0]
			suspects = append(suspects, person)
			max = rating
		} else if rating == max {
			suspects = append(suspects, person)
		}
	}

	slices.SortFunc(suspects, strings.Compare)

	return suspects
}

func writeStatements(out *bufio.Writer, suspects []string, action string) {
	for i := 0; i < len(suspects); i++ {
		fmt.Fprintf(out, "%s is %s.", suspects[i], action)
		fmt.Fprintln(out)
	}
}

func readStatements(in *bufio.Reader, count int) []statement {
	statements := make([]statement, count)
	for i := 0; i < count; i++ {
		bytes, _, _ := in.ReadLine()
		statements[i] = parseStatement(string(bytes))
	}
	return statements
}

func parseStatement(value string) statement {
	words := strings.Split(value, " ")
	statement := statement{}

	statement.subject = words[0]
	statement.subject = statement.subject[:len(statement.subject)-1]
	statement.object = words[1]

	statement.isPositiveRelated = true
	if words[2] == "am" {
		statement.object = statement.subject
	}

	if words[3] == "not" {
		statement.isPositiveRelated = false
	}

	statement.action = words[len(words)-1]
	statement.action = statement.action[:len(statement.action)-1]

	if statement.object == statement.subject {
		statement.isAboutSelf = true
	}

	return statement
}

type statement struct {
	subject           string
	object            string
	action            string
	isPositiveRelated bool
	isAboutSelf       bool
}
