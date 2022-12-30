package dayOne

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.in
var input string

func DayOnePartOne() {
	elves := ParseInput(input)

	max := 0
	for _, value := range elves {
		weight := value.weight
		if weight > max {
			max = weight
		}
	}

	fmt.Printf("Max calories is %d \r\n", max)
}

func DayOnePartTwo() {
	elves := ParseInput(input)

	weights := []int{}

	for _, elf := range elves {
		weights = append(weights, elf.weight)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(weights)))

	maxWeights := 0

	for i := 0; i < 3; i++ {
		maxWeights += weights[i]
	}

	fmt.Printf("Sum of top 3 elves: %d", maxWeights)
}

func main() {
	DayOnePartOne()
	DayOnePartTwo()
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

func ParseInput(input string) []elf {
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

	return elves
}
