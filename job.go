package main

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type job struct {
	topic    string
	producer *kafka.Producer
	msg      KafkaMsg
}

func (self *job) Name() string {
	return self.topic
}

func (self *job) Execute() error {
	b, _ := json.Marshal(self.msg)
	self.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &(self.topic), Partition: kafka.PartitionAny},
		Value:          b,
	}, nil)
	return nil
}
