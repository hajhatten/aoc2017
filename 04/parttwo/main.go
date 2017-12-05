package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

type anagramRuneSlice []rune

func (r anagramRuneSlice) Len() int           { return len(r) }
func (r anagramRuneSlice) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r anagramRuneSlice) Less(i, j int) bool { return r[i] < r[j] }

func main() {
	file, err := os.Open("../../input/04")
	if err != nil {
		log.Fatal(err)
	}

	var validPolicy1 [][]string

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
			validPolicy1 = append(validPolicy1, words)
		}
	}

	validPolicy2 := len(validPolicy1)

PhraseLoop:
	for _, phrase := range validPolicy1 {
		for x, w := range phrase {
			for y, w2 := range phrase {
				if x == y {
					continue
				}
				if len(w) == len(w2) {
					word1 := anagramRuneSlice(w)
					word2 := anagramRuneSlice(w2)
					sort.Sort(word1)
					sort.Sort(word2)

					if string(word1) == string(word2) {
						validPolicy2--
						continue PhraseLoop
					}
				}
			}
		}
	}

	log.Println(validPolicy2)
}
