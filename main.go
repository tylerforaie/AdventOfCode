package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func DayOnePartOne() {
	file, err := os.Open("dayOne/input.in")
	elves := make([]elf, 0)
	items := make([]int, 0)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			elf := NewElf(items)
			elves = append(elves, *elf)
			items = nil
			continue
		} else {
			value := scanner.Text()
			convertedValue, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			items = append(items, convertedValue)
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
