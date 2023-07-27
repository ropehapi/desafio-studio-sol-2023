package lib

import (
	"testing"
)

func TestNewPublisher(t *testing.T) {
	publisher := NewPublisher()

	// Verificar se o publisher foi criado corretamente.
	if publisher == nil {
		t.Error("Falha ao criar o publisher")
	}

	if len(publisher.subscribers) != 0 {
		t.Error("O mapa de subscribers não deveria conter nenhum elemento")
	}
}

func TestSubscribe(t *testing.T) {
	publisher := NewPublisher()

	subscriber1 := NewSubscriber("queue1")
	subscriber2 := NewSubscriber("queue1")

	// Verificar se o subscriber é adicionado corretamente ao publisher.
	publisher.Subscribe(subscriber1.Queue, subscriber1.Message)
	publisher.Subscribe(subscriber2.Queue, subscriber2.Message)

	if len(publisher.subscribers["queue1"]) != 2 {
		t.Error("Falha ao adicionar subscribers ao publisher")
	}
}

func TestUnsubscribe(t *testing.T) {
	publisher := NewPublisher()

	subscriber1 := NewSubscriber("queue1")
	subscriber2 := NewSubscriber("queue1")

	// Adicionar subscribers.
	publisher.Subscribe(subscriber1.Queue, subscriber1.Message)
	publisher.Subscribe(subscriber2.Queue, subscriber2.Message)

	// Verificar se o subscriber é removido corretamente do publisher.
	publisher.Unsubscribe(subscriber1.Queue, subscriber1.Message)
	if len(publisher.subscribers["queue1"]) != 1 {
		t.Error("Falha ao remover o subscriber1 do publisher")
	}

	publisher.Unsubscribe(subscriber2.Queue, subscriber2.Message)
	if len(publisher.subscribers["queue1"]) != 0 {
		t.Error("Falha ao remover o subscriber2 do publisher")
	}
}

func TestNewSubscriber(t *testing.T) {
	queueName := "test_queue"
	subscriber := NewSubscriber(queueName)

	// Verificar se o subscriber foi criado corretamente.
	if subscriber.Queue != queueName {
		t.Errorf("Esperado: %s, Recebido: %s", queueName, subscriber.Queue)
	}

	if subscriber.Message == nil {
		t.Error("O canal de mensagens não deveria ser nulo")
	}
}
