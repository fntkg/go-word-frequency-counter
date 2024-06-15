# Word Frequency Counter in Go: Documentation

## Overview

This documentation explains how a word frequency counter was implemented in Go, equivalent to the following bash pipeline:

```bash
tr -s ' ' '\n' | sort | uniq -c | sort -n | tail
```

The Go script processes text input to count the frequency of each word and outputs the top most frequent words.

## Usage

To use the script, run the following command:

```bash
cat /path/to/file.txt | go run word-frequency-counter.go
```

This command reads the contents of `file.txt`, processes it with the Go script, and outputs the top 10 most frequent words.

## Pipeline Description

The bash pipeline performs these steps:

1. **`tr -s ' ' '\n'`**: Changes spaces to newline characters to split the text into words.
2. **`sort`**: Sorts the words alphabetically.
3. **`uniq -c`**: Counts the occurrences of each unique word.
4. **`sort -n`**: Sorts the counts numerically.
5. **`tail`**: Shows the last 10 lines, which are the 10 most frequent words.

## Go Implementation

### Input Handling

The script reads from `stdin` using the `bufio` package, which allows handling large text inputs efficiently.

### Word Counting

A map is used to store how many times each word appears. All words are converted to lowercase to avoid case-sensitive duplicates.

### Efficient Sorting with Min Heap

Sorting all words and their counts would be slow for large inputs. Instead, we use a fixed-size min heap (size 10) to keep track of the top 10 most frequent words. This keeps memory usage constant, no matter how large the input.

### Min Heap Implementation

We use Go's `container/heap` package to manage the min heap. The algorithm processes the word frequency map, inserting each word into the heap if its count is greater than the smallest count in the heap. If the heap already has 10 elements, the smallest one is removed.

### Output

After processing, the heap is used to retrieve the top 10 words in order of frequency, similar to the final sort and `tail` in the bash pipeline.

## Further Optimizations

### Parallelization

To handle even larger inputs, the script could use Goroutines and channels to process text in parallel. This would split the text file into chunks, each handled by a separate Goroutine, improving performance on multi-core systems.

However, this requires careful handling of concurrency and synchronization and is beyond the current scope. Future work could explore this for further optimization.

## Conclusion

This Go script efficiently counts word frequencies and finds the top 10 most frequent words, closely mimicking the given bash pipeline. By using a min heap, the solution remains fast even for large inputs, showing an effective way to handle large-scale text processing in Go.