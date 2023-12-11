package solution

import (
	"fmt"
	"strings"
)

type Pos struct {
	I, J int
}

func Solution1(body string) int {
	m := extractImage(body)
	m = expandUniverse(m)

	galaxies := []Pos{}
	count := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == '#' {
				count++
				// m[i][j] = rune(count + '0')
				galaxies = append(galaxies, Pos{I: i, J: j})
			}
		}
	}

	for _, r := range m {
		fmt.Println(string(r))
	}

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			s := galaxies[i]
			t := galaxies[j]
			dist := ShortestPath(s, t, m)
			fmt.Println("Start at", i+1, "End at", j+1, dist)
			sum += dist
		}
	}

	return sum
}

func ShortestPath(s, t Pos, m [][]rune) int {
	visited := make(map[Pos]struct{})
	toVisit := []Pos{s}
	parent := make(map[Pos]*Pos)

	visited[s] = struct{}{}
	parent[s] = &s

	done := false
	for len(toVisit) > 0 && !done {
		v := toVisit[0]
		toVisit = toVisit[1:]

		for _, n := range Neighbors(v, m) {
			if _, ok := visited[n]; !ok {
				visited[n] = struct{}{}
				parent[n] = &v
				toVisit = append(toVisit, n)
				if t == v {
					done = true
					break
				}
			}
		}
	}

	path := []Pos{}
	v := t
	for parent[v] != &s {
		path = append([]Pos{v}, path...)
		v = *parent[v]
	}
	path = append([]Pos{s}, path...)

	return len(path) - 1
}

func Neighbors(p Pos, m [][]rune) []Pos {
	final := make([]Pos, 0, 4)
	for _, p := range [4]Pos{
		{I: p.I - 1, J: p.J}, // top
		{I: p.I + 1, J: p.J}, // bottom
		{I: p.I, J: p.J - 1}, // left
		{I: p.I, J: p.J + 1}, // right
	} {
		if isOutOfBounds(p, len(m), len(m[0])) {
			continue
		}

		final = append(final, p)
	}
	return final
}

func isOutOfBounds(p Pos, n, m int) bool {
	return !(p.I >= 0 && p.I < n && p.J >= 0 && p.J < m)
}

func extractImage(body string) [][]rune {
	lines := strings.Split(body, "\n")
	m := make([][]rune, 0, len(lines))
	for _, line := range lines {
		row := make([]rune, 0, len(line))
		for _, c := range line {
			row = append(row, c)
		}
		m = append(m, row)
	}
	return m
}

func expandUniverse(m [][]rune) [][]rune {
	rows := []int{}
	cols := []int{}

	for i, r := range m {
		if !strings.Contains(string(r), "#") {
			rows = append(rows, i)
		}
	}

	for j := 0; j < len(m[0]); j++ {
		line := make([]rune, 0, len(m))
		for i := 0; i < len(m); i++ {
			line = append(line, m[i][j])
		}
		if !strings.Contains(string(line), "#") {
			cols = append(cols, j)
		}
	}

	expRow := make([]rune, len(m[0]))
	for i := 0; i < len(expRow); i++ {
		expRow[i] = '.'
	}

	expansions := 0
	for _, r := range rows {
		before := m[:r+expansions]
		after := m[r+expansions:]

		m = [][]rune{}
		m = append(m, before...)
		m = append(m, expRow)
		m = append(m, after...)

		expansions++
	}

	expCol := make([]rune, len(m))
	for i := 0; i < len(expCol); i++ {
		expCol[i] = '.'
	}

	expansions = 0
	for _, c := range cols {
		for i, r := range m {
			before := r[:c+expansions]
			after := r[c+expansions:]

			exp := []rune{}
			exp = append(exp, before...)
			exp = append(exp, '.')
			exp = append(exp, after...)
			m[i] = exp
		}
		expansions++
	}

	return m
}
