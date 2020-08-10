package main

import (
	"fmt"

	"github.com/indikamaligaspe/go-concurrnecy/src/movies/channels"
	"github.com/indikamaligaspe/go-concurrnecy/src/movies/waitgroups"
)

func main() {
	fmt.Println("Starting With WAITGROUPS")
	waitgroups.StartWaitGroup()
	fmt.Println("Starting With CHANNELS")
	channels.StartChannels()
}
