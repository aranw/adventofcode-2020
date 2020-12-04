package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var policyRegex = regexp.MustCompile(`(\d+)-(\d+)\s(\w)`)

type Policy struct {
	Min    int
	Max    int
	Letter string
}

func (p *Policy) SledRentalValid(pass string) bool {
	if !strings.Contains(pass, p.Letter) {
		return false
	}

	count := strings.Count(pass, p.Letter)
	if !(count >= p.Min && count <= p.Max) {
		return false
	}

	return true
}

func (p *Policy) TobogganCorporateValid(pass string) bool {
	if !strings.Contains(pass, p.Letter) {
		return false
	}

	f := string(pass[p.Min-1])
	s := string(pass[p.Max-1])

	return (f == p.Letter) != (s == p.Letter)
}

func NewPolicy(l string) *Policy {
	p := policyRegex.FindStringSubmatch(l)

	min, _ := strconv.Atoi(p[1])
	max, _ := strconv.Atoi(p[2])
	letter := p[3]

	return &Policy{
		Min:    min,
		Max:    max,
		Letter: letter,
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to read input.txt: %+v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sledRentalValid := 0
	tobogganCorporateValid := 0

	for scanner.Scan() {
		l := strings.Split(scanner.Text(), ":")

		pass := strings.TrimSpace(l[1])

		p := NewPolicy(l[0])
		if p.SledRentalValid(pass) {
			sledRentalValid++
		}
		if p.TobogganCorporateValid(pass) {
			tobogganCorporateValid++
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sledRentalValid)
	fmt.Println(tobogganCorporateValid)
}
