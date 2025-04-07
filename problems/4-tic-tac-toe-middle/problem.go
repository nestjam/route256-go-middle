package main

type cell = byte

const (
	e cell = iota
	o
	x
)

type board struct {
	b    [][]cell
	n, m int
}

func newBoard(b [][]cell) board {
	return board{b: b, n: len(b), m: len(b[0])}
}

func (b *board) lookupMajorDiagonal(i, j, maxCount int) (line, bool) {
	line := line{start: newPoint(i, j)}

	l, k := i, j
	for l < b.n && k < b.m {
		if b.b[l][k] != b.b[i][j] {
			break
		}

		line.count++
		l++
		k++

		if line.count >= maxCount {
			return line, false
		}
	}

	line.end = newPoint(l-1, k-1)
	return line, true
}

func (b *board) lookupMinorDiagonal(i, j, maxCount int) (line, bool) {
	line := line{start: newPoint(i, j)}

	l, k := i, j
	for l < b.n && k >= 0 {
		if b.b[l][k] != b.b[i][j] {
			break
		}

		line.count++
		l++
		k--

		if line.count >= maxCount {
			return line, false
		}
	}

	line.end = newPoint(l-1, k+1)
	return line, true
}

func (b *board) lookupRight(i, j, maxCount int) (line, bool) {
	line := line{start: newPoint(i, j)}

	l := j
	for ; l < b.m; l++ {
		if b.b[i][l] != b.b[i][j] {
			break
		}

		line.count++

		if line.count >= maxCount {
			return line, false
		}
	}

	line.end = newPoint(i, l-1)
	return line, true
}

func (b *board) lookupDown(i, j, maxCount int) (line, bool) {
	line := line{start: newPoint(i, j)}

	l := i
	for ; l < b.n; l++ {
		if b.b[l][j] != b.b[i][j] {
			break
		}

		line.count++

		if line.count >= maxCount {
			return line, false
		}
	}

	line.end = newPoint(l-1, j)
	return line, true
}

func (b *board) isEmpty(p point) bool {
	if p.i < 0 || p.i >= b.n || p.j < 0 || p.j >= b.m {
		return false
	}

	return b.b[p.i][p.j] == e
}

func (b *board) isCross(p point) bool {
	if p.i < 0 || p.i >= b.n || p.j < 0 || p.j >= b.m {
		return false
	}

	return b.b[p.i][p.j] == x
}

type line struct {
	start, end point
	count      int
}

type point struct {
	i, j int
}

func newPoint(i, j int) point {
	return point{i, j}
}

func (p point) add(i, j int) point {
	return newPoint(p.i+i, p.j+j)
}

func (b *board) cell(i, j int) cell {
	return b.b[i][j]
}

func canCrossWin(b *board, k int) bool {
	horizontalLines, ok := findHorizontalLines(b, k)
	if !ok {
		return false
	}

	verticalLines, ok := findVerticalLines(b, k)
	if !ok {
		return false
	}

	majorDiagonalLines, ok := findMajorDiagonalLines(b, k)
	if !ok {
		return false
	}

	minorDiagonalLines, ok := findMinorDiagonalLines(b, k)
	if !ok {
		return false
	}

	if k == 1 {
		for i := 0; i < b.n; i++ {
			for j := 0; j < b.m; j++ {
				if b.cell(i, j) == e {
					return true
				}
			}
		}
	} else {
		if canCrossWinInHorizontalLines(horizontalLines, b, k) {
			return true
		}

		if canCrossWinInVerticalLines(verticalLines, b, k) {
			return true
		}

		if canCrossWinInMajorDiagonalLines(majorDiagonalLines, b, k) {
			return true
		}

		if canCrossWinInMinorDiagonalLines(minorDiagonalLines, b, k) {
			return true
		}
	}

	return false
}

func canCrossWinInMinorDiagonalLines(lines map[point]line, b *board, k int) bool {
	for _, line := range lines {
		if line.count == k-1 && b.isEmpty(line.start.add(-1, 1)) {
			return true
		}

		count := line.count
		if b.isEmpty(line.end.add(1, -1)) {
			count++

			if nextLine, ok := lines[line.end.add(2, -2)]; ok {
				count += nextLine.count
			}

			if count >= k {
				return true
			}
		}
	}

	return false
}

