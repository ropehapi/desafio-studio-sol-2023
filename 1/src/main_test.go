package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPublishSubscribe(t *testing.T) {
	publisher := NewPublisher()

	// Criando subscribers para duas queues.
	subscriber1 := NewSubscriber("queue1")
	subscriber2 := NewSubscriber("queue2")

	publisher.Subscribe(subscriber1.Queue, subscriber1.Message)
	publisher.Subscribe(subscriber2.Queue, subscriber2.Message)

	// Publicando mensagens.
	publisher.Publish(Message{Queue: "queue1", Body: "Mensagem 1 na queue1"})
	publisher.Publish(Message{Queue: "queue2", Body: "Mensagem 1 na queue2"})

	// Aguardar um pouco para que as goroutines recebam as mensagens.
	time.Sleep(time.Second)

	// Verificar se as mensagens foram recebidas pelos subscribers.
	assert.Equal(t, "Mensagem 1 na queue1", <-subscriber1.Message)
	assert.Equal(t, "Mensagem 1 na queue2", <-subscriber2.Message)

	// Removendo a inscrição de subscriber2 da queue2.
	publisher.Unsubscribe(subscriber2.Queue, subscriber2.Message)

	// Publicando mais mensagens.
	publisher.Publish(Message{Queue: "queue1", Body: "Mensagem 2 na queue1"})
	publisher.Publish(Message{Queue: "queue2", Body: "Mensagem 2 na queue2"})

	// Aguardar um pouco para que as goroutines recebam as mensagens.
	time.Sleep(time.Second)

	// Verificar se subscriber2 não recebeu a segunda mensagem.
	select {
	case msg := <-subscriber2.Message:
		t.Errorf("Subscriber2 recebeu uma mensagem inesperada: %s", msg)
	default:
		// Subscriber2 não recebeu uma mensagem, isso é esperado.
	}

	// Verificar se subscriber1 recebeu a segunda mensagem.
	assert.Equal(t, "Mensagem 2 na queue1", <-subscriber1.Message)
}
