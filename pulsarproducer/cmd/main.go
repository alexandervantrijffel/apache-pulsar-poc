package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
	log.Printf("hi")
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	defer client.Close()
	log.Printf("client initialized")
	go produceMessage(client)
	consumeMessages(client)
}
func produceMessage(client pulsar.Client) {

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "myevents",
		MessageRouter: func(msg *pulsar.ProducerMessage, tm pulsar.TopicMetadata) int {
			fmt.Println("Routing message ", msg, " -- Partitions: ", tm.NumPartitions())
			return 2
		},
	})
	if err != nil {
		log.Fatalf("failed to create producer: %v", err)
		return
	}

	defer producer.Flush()
	defer producer.Close()

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte("hello"),
	})
	if err != nil {
		log.Fatalf("Failed to produce message: %v", err)
		return
	}

	if err != nil {
		fmt.Println("Failed to publish message", err)
	}
	fmt.Printf("Published message to %s with sequenceid %d\n", producer.Topic(),
		producer.LastSequenceID)
}
func consumeMessages(client pulsar.Client) {
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "myevents",
		SubscriptionName: "test",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()
	ctx, canc := context.WithTimeout(context.Background(), 5*time.Second)
	msg, err := consumer.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("RECV %s\n", msg.Payload())
	canc()
}
