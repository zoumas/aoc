package main

import (
	"fmt"
	"log"
	"os"

	solution "github.com/zoumas/aoc/2023/2"
)

func main() {
	data, err := os.ReadFile("2.in")
	if err != nil {
		log.Fatal(err)
	}
	res := solution.Solution2(string(data))
	fmt.Println(res)
}
