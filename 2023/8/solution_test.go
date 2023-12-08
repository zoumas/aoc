package solution_test

import (
	"testing"

	solution "github.com/zoumas/aoc/2023/8"
)

// func TestSolution1(t *testing.T) {
// 	t.Run("base case", func(t *testing.T) {
// 		body := `RL
//
// AAA = (BBB, CCC)
// BBB = (DDD, EEE)
// CCC = (ZZZ, GGG)
// DDD = (DDD, DDD)
// EEE = (EEE, EEE)
// GGG = (GGG, GGG)
// ZZZ = (ZZZ, ZZZ)`
//
// 		want := 2
// 		got := solution.Solution1(body)
// 		if want != got {
// 			t.Errorf("\ngot: %d\nwant: %d", got, want)
// 		}
// 	})
//
// 	t.Run("repeating", func(t *testing.T) {
// 		body := `LLR
//
// AAA = (BBB, BBB)
// BBB = (AAA, ZZZ)
// ZZZ = (ZZZ, ZZZ)`
//
// 		want := 6
// 		got := solution.Solution1(body)
// 		if want != got {
// 			t.Errorf("\ngot: %d\nwant: %d", got, want)
// 		}
// 	})
// }

func TestSolution2(t *testing.T) {
	body := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

	want := 6
	got := solution.Solution2(body)
	if want != got {
		t.Errorf("\ngot: %d\nwant: %d", got, want)
	}
}
