package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

const MaxHeapSize = 10

// An Item is something we manage in a priority queue.
type Item struct {
	word  string // The word of the item.
	count int    // The count (priority) of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not the highest count, so we use less than here.
	return pq[i].count < pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	//n := len(*pq)
	item := x.(*Item)
	//item.index = n
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the count and word of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, word string, count int) {
	item.word = word
	item.count = count
	heap.Fix(pq, item.index)
}

func main() {
	// Scan file from stdin
	items := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Read and count words
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		items[word]++
	}

	// Create a count queue, put the items in it, and
	// establish the count queue (heap) invariants.
	//pq := make(PriorityQueue, len(items))
	pq := make(PriorityQueue, 0, MaxHeapSize)
	heap.Init(&pq)
	i := 0
	for word, count := range items {
		item := &Item{
			word:  word,
			count: count,
			index: i,
		}
		i++
		// Add data if heap not full or new item has higher count
		if i <= MaxHeapSize || count > pq[0].count {
			if i > MaxHeapSize {
				heap.Pop(&pq)
			}
			heap.Push(&pq, item)
		}

	}

	// Take the items out; they arrive in increasing count order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%d %s \n", item.count, item.word)
	}
}
