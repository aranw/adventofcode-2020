package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

func main() {
	c, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to read input.txt: %+v", err)
	}

	responses := strings.Split(string(c), "\n\n")

	fmt.Println(part1(responses))
	fmt.Println(part2(responses))
}

func part1(responses []string) int {
	score := 0

	for _, response := range responses {
		answers := make(map[rune]struct{})
		for _, c := range response {
			if unicode.IsLetter(c) {
				answers[c] = struct{}{}
			}

		}

		score += len(answers)
	}
	return score
}

func part2(responses []string) int {
	score := 0

	for _, response := range responses {
		answers := strings.Split(response, "\n")
		counts := make(map[rune]int)

		for _, a := range answers {
			for _, c := range a {
				if unicode.IsLetter(c) {
					if _, ok := counts[c]; !ok {
						counts[c] = 1
					} else {
						counts[c]++
					}
				}
			}
		}

		for _, c := range counts {
			if c == len(answers) {
				score++
			}
		}
	}

	return score
}
