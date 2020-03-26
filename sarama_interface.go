package kafka

import "github.com/Shopify/sarama"

//go:generate moq -out ./mock/sarama.go -pkg mock . Sarama
//go:generate moq -out ./mock/sarama_client.go -pkg mock . SaramaClient
//go:generate moq -out ./mock/sarama_async_producer.go -pkg mock . SaramaAsyncProducer
//go:generate moq -out ./mock/sarama_consumer_group.go -pkg mock . SaramaConsumerGroup

// Sarama is an interface representing the Sarama library.
type Sarama interface {
	NewClient(addrs []string, conf *sarama.Config) (SaramaClient, error)
	NewAsyncProducerFromClient(client SaramaClient) (SaramaAsyncProducer, error)
	NewConsumerGroupFromClient(groupID string, client SaramaClient) (SaramaConsumerGroup, error)
}

// SaramaClient is a wrapper around sarama.Client
type SaramaClient = sarama.Client

// SaramaAsyncProducer is a wrapper around sarama.AsyncProducer
type SaramaAsyncProducer = sarama.AsyncProducer

// SaramaConsumerGroup is a wrapper around sarama.ConsumerGroup
type SaramaConsumerGroup = sarama.ConsumerGroup

// SaramaLib implements Sarama interface and wraps the real calls to Sarama library.
type SaramaLib struct{}

// NewClient creates a new sarama Client using the provided broker addresses and configuration
func (s *SaramaLib) NewClient(addrs []string, conf *sarama.Config) (SaramaClient, error) {
	return sarama.NewClient(addrs, conf)
}

// NewAsyncProducerFromClient creates a new SaramaAsyncProducer using the given client
func (s *SaramaLib) NewAsyncProducerFromClient(client SaramaClient) (SaramaAsyncProducer, error) {
	return sarama.NewAsyncProducerFromClient(client)
}

// NewConsumerGroupFromClient creates a new SaramaConsumerGroup using the given client
func (s *SaramaLib) NewConsumerGroupFromClient(groupID string, client SaramaClient) (SaramaConsumerGroup, error) {
	return sarama.NewConsumerGroupFromClient(groupID, client)
}
