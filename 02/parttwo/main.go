package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../../input/02")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		columns := strings.Split(scanner.Text(), "\t")
		for i, x := range columns {
			xn, err := strconv.Atoi(x)
			if err != nil {
				log.Fatal(err)
			}
			for j, y := range columns {
				yn, err := strconv.Atoi(y)
				if err != nil {
					log.Fatal(err)
				}
				if i != j && xn%yn == 0 {
					sum += (xn / yn)
				}
			}
		}

	}

	log.Println(sum)
}
