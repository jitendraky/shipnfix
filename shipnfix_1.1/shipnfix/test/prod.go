package main

import (
	"github.com/rafaeljesus/nsq-event-bus"
	"log"
)

type event struct{ Body string }

func main() {
	topic := "events"
  
	emitter, err := bus.NewEmitter(bus.EmitterConfig{})
	
	if err != nil {
		log.Fatal("[ERRO]", err)
	}
	
	message := "[Emitter 1] sending message"
	e := event{message}
	
	if err = emitter.Emit(topic, &e); err != nil {
		log.Println("error while was emitting message", err)
	}
	
	log.Println("[Message emitted]", message)
}
