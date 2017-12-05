package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../../input/04")
	if err != nil {
		log.Fatal(err)
	}
	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		var containsDup bool

		for i, w := range words {
			for i2, w2 := range words {
				if i == i2 {
					continue
				}

				if w == w2 {
					containsDup = true
				}
			}
		}

		if !containsDup {
			sum++
		}
	}
	log.Println(sum)
}
