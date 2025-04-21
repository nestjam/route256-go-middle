package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var questionRegexp = regexp.MustCompile(`How old is (\w+)?`)
var constFactRegexp = regexp.MustCompile(`(\w+) is (\d+) years old`)
var sameFactRegexp = regexp.MustCompile(`(\w+) is the same age as (\w+)`)
var youngerFactRegexp = regexp.MustCompile(`(\w+) is (\d+) years younger than (\w+)`)
var olderFactRegexp = regexp.MustCompile(`(\w+) is (\d+) years older than (\w+)`)

func main() {
	var in *bufio.Reader
	in = bufio.NewReader(os.Stdin)
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setCount int

	fmt.Fscan(in, &setCount)
	fmt.Fscanln(in)

	var target string
	facts := make([]fact, 0)

	for {
		bytes, _, err := in.ReadLine()

		if err != nil {
			answer := solve(target, facts)
			fmt.Fprintln(out, answer)
			break
		}

		if isQuestion(bytes) == true {
			if len(facts) != 0 {
				answer := solve(target, facts)
				fmt.Fprintln(out, answer)
			}

			target = parseTarget(string(bytes))
			facts = facts[:0]
			continue
		}

		fact := parseFact(bytes)
		facts = append(facts, fact)
	}
}

func solve(target string, facts []fact) int {
	for i := 0; i < len(facts); i++ {
		if facts[i].hasOp(target) {
			if v, err := facts[i].calc(target, facts); err == nil {
				return v
			}
		}
	}

	panic("failed")
}

func readFacts(in *bufio.Reader) []fact {
	facts := make([]fact, 0)

	for {
		bytes, _, err := in.ReadLine()

		if isQuestion(bytes) == true {
			break
		}

		if err != nil {
			break
		}
		fact := parseFact(bytes)
		facts = append(facts, fact)
	}

	return facts
}

func parseFact(bytes []byte) fact {

	if constFactRegexp.Match(bytes) {
		matches := constFactRegexp.FindAllStringSubmatch(string(bytes), -1)
		v, _ := strconv.Atoi(matches[0][2])

		return constFact{
			op:    matches[0][1],
			value: v,
		}
	} else if sameFactRegexp.Match(bytes) {
		matches := sameFactRegexp.FindAllStringSubmatch(string(bytes), -1)

		return equalityFact{
			op:  matches[0][1],
			op2: matches[0][2],
		}
	} else if youngerFactRegexp.Match(bytes) {
		matches := youngerFactRegexp.FindAllStringSubmatch(string(bytes), -1)
		v2, _ := strconv.Atoi(matches[0][1])

		return youngerFact{
			op:    matches[0][1],
			op2:   matches[0][3],
			value: v2,
		}
	} else if olderFactRegexp.Match(bytes) {
		matches := olderFactRegexp.FindAllStringSubmatch(string(bytes), -1)
		v2, _ := strconv.Atoi(matches[0][1])

		return youngerFact{
			op:    matches[0][1],
			op2:   matches[0][3],
			value: v2,
		}
	}

	panic("failed parse fact")
}

func parseTarget(s string) string {
	matches := questionRegexp.FindAllStringSubmatch(s, -1)
	return matches[0][1]
}

func isQuestion(bytes []byte) bool {
	return questionRegexp.Match(bytes)
}

type fact interface {
	getId() int
	hasOp(op string) bool
	calc(op string, facts []fact) (int, error)
}

type olderFact struct {
	op    string
	op2   string
	id    int
	value int
}

func (f olderFact) hasOp(op string) bool {
	return f.op == op || f.op2 == op
}

func (f olderFact) calc(op string, facts []fact) (int, error) {
	if op != f.op && op != f.op2 {
		panic("wrong op")
	}

	for i := 0; i < len(facts); i++ {
		fact := facts[i]

		if fact.getId() == f.id || !fact.hasOp(f.op) {
			continue
		}

		if v, err := fact.calc(f.op, facts); err == nil {
			if op == f.op {
				return v, nil
			} else {
				return v - f.value, nil
			}
		}
	}

	for i := 0; i < len(facts); i++ {
		fact := facts[i]

		if fact.getId() == f.id || !fact.hasOp(f.op2) {
			continue
		}

		if v, err := fact.calc(f.op2, facts); err == nil {
			if op == f.op2 {
				return v, nil
			} else {
				return v + f.value, nil
			}
		}
	}

	return 0, errors.New("no result")
}

func (f olderFact) getId() int {
	return f.id
}

type youngerFact struct {
	op    string
	op2   string
	id    int
	value int
}

func (f youngerFact) hasOp(op string) bool {
	return f.op == op || f.op2 == op
}

func (f youngerFact) calc(op string, facts []fact) (int, error) {
	if op != f.op && op != f.op2 {
		panic("wrong op")
	}

	for i := 0; i < len(facts); i++ {
		fact := facts[i]

		if fact.getId() == f.id || !fact.hasOp(f.op) {
			continue
		}

		if v, err := fact.calc(f.op, facts); err == nil {
			if op == f.op {
				return v, nil
			} else {
				return v + f.value, nil
			}
		}
	}

	for i := 0; i < len(facts); i++ {
		fact := facts[i]

		if fact.getId() == f.id || !fact.hasOp(f.op2) {
			continue
		}

		if v, err := fact.calc(f.op2, facts); err == nil {
			if op == f.op2 {
				return v, nil
			} else {
				return v - f.value, nil
			}
		}
	}

	return 0, errors.New("no result")
}

func (f youngerFact) getId() int {
	return f.id
}

type constFact struct {
	op    string
	value int
	id    int
}

func (f constFact) hasOp(op string) bool {
	return f.op == op
}

func (f constFact) calc(op string, facts []fact) (int, error) {
	if op != f.op {
		panic("wrong op")
	}

	return f.value, nil
}

func (f constFact) getId() int {
	return f.id
}

type equalityFact struct {
	op  string
	op2 string
	id  int
}

func (f equalityFact) hasOp(op string) bool {
	return f.op == op || f.op2 == op
}

func (f equalityFact) calc(op string, facts []fact) (int, error) {
	if op != f.op && op != f.op2 {
		panic("wrong op")
	}

	for i := 0; i < len(facts); i++ {
		fact := facts[i]

		if fact.getId() == f.id || !fact.hasOp(f.op) {
			continue
		}

		if v, err := fact.calc(f.op, facts); err == nil {
			return v, nil
		}
	}

	for i := 0; i < len(facts); i++ {
		fact := facts[i]

		if fact.getId() == f.id || !fact.hasOp(f.op2) {
			continue
		}

		if v, err := fact.calc(f.op2, facts); err != nil {
			return v, nil
		}
	}

	return 0, errors.New("no result")
}

func (f equalityFact) getId() int {
	return f.id
}