func findMinorDiagonalLines(b *board, k int) (lines map[point]line, ok bool) {
	lines = make(map[point]line)

	for i := 0; i < b.n-1; i++ {
		l, j := i, b.m-1
		for j >= 0 && l < b.n {
			cell := b.cell(l, j)

			if cell == e {
				l++
				j--
				continue
			}

			line, ok := b.lookupMinorDiagonal(l, j, k)

			if !ok {
				return lines, false
			}

			if cell == x {
				lines[line.start] = line
			}

			l += line.count
			j -= line.count
		}
	}

	for j := b.m - 2; j > 0; j-- {
		i, l := 0, j
		for l >= 0 && i < b.n {
			cell := b.cell(i, l)

			if cell == e {
				i++
				l--
				continue
			}

			line, ok := b.lookupMinorDiagonal(i, l, k)

			if !ok {
				return lines, false
			}

			if cell == x {
				lines[line.start] = line
			}

			i += line.count
			l -= line.count
		}
	}

	return lines, true
}

func canCrossWinInMajorDiagonalLines(lines map[point]line, b *board, k int) bool {
	for _, line := range lines {
		if line.count == k-1 && b.isEmpty(line.start.add(-1, -1)) {
			return true
		}

		count := line.count
		if b.isEmpty(line.end.add(1, 1)) {
			count++

			if nextLine, ok := lines[line.end.add(2, 2)]; ok {
				count += nextLine.count
			}

			if count >= k {
				return true
			}
		}
	}

	return false
}

func findMajorDiagonalLines(b *board, k int) (lines map[point]line, ok bool) {
	lines = make(map[point]line)

	for i := 0; i < b.n-1; i++ {
		l, j := i, 0
		for j < b.m && l < b.n {
			cell := b.cell(l, j)

			if cell == e {
				l++
				j++
				continue
			}

			line, ok := b.lookupMajorDiagonal(l, j, k)

			if !ok {
				return lines, false
			}

			if cell == x {
				lines[line.start] = line
			}

			l += line.count
			j += line.count
		}
	}

	for j := 1; j < b.m-1; j++ {
		i, l := 0, j
		for l < b.m && i < b.n {
			cell := b.cell(i, l)

			if cell == e {
				i++
				l++
				continue
			}

			line, ok := b.lookupMajorDiagonal(i, l, k)

			if !ok {
				return lines, false
			}

			if cell == x {
				lines[line.start] = line
			}

			i += line.count
			l += line.count
		}
	}

	return lines, true
}

func canCrossWinInVerticalLines(lines map[point]line, b *board, k int) bool {
	for _, line := range lines {
		if line.count == k-1 && b.isEmpty(line.start.add(-1, 0)) {
			return true
		}

		count := line.count
		if b.isEmpty(line.end.add(1, 0)) {
			count++

			if nextLine, ok := lines[line.end.add(2, 0)]; ok {
				count += nextLine.count
			}

			if count >= k {
				return true
			}
		}
	}

	return false
}

func findVerticalLines(b *board, k int) (lines map[point]line, ok bool) {
	lines = make(map[point]line)

	for j := 0; j < b.m; j++ {
		count := 0

		for i := 0; i < b.n; i++ {
			cell := b.cell(i, j)

			if cell == e {
				count++

				if count > b.n-k && !b.isCross(newPoint(i+1, j)) {
					break
				}

				continue
			}

			line, ok := b.lookupDown(i, j, k)

			if !ok {
				return lines, false
			}

			if cell == x {
				lines[line.start] = line
			}

			i += line.count - 1
			count += line.count
		}
	}

	return lines, true
}

func canCrossWinInHorizontalLines(lines map[point]line, b *board, k int) bool {
	for _, line := range lines {
		if line.count == k-1 && b.isEmpty(line.start.add(0, -1)) {
			return true
		}

		count := line.count
		if b.isEmpty(line.end.add(0, 1)) {
			count++

			if nextLine, ok := lines[line.end.add(0, 2)]; ok {
				count += nextLine.count
			}

			if count >= k {
				return true
			}
		}
	}

	return false
}

func findHorizontalLines(b *board, k int) (lines map[point]line, ok bool) {
	lines = make(map[point]line)

	for i := 0; i < b.n; i++ {
		count := 0

		for j := 0; j < b.m; j++ {
			cell := b.cell(i, j)

			if cell == e {
				count++

				if count > b.m-k && !b.isCross(newPoint(i, j+1)) {
					break
				}

				continue
			}

			line, ok := b.lookupRight(i, j, k)

			if !ok {
				return lines, false
			}

			if b.cell(i, j) == x {
				lines[line.start] = line
			}

			j += line.count - 1
			count += line.count
		}
	}

	return lines, true
}
