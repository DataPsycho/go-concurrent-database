package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	// Solution of inmemore cash
	// as the both go routines is accessing our cache map
	// at the same time, we can use mutex for that
	m := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryCache(id, m); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		// Pointer to a waitgroup
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryDatabase(id, m); ok {
				fmt.Println("from the database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		time.Sleep(2000 * time.Millisecond)
	}
	// Wait untill all wait group are done
	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	// The Rlock will allow multiple reader to read from cache
	// at the same time, so some routine if want to write in to cache
	// Mutex will up all the reader and then give access to writers
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(3 * time.Second)
	for _, b := range books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
