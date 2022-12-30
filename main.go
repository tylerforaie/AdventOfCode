package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed dayOne/input.in
var input string

func DayOnePartOne() {
	elves := make([]elf, 0)
	items := make([]int, 0)

	input = strings.TrimRight(input, "\r\n")
	for _, value := range strings.Split(input, "\r\n") {
		if value != "" {
			convertedValue, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			items = append(items, convertedValue)
		} else {
			elf := NewElf(items)
			elves = append(elves, *elf)
			items = nil
		}
	}

	max := 0
	for _, value := range elves {
		weight := value.weight
		if weight > max {
			max = weight
		}
	}

	fmt.Printf("Max calories is %d", max)
}

func main() {
	DayOnePartOne()
}

type elf struct {
	weight int
	items  []int
}

func (e elf) calculateWeight() int {
	sum := 0
	for _, value := range e.items {
		sum += value
	}

	return sum
}

func NewElf(itemList []int) *elf {
	e := new(elf)
	e.items = itemList
	e.weight = e.calculateWeight()

	return e
}
