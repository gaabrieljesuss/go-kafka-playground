package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"go_kafka_playground/api-gateway/src/core/domain/errors"
	"go_kafka_playground/api-gateway/src/infra"
)

var logger = infra.Logger()

func ProduceMessage(brokers []string, topic string, message string) errors.Error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		logger.Log().Msg(fmt.Sprintf("Failed to start Sarama producer: %s", err.Error()))
		return errors.NewUnexpected()
	}

	defer func() {
		if err := producer.Close(); err != nil {
			logger.Log().Msg(fmt.Sprintf("Error closing producer: %v", err))
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		logger.Log().Msg(fmt.Sprintf("Failed to send message: %s", err.Error()))
		return errors.NewUnexpected()
	}

	logger.Log().Msg(fmt.Sprintf("Message sent to partition %d at offset %d", partition, offset))
	return nil
}
