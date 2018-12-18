package main

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)


func Producer() {
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		fmt.Println("NewProducer", err)
		panic(err)
	}

	i := 1
	for {
		if err := producer.Publish("test", []byte(fmt.Sprintf("Hello World %d", i))); err != nil {
				fmt.Println("Publish", err)
			panic(err)
		}

		time.Sleep(time.Second * 1)

		i++
		fmt.Println(i)
	}
}

func main() {
	Producer()
}
