package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	publisher := NewPublisher()

	// Criando subscribers para duas queues.
	subscriber1 := NewSubscriber("queue1")
	subscriber2 := NewSubscriber("queue2")

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
	publisher.Publish(Message{Queue: "queue1", Body: "Mensagem 1 na queue1"})
	publisher.Publish(Message{Queue: "queue2", Body: "Mensagem 1 na queue2"})

	// Aguardar um pouco para que as goroutines recebam as mensagens.
	time.Sleep(time.Second * 2)

	// Removendo a inscrição de subscriber2 da queue2.
	publisher.Unsubscribe(subscriber2.Queue, subscriber2.Message)

	// Publicando mais mensagens.
	publisher.Publish(Message{Queue: "queue1", Body: "Mensagem 2 na queue1"})
	publisher.Publish(Message{Queue: "queue2", Body: "Mensagem 2 na queue2"})

	// Aguardar um pouco para que as goroutines recebam as mensagens.
	time.Sleep(time.Second * 2)
}

// Message representa uma mensagem a ser publicada.
type Message struct {
	Queue string
	Body  string
}

// Publisher é responsável por publicar mensagens nas queues.
type Publisher struct {
	subscribers map[string][]chan<- string
	lock        sync.RWMutex
}

// NewPublisher cria um novo Publisher.
func NewPublisher() *Publisher {
	return &Publisher{
		subscribers: make(map[string][]chan<- string),
	}
}

// Subscribe inscreve um canal (channel) para receber mensagens de uma queue específica.
func (p *Publisher) Subscribe(queue string, ch chan<- string) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.subscribers[queue] = append(p.subscribers[queue], ch)
}

// Unsubscribe cancela a inscrição de um canal (channel) em uma queue específica.
func (p *Publisher) Unsubscribe(queue string, ch chan<- string) {
	p.lock.Lock()
	defer p.lock.Unlock()

	subscribers := p.subscribers[queue]
	for i, subscriber := range subscribers {
		if subscriber == ch {
			p.subscribers[queue] = append(subscribers[:i], subscribers[i+1:]...)
			break
		}
	}
}

// Publish publica uma mensagem em uma queue específica e registra no arquivo de log.
func (p *Publisher) Publish(message Message) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	for _, ch := range p.subscribers[message.Queue] {
		ch <- message.Body
	}

	p.writeToLogFile(message)
}

// writeToLogFile registra a mensagem no arquivo de log.
func (p *Publisher) writeToLogFile(message Message) {
	fileName := fmt.Sprintf("%s.log", message.Queue)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo %s: %s\n", fileName, err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(message.Body + "\n"); err != nil {
		fmt.Printf("Erro ao escrever no arquivo %s: %s\n", fileName, err)
	}
}

// Subscriber representa um subscriber que recebe mensagens de uma queue específica.
type Subscriber struct {
	Queue   string
	Message chan string
}

// NewSubscriber cria um novo Subscriber.
func NewSubscriber(queue string) *Subscriber {
	return &Subscriber{
		Queue:   queue,
		Message: make(chan string),
	}
}
