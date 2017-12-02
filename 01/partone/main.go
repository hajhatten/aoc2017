package main

import (
	"io/ioutil"
	"log"
)

func main() {
	file, err := ioutil.ReadFile("../../input/01")
	if err != nil {
		log.Fatal(err)
	}

	input := string(file)
	offset := 1

	var sum int

	for i, r := range input {
		n := int(r - '0')
		n2 := int(input[(i+offset)%len(input)] - '0')
		if n == n2 {
			sum += n
		}
	}
	log.Println(sum)
}
