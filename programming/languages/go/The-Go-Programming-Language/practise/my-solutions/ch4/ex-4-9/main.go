package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	counts := make(map[string]int)

	f, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		counts[input.Text()]++
	}

	// sort by key
	words := make([]string, 0, len(counts))
	for word := range counts {
		words = append(words, word)
	}
	sort.Strings(words)

	for _, word := range words {
		fmt.Printf("%s:\t%d\n", word, counts[word])
	}
}
