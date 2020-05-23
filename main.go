package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pheever/stresster/core/feed"
)

type foo struct {
	Result,
	Sample feed.Feeder
	Channel chan int
}

func main() {
	f := foo{}
	for i := 0; i < 20000; i++ {
		go receiver(i, &f.Sample)
	}
	time.Sleep(1 * time.Second)
	go func() {
		var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < 3; i++ {
			f.Sample.Feed() <- seededRand.Intn(999999)
		}
	}()
	time.Sleep(time.Duration(1 * time.Second))
	fmt.Println(<-f.Sample.Close())
	panic("Show me the benjies")
}

func receiver(id int, feed feed.Feed) {
	fmt.Printf("Hi i'm %d\n", id)
	for i := range feed.Subscribe() {
		if number, ok := i.(int); ok {
			fmt.Printf("%d read %d\n", id, number)
		} else {
			fmt.Printf("%d NOT AN INT\n", id)
		}
	}
	fmt.Printf("%d closed\n", id)
}
