package main

import (
	"fmt"
	"time"

	"github.com/ropehapi/pubsub/lib"
)

//Faz uso da biblioteca criada
func main() {
	publisher := lib.NewPublisher()

	// Criando subscribers para duas queues.
	subscriber1 := lib.NewSubscriber("queue1")
	subscriber2 := lib.NewSubscriber("queue2")

	publisher.Subscribe(subscriber1.Queue, subscriber1.Message)
	publisher.Subscribe(subscriber2.Queue, subscriber2.Message)

	// Lendo as mensagens dos subscribers.
	go func() {
		for msg := range subscriber1.Message {
			fmt.Printf("Subscriber1 - Received: %s\n", msg)
		}
	}()

	go func() {
		for msg := range subscriber2.Message {
			fmt.Printf("Subscriber2 - Received: %s\n", msg)
		}
	}()

	// Publicando mensagens.
	publisher.Publish(lib.Message{Queue: "queue1", Body: "Mensagem 1 na queue1"})
	publisher.Publish(lib.Message{Queue: "queue2", Body: "Mensagem 1 na queue2"})

	// Aguardar um pouco para que as goroutines recebam as mensagens.
	time.Sleep(time.Second * 2)

	// Removendo a inscrição de subscriber2 da queue2.
	publisher.Unsubscribe(subscriber2.Queue, subscriber2.Message)

	// Publicando mais mensagens.
	publisher.Publish(lib.Message{Queue: "queue1", Body: "Mensagem 2 na queue1"})
	publisher.Publish(lib.Message{Queue: "queue2", Body: "Mensagem 2 na queue2"})

	// Aguardar um pouco para que as goroutines recebam as mensagens.
	time.Sleep(time.Second * 2)
}