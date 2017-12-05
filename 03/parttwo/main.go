package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type position struct {
	x int
	y int
}

func main() {
	file, err := ioutil.ReadFile("../../input/03")
	if err != nil {
		log.Fatal(err)
	}

	n, err := strconv.Atoi(string(file))
	if err != nil {
		log.Fatal(err)
	}

	numbers := make(map[position]int)
	current := position{0, 0}
	direction := position{1, 0}
	numbers[current] = 1

	var sum int
	for {
		nextPosition := position{current.x + direction.x, current.y + direction.y}
		sum = calculateValueForPosition(nextPosition, numbers)
		numbers[nextPosition] = sum
		leftDirection, err := turnLeft(direction)
		if err != nil {
			log.Fatal(err)
		}
		leftPosition := position{nextPosition.x + leftDirection.x, nextPosition.y + leftDirection.y}
		if _, ok := numbers[leftPosition]; ok == false {
			direction = *leftDirection
		}
		current = nextPosition
		if sum > n {
			break
		}
	}

	log.Println(sum)
}

func calculateValueForPosition(p position, numbers map[position]int) int {
	var sum int
	diffs := []int{-1, 0, 1}
	for _, x := range diffs {
		for _, y := range diffs {
			if x == 0 && y == 0 {
				continue
			}
			if val, ok := numbers[position{x: p.x + x, y: p.y + y}]; ok {
				sum += val
			}
		}
	}
	return sum
}

func turnLeft(direction position) (*position, error) {
	if direction.x == 0 && direction.y == 1 {
		return &position{-1, 0}, nil
	} else if direction.x == -1 && direction.y == 0 {
		return &position{0, -1}, nil
	} else if direction.x == 0 && direction.y == -1 {
		return &position{1, 0}, nil
	} else if direction.x == 1 && direction.y == 0 {
		return &position{0, 1}, nil
	} else {
		return nil, fmt.Errorf("unknown direction")
	}
}
