package solution_test

import (
	"testing"

	solution "github.com/zoumas/aoc/2023/11"
)

// func TestSolution1(t *testing.T) {
// 	given := `...#......
// .......#..
// #.........
// ..........
// ......#...
// .#........
// .........#
// ..........
// .......#..
// #...#.....`
//
// 	want := 374
//
// 	got := solution.Solution1(given)
//
// 	if got != want {
// 		t.Errorf("\ngot: %d\nwant: %d", got, want)
// 	}
// }

func TestSolution2(t *testing.T) {
	given := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	want := 8410

	got := solution.Solution2(given, 100-1)

	if got != want {
		t.Errorf("\ngot: %d\nwant: %d", got, want)
	}
}
