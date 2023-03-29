package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_AtomicPush tests that PopFront operation is atomic.
func Test_AtomicPopFront(t *testing.T) {
	var wg sync.WaitGroup
	var list List
	var value int
	wg.Add(1)
	go func() {
		defer wg.Done()
		value = list.PopFront()
	}()
	list.PushFront(1)
	wg.Wait()
	assert.Equal(t, 1, value)
}
