package concurrency

import (
	"fmt"
	"sync"
)

/*
Reference : https://youtu.be/5Z8skvm4g64?si=XAYq7x51f2C_DRr6
*/

const KEY string = "sample_key"

type SafeCounter struct {
	mu     sync.Mutex
	NumMap map[string]int // maps are not thread-safe in go
}

func (s *SafeCounter) Add(num int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.NumMap[KEY] = num
}
func UpdateMapUsingMutex() {
	sc := SafeCounter{
		NumMap: make(map[string]int),
	}
	var wg sync.WaitGroup

	for i := 1; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sc.Add(i)
		}(i)
	}
	wg.Wait()

	fmt.Printf("UpdateMapUsingMutex() %v", sc.NumMap[KEY])

}
