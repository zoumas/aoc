package solution_test

import (
	"testing"

	solution "github.com/zoumas/aoc/2023/6"
)

func TestSolution(t *testing.T) {
	given := `Time:      7  15   30
Distance:  9  40  200`

	got := solution.Solution(given)
	want := 288

	if got != want {
		t.Errorf("\ngot: %d\nwant: %d", got, want)
	}
}

func TestSolution2(t *testing.T) {
	given := `Time:      7  15   30
Distance:  9  40  200`
	got := solution.Solution2(given)
	want := 71503

	if got != want {
		t.Errorf("\ngot: %d\nwant: %d", got, want)
	}
}

func BenchmarkSolution2(b *testing.B) {
	given := `Time:        55     99     97     93
Distance:   401   1485   2274   1405`

	for i := 0; i < b.N; i++ {
		got := solution.Solution2(given)
		want := 39132886
		if got != want {
			b.Errorf("\ngot: %d\nwant: %d", got, want)
		}
	}
}
