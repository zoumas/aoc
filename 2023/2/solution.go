package solution

import (
	"log"
	"strconv"
	"strings"
)

func Solution1(body string) int {
	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	var idSum int

	for i, game := range strings.Split(body, "\n") {
		if game == "" {
			continue
		}
		validGame := true

		gameTagAndData := strings.Split(game, ":")
		if len(gameTagAndData) != 2 {
			log.Print(i, gameTagAndData)
			log.Fatalf("Bad input: %v\n", gameTagAndData)
		}
		tag := gameTagAndData[0]
		idString := (strings.Fields(tag))[1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			log.Fatal(err)
		}

		data := gameTagAndData[1]
		recordSet := strings.Split(data, ";")
	RecordLoop:
		for _, record := range recordSet {
			for _, pair := range strings.Split(record, ",") {
				fields := strings.Fields(pair)
				quantity, err := strconv.Atoi(fields[0])
				if err != nil {
					log.Fatal(err)
				}
				color := fields[1]

				if quantity > maxCubes[color] {
					validGame = false
					break RecordLoop
				}
			}
		}

		if validGame {
			idSum += id
		}
	}

	return idSum
}

func Solution2(body string) int {
	var powerSum int

	for i, game := range strings.Split(body, "\n") {
		if game == "" {
			continue
		}
		validGame := true

		gameTagAndData := strings.Split(game, ":")
		if len(gameTagAndData) != 2 {
			log.Print(i, gameTagAndData)
			log.Fatalf("Bad input: %v\n", gameTagAndData)
		}

		data := gameTagAndData[1]
		maxCubesSeen := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}

		for _, record := range strings.Split(data, ";") {
			for _, pair := range strings.Split(record, ",") {
				fields := strings.Fields(pair)
				quantity, err := strconv.Atoi(fields[0])
				if err != nil {
					log.Fatal(err)
				}
				color := fields[1]

				if quantity > maxCubesSeen[color] {
					maxCubesSeen[color] = quantity
				}
			}
		}

		if validGame {
			powerSum += maxCubesSeen["red"] * maxCubesSeen["blue"] * maxCubesSeen["green"]
		}
	}

	return powerSum
}
