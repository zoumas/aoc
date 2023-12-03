package solution

import "testing"

func TestSolution(t *testing.T) {
	given := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	want := 4361
	got := Solution(given)

	if got != want {
		t.Errorf("\ngot: %d\nwant: %d", got, want)
	}
}

func TestSolution2(t *testing.T) {
	given := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	want := 467835
	got := Solution2(given)
	if got != want {
		t.Errorf("\ngot: %d\nwant: %d", got, want)
	}
}
