package kafka

import (
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

/*
	[KAFKA]
	1. Purpose:
		- Distributed, high-throughput, fault-tolerant event streaming platform.
		- Optimized for handling:
			1. High-volume, real-time data streams.
			2. Log aggregation and data pipelines.
			3. Decoupled communicationg between microservices
		===> Orchestrator for data streaming.

	2. Components:
		1. Producer: A component sends (publishes) data to Kafka topics
		2. Consumer: A component reads data from Kafka topics.
		3. Topic: Channel that groups messages. Messages is single unit of data.
		4. Partition: Each topic is split into partitions (instead of 1 single log/partition) -
		a unit of parallelism and scalability. Ensures:
			- Parallel writes/reads
			- Horizontal scalability
			- No bottleneck
			- Order is preserved within a partition

		Example:
						🧾 Message												📂 Partition						🗂 Topic							#️⃣ Offset
						{"user":"u1", "action":"login"}					0								user-activity								0
						{"user":"u2", "action":"logout"}				1								user-activity								0
						{"user":"u1", "action":"purchase"}			0								user-activity								1
						{"user":"u3", "action":"login"}					2								user-activity								0

	3. Example: Uber
		- Producer: Drive app sending GPS (driver-locations)
		- Topic: driver-locations
		- Parition: Driver-specific data (shared using drive_id)
			- driver_id = 123 => Parition 0
			- driver_id = 456 => Parition 1
			...
		- Consumer: Reads driver-locations topic and uses for: Map_UI, Matching-Engine, etc whatever consumer
		subscribes to this topic.
*/

var (
	kakfaProducer *kafka.Writer
)

const ()

// for producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// for consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		// server stores topic partitions handles read/write requests from clients, and coordinate
		// message distribution across the Kafka cluster
		// e.g: Broker: post office
		Brokers:        brokers, // []string{"localhost011", "localhost"}
		GroupID:        groupID, // consumer group this reader belongs to
		Topic:          topic,
		MinBytes:       10e3,              // 10KB
		MaxBytes:       10e6,              // 10MB
		CommitInterval: time.Second,       // how often consumer saves its current reading offset from the partition.
		StartOffset:    kafka.FirstOffset, // listen starting from 1st offset for new consumer
	})
}
