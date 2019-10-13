package main

import (
	"log"
	"runtime"
	"sync"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

func doSomethingWith(c *gosocketio.Client, wg *sync.WaitGroup) {
	err := c.On("message", func(h *gosocketio.Channel, message string) {
		log.Println(message)
		wg.Done()
	})

	if err != nil {
		log.Printf("error: %v", err)
	}	
}

func connect() (*gosocketio.Client, error) {
	c, err := gosocketio.Dial(
		gosocketio.GetUrl("127.0.0.1", 3000, false),
		transport.GetDefaultWebsocketTransport())

	return c, err
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	c, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Fatal("Disconnected")
		c, err = connect()
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("Connected")
	})
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	for {
		wg.Add(1)
		go doSomethingWith(c, wg)
		wg.Wait()
	}
}
