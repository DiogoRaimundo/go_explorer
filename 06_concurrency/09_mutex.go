/*
Channels are great for communication among goroutines.
If we just need to make sure only one goroutine can access a variable at a time to avoid conflicts.
For that, we use Go's standard library sync.Mutex.
*/
package goTour06

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	// Uncomment to cause more race conditions
	// time.Sleep(time.Second)

	// Lock so only one goroutine at a time can access the map c.v.
	c.mu.Lock()

	c.v[key]++

	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	// Lock so only one goroutine at a time can access the map c.v.
	c.mu.Lock()

	//We use defer to ensure the mutex will be unlocked.
	defer c.mu.Unlock()

	return c.v[key]
}

func RunExample09() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
