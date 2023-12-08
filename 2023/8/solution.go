package solution

import (
	"fmt"
	"log"
	"strings"
)

type I struct {
	D map[byte]string
}

func Solution1(body string) int {
	lines := strings.Split(body, "\n")
	instructions := lines[0]
	fmt.Println(instructions)

	m := map[string]I{}

	lines = lines[2:]
	for _, line := range lines {
		ls := strings.Split(line, " = ")
		src := ls[0]

		dsts := ls[1]
		dst := strings.Split(dsts, ", ")
		left := dst[0][1:]
		right := dst[1][:len(dst[1])-1]

		v := map[byte]string{}
		v['L'] = left
		v['R'] = right
		m[src] = I{D: v}
	}

	for k, v := range m {
		log.Println(k, v)
	}

	from := "AAA"
	to := "ZZZ"

	count := 0
	for {
		for _, c := range instructions {
			d := byte(c)
			v, ok := m[from]
			if !ok {
				log.Fatal("Value does not exist", from)
			}

			log.Println(v.D[d])

			count++
			if v.D[d] == to {
				log.Println("Found", to, "after ", count, "steps")
				return count
			}
			from = v.D[d]
		}
	}
}

func Solution2(body string) int {
	lines := strings.Split(body, "\n")
	directions := lines[0]
	fmt.Println(directions)

	m := map[string]I{}

	srcs := []string{}

	lines = lines[2:]
	for _, line := range lines {
		ls := strings.Split(line, " = ")
		src := ls[0]
		if src[len(src)-1] == 'A' {
			srcs = append(srcs, src)
		}

		dsts := ls[1]
		dst := strings.Split(dsts, ", ")
		left := dst[0][1:]
		right := dst[1][:len(dst[1])-1]

		v := map[byte]string{}
		v['L'] = left
		v['R'] = right
		m[src] = I{D: v}
	}

	log.Println("Starts", srcs)
	for k, v := range m {
		log.Println(k, v)
	}

	all := len(srcs)
	steps := make([]int, 0, all)
	for _, src := range srcs {
		count := 0
	L:
		for {
			for i := range directions {
				d := directions[i]
				v, ok := m[src]
				if !ok {
					log.Fatal(src, "Dead End")
				}

				dest := v.D[d]
				count++
				last := dest[len(dest)-1]
				if last == 'Z' {
					log.Println(src, "reached", dest)
					steps = append(steps, count)
					break L
				}
				src = v.D[d]
			}
		}
	}

	log.Println(steps)
	return LCM(steps)
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(ns []int) int {
	r := ns[0]
	for i := 1; i < len(ns); i++ {
		r = ns[i] * r / GCD(ns[i], r)
	}
	return r
}
