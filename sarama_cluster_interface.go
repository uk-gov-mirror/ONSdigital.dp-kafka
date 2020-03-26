package kafka

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

//go:generate moq -out ./mock/sarama_cluster.go -pkg mock . SaramaCluster
//go:generate moq -out ./mock/sarama_cluster_consumer.go -pkg mock . SaramaClusterConsumer

// SaramaCluster is an interface representing the bsm sarama-cluster library.
type SaramaCluster interface {
	NewConsumerFromClient(client *cluster.Client, groupID string, topics []string) (SaramaClusterConsumer, error)
	NewClient(addrs []string, config *cluster.Config) (*cluster.Client, error)
}

// SaramaClusterConsumer is an interface representing the bsm sarama-cluster Consumer struct
type SaramaClusterConsumer interface {
	Close() (err error)
	Messages() <-chan *sarama.ConsumerMessage
	CommitOffsets() error
	Errors() <-chan error
	Notifications() <-chan *cluster.Notification
	MarkOffset(msg *sarama.ConsumerMessage, metadata string)
}

// SaramaClusterLib implements SaramaCluster interface and wraps the real calls to bsm sarama-cluster library.
type SaramaClusterLib struct{}

// NewConsumerFromClient creates a new sarama cluster consumer.
func (c *SaramaClusterLib) NewConsumerFromClient(client *cluster.Client, groupID string, topics []string) (SaramaClusterConsumer, error) {
	return cluster.NewConsumerFromClient(client, groupID, topics)
}

// NewClient creates a new sarama cluster client.
func (c *SaramaClusterLib) NewClient(addrs []string, config *cluster.Config) (*cluster.Client, error) {
	return cluster.NewClient(addrs, config)
}
