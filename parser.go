package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type WordFreq struct {
	Word  string
	Count int
}

type MinHeap []WordFreq

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].Count < h[j].Count } // Min-heap
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(WordFreq)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

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

	h := &MinHeap{}
	heap.Init(h)

	for word, count := range wordCount {
		heap.Push(h, WordFreq{Word: word, Count: count})
		if h.Len() > 3 {
			heap.Pop(h) // Keep only top 3
		}
	}

	result := make([]WordFreq, h.Len())
	for i := len(result) - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(WordFreq)
	}

	fmt.Println("Top 3 most common words:")
	for _, wf := range result {
		fmt.Printf("%s: %d\n", wf.Word, wf.Count)
	}
	return nil
}
