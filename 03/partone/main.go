package main

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
)

func main() {
	file, err := ioutil.ReadFile("../../input/03")
	if err != nil {
		log.Fatal(err)
	}

	n, err := strconv.Atoi(string(file))
	if err != nil {
		log.Fatal(err)
	}

	var steps int
	var turns int
	var posX float64
	var posY float64

	for steps < n-1 {
		length := (turns / 2) + 1
		for i := 0; i < length; i++ {
			if steps == n-1 {
				break
			}
			steps++
			direction := turns % 4
			switch direction {
			case 0:
				posX++
			case 1:
				posY++
			case 2:
				posX--
			default:
				posY--
			}
		}
		turns++
	}
	log.Println(math.Abs(posX) + math.Abs(posY))
}
