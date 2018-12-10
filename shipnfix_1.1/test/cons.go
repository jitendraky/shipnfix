package main

import (
	"github.com/rafaeljesus/nsq-event-bus"
	"log"
	"sync"
)

type event struct{ Body string }

var wg sync.WaitGroup

func main() {
	wg.Add(1) // just to test purposes, the program will await for one message

	if err := bus.On(bus.ListenerConfig{
		Lookup:      []string{"localhost:4161"},
		Topic:       "events",
		Channel:     "consumer1",
		HandlerFunc: handler,
	}); err != nil {
		// handle failure to listen a message
		log.Println("Error while consuming message", err)
	}
  
	wg.Wait()
}

func handler(message *bus.Message) (reply interface{}, err error) {
	e := event{}
	if err = message.DecodePayload(&e); err != nil {
		// handle failure to decode a message
		log.Println("Error while consuming message", err)
		message.Finish()
		wg.Done()
		return
	}
	
	log.Println("[Consumer 1] Consuming message", e)
	message.Finish()
	wg.Done()
	return
}
