package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../../input/05")
	if err != nil {
		log.Fatal(err)
	}

	var maze []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		maze = append(maze, n)
	}

	var jumps int
	var position int
	for position >= 0 && position < len(maze) {
		if maze[position] >= 3 {
			maze[position]--
			position = position + maze[position] + 1
		} else {
			maze[position]++
			position = position + maze[position] - 1
		}
		jumps++
		log.Println(jumps)
		log.Println(position)
	}

	log.Println(jumps)
}
