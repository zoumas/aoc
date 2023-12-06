package solution

import (
	"strconv"
	"strings"
)

func unwrap[T any](t T, e error) T {
	if e != nil {
		panic(e)
	}
	return t
}

func Solution(body string) int {
	lines := strings.Split(body, "\n")
	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]

	p := 1
	for r := 0; r < len(times); r++ {
		var w int

		t := unwrap(strconv.Atoi(times[r]))
		d := unwrap(strconv.Atoi(distances[r]))

		for h := 1; h < t; h++ {
			rt := t - h
			dc := rt * h
			if dc > d {
				w++
			}
		}

		p *= w
	}

	return p
}

func Solution2(body string) int {
	lines := strings.Split(body, "\n")
	t := unwrap(strconv.Atoi(strings.Join(strings.Fields(lines[0])[1:], "")))
	d := unwrap(strconv.Atoi(strings.Join(strings.Fields(lines[1])[1:], "")))

	w := 0
	for h := 1; h < t; h++ {
		rt := t - h
		dc := rt * h
		if dc > d {
			w++
		}
	}

	return w
}
