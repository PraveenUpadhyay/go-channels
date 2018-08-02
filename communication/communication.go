package communication

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var once sync.Once
var obj communication

type communication struct {
	clearQueue chan bool
	offline    chan bool
}

func New() communication {
	once.Do(func() {
		obj = communication{
			clearQueue: make(chan bool),
			offline:    make(chan bool),
		}
		go obj.emptyQueue()
	})
	return obj
}
func (c communication) SendMessage(msg string, forceSend bool) {
	select {
	case <-c.offline:
		if forceSend {
			fmt.Printf("Offline & ForceSend - %s,%v\n", msg, forceSend)
			if !success() {
				go func() { c.offline <- true }()

			}
			go func() {
				fmt.Println("////////////////")
				c.clearQueue <- true
			}()
			//if success then don't write/enable offline channel
			return
		}
		fmt.Printf("Offline & Persist - %s,%v\n", msg, forceSend)
		go func() { c.offline <- true }()
	default:
		//online
		fmt.Printf("online - %s,%v\n", msg, forceSend)
		if !success() {
			go func() { c.offline <- true }()
		}
	}
}

func (c communication) emptyQueue() {
	for {
		select {
		case <-c.clearQueue:
			fmt.Println("==========================================================")
			for i := 1; i < 10; i++ {
				fmt.Printf("Clearing Queue %d\n", i)
				if !success() {
					go func() { c.offline <- true }()
					return
				}
				time.Sleep(1 * time.Second)
			}
		}
	}
}

func success() bool {
	rand.NewSource(10)
	n := rand.Intn(10)
	return n%2 == 0
}
