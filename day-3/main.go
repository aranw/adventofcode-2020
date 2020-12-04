package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	c, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to read input.txt: %+v", err)
	}

	slopeMap := strings.Split(string(c), "\n")

	hits1 := checkSlope(slopeMap, 3, 1)
	hits2 := checkSlope(slopeMap, 1, 1)
	hits3 := checkSlope(slopeMap, 5, 1)
	hits4 := checkSlope(slopeMap, 7, 1)
	hits5 := checkSlope(slopeMap, 1, 2)

	fmt.Println(hits1)
	fmt.Println(hits2 * hits1 * hits3 * hits4 * hits5)
}

func checkSlope(slope []string, right, down int) int {
	trees := 0

	for x, y := 0, 0; y < len(slope); x, y = x+right, y+down {
		if string(slope[y][x%len(slope[y])]) == "#" {
			trees++
		}
	}

	return trees
}
