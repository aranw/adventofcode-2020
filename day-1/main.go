package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to read input.txt: %+v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var nums = make([]int, 0)
processor:
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Printf("failed to conver integer to string: %+v", err)
			continue
		}

		nums = append(nums, n)
		for _, a := range nums {
			if n+a == 2020 {
				fmt.Printf("%d * %d = %d\n", n, a, n*a)

			}

			for _, b := range nums {
				if n+a+b == 2020 {
					fmt.Printf("%d * %d * %d = %d\n", n, a, b, n*a*b)
					break processor
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
