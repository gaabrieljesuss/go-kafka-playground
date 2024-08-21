package main

import "go_kafka_playground/user-service/src/infra/kafka"

func main() {
	// Start Kafka consumer
	kafka.StartConsumer()
}
