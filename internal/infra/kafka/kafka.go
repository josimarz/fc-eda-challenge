package kafka

import (
	"encoding/json"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (c *Consumer) Consume(ch chan<- *ckafka.Message) error {
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		return err
	}
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		return err
	}
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			ch <- msg
		}
	}
}

type Producer struct {
	ConfigMap *ckafka.ConfigMap
}

func NewProducer(configMap *ckafka.ConfigMap) *Producer {
	return &Producer{ConfigMap: configMap}
}

func (p *Producer) Publish(msg interface{}, key []byte, topic string) error {
	producer, err := ckafka.NewProducer(p.ConfigMap)
	if err != nil {
		return err
	}
	value, err := json.Marshal(&msg)
	if err != nil {
		return err
	}
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          value,
		Key:            key,
	}
	if err := producer.Produce(message, nil); err != nil {
		return err
	}
	return nil
}
