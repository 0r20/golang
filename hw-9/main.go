// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var i int

// func worker(mu *sync.Mutex) {
// 	mu.Lock()
// 	i++
// 	mu.Unlock()
// }

// func main() {
// 	wg := &sync.WaitGroup{}
// 	mu := &sync.Mutex{}

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			worker(mu)
// 			defer wg.Done()
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Printf("i = %d\n", i)
// }

package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func setup() {
	fmt.Print("Hello")
}

func doprint() {
	once.Do(setup)
}

func main() {
	go doprint()
	go doprint()
}
