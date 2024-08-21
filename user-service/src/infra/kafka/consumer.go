package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"go_kafka_playground/user-service/src/infra"
)

var logger = infra.Logger()

func StartConsumer() {
	brokers := []string{"localhost:9092"}
	topic := "user-topic"
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		logger.Log().Msg(fmt.Sprintf("Error creating consumer: %s", err.Error()))
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		logger.Log().Msg(fmt.Sprintf("Error starting consumer for partition: %s", err.Error()))
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		logger.Log().Msg(fmt.Sprintf("Consumed message: %s", string(message.Value)))
	}
}
