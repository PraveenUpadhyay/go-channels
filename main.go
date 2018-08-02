package main

import (
	"fmt"
	"sync"

	"github.com/PraveenUpadhyay/go-channels/communication"
)

func main() {
	max := 10
	var wg sync.WaitGroup
	wg.Add(max)

	for i := 1; i <= max; i++ {
		var forceSend bool
		if i%2 == 0 {
			forceSend = true
		}
		go func(i int) {
			for n := 1; n <= 1000; n++ {
				communication.New().SendMessage(fmt.Sprintf("Go-routine-%d-%d", i, n), forceSend)
				//time.Sleep(1 * time.Second)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
