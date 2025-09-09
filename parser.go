package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Top3Words(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	wordCount := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			word = strings.ToLower(strings.Trim(word, ".,!?:;\"'()[]{}"))
			if word != "" {
				wordCount[word]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	type kv struct {
		Word  string
		Count int
	}
	var wordList []kv
	for w, c := range wordCount {
		wordList = append(wordList, kv{w, c})
	}

	sort.Slice(wordList, func(i, j int) bool {
		return wordList[i].Count > wordList[j].Count
	})

	fmt.Println("Top 3 most common words:")
	for i := 0; i < len(wordList) && i < 3; i++ {
		fmt.Printf("%s: %d\n", wordList[i].Word, wordList[i].Count)
	}

	return nil
}


