package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	w := newWords()

	filenames := os.Args[1:]

	for _, f := range filenames {
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()

	fmt.Println("Words that appear more than once.")
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%d %s\n", count, word)
		}
	}
}

type words struct {
	sync.Mutex
	found map[string]int
}

// newWords returns the pointer to a newly created data structure words.
func newWords() *words {
	return &words{found: map[string]int{}}
}

// add is a method on words, that adds the number n to the tally.
func (w *words) add(word string, n int) {
	w.Lock()
	w.found[word] += n
	w.Unlock()
}

// tallyWords gets the number
func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}
