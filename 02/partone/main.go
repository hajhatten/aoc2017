package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		columns := strings.Split(scanner.Text(), "\t")

		highest, err := strconv.Atoi(columns[0])
		if err != nil {
			log.Fatal(err)
		}

		lowest, err := strconv.Atoi(columns[0])
		if err != nil {
			log.Fatal(err)
		}

		for _, c := range columns {
			n, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal(err)
			}
			if n > highest {
				highest = n
			}
			if n < lowest {
				lowest = n
			}
		}

		sum += (highest - lowest)
	}

	log.Println(sum)
}
