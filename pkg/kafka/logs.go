package kafka

import (
	"context"
	"fmt"

	"progo/core/config"
	corekafka "progo/core/kafka"
)

var logsTopic = "logs"

// NewLogsConsumer creates a new consumer for logs messages.
func NewLogsConsumer() corekafka.Consumer {

	brokers := getBrokers()

	return corekafka.NewKafkaConsumer(logsTopic, brokers)
}

// Read gets Kafka messages and store them to database.
func Read(r corekafka.Consumer) {
	fmt.Println("Reading")
	for {
		fmt.Println("Here")
		m, err := r.ReadMessage(context.Background())
		fmt.Println("Here2")
		if err != nil {
			fmt.Println(err)

			break
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	r.Close()
}

func getBrokers() []string {
	brokers := config.Get("KAFKA_BROKERS")

	return []string{brokers}
}
