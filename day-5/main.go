package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to read input.txt: %+v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	seatID := 0

	seats := make(map[int]struct{})

	for scanner.Scan() {
		l := scanner.Text()

		lowRow := 0
		highRow := 127
		rowNum := 0
		row := l[:7]

		for _, r := range row {
			if r == 'B' {
				lowRow = lowRow + ((highRow-lowRow)/2 + 1)
				rowNum = highRow
			} else {
				highRow = highRow - ((highRow-lowRow)/2 + 1)
				rowNum = lowRow
			}
		}

		lowColumn := 0
		highColumn := 7
		column := l[7:]
		coulmnNum := 0

		for _, c := range column {
			if c == 'R' {
				lowColumn = lowColumn + ((highColumn-lowColumn)/2 + 1)
				coulmnNum = highColumn
			} else {
				highColumn = highColumn - ((highColumn-lowColumn)/2 + 1)
				coulmnNum = lowColumn
			}
		}

		currentSeatID := rowNum*8 + coulmnNum

		if currentSeatID > seatID {
			seatID = currentSeatID
		}
		seats[currentSeatID] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(seatID)

	for i := 0; i < seatID; i++ {
		if _, ok := seats[i]; !ok {
			fmt.Println(i)
		}
	}
}
