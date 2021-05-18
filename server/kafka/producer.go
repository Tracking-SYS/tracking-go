// =============================================================================
//
// Produce messages to Confluent Cloud
// Using Confluent Golang Client for Apache Kafka
//
// =============================================================================

package kafka

/**
 * Copyright 2020 Confluent Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"context"
	"encoding/json"

	"fmt"
	"os"
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	"factory/exam/server/kafka/ccloud"
)

// ProduceRecordValue represents the struct of the value in a Kafka message
type ProduceRecordValue ccloud.RecordValue

type KafkaProducer struct {
}

//NewKafkaProducer ...
func NewKafkaProducer() (*KafkaProducer, error) {
	return &KafkaProducer{}, nil
}

//Close ...
func (kp *KafkaProducer) Close() error {
	return nil
}

//Start ...
func (kp *KafkaProducer) Start() error {

	// Initialization
	configFile, topic := ccloud.ParseArgs()
	conf := ccloud.ReadCCloudConfig(*configFile)

	// Create Producer instance
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": conf["bootstrap.servers"],
		"sasl.mechanisms":   conf["sasl.mechanisms"],
		"security.protocol": conf["security.protocol"],
		"sasl.username":     conf["sasl.username"],
		"sasl.password":     conf["sasl.password"]})
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Create topic if needed
	CreateTopic(p, *topic)

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	deliveryChan := make(chan kafka.Event)

	recordKey := "tracking"
	data := &ProduceRecordValue{
		Count: 0}
	recordValue, _ := json.Marshal(&data)
	fmt.Printf("Preparing to produce record: %s\t%s\n", recordKey, recordValue)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: topic, Partition: int32(kafka.PartitionAny)},
		Key:            []byte(recordKey),
		Value:          []byte(recordValue),
	}, deliveryChan)
	if err != nil {
		fmt.Printf("Produce message has error: key %v, val %v", recordKey, recordValue)
		return err
	}

	// Wait for all messages to be delivered
	kafkaEvent := <-deliveryChan
	msg := kafkaEvent.(*kafka.Message)
	if msg.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", msg.TopicPartition.Error)
		return msg.TopicPartition.Error
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset)
	}

	close(deliveryChan)

	return nil
}

// CreateTopic creates a topic using the Admin Client API
func CreateTopic(p *kafka.Producer, topic string) {

	a, err := kafka.NewAdminClientFromProducer(p)
	if err != nil {
		fmt.Printf("Failed to create new admin client from producer: %s", err)
		os.Exit(1)
	}
	// Contexts are used to abort or limit the amount of time
	// the Admin call blocks waiting for a result.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Create topics on cluster.
	// Set Admin options to wait up to 60s for the operation to finish on the remote cluster
	maxDur, err := time.ParseDuration("60s")
	if err != nil {
		fmt.Printf("ParseDuration(60s): %s", err)
		os.Exit(1)
	}
	results, err := a.CreateTopics(
		ctx,
		// Multiple topics can be created simultaneously
		// by providing more TopicSpecification structs here.
		[]kafka.TopicSpecification{{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 3}},
		// Admin options
		kafka.SetAdminOperationTimeout(maxDur))
	if err != nil {
		fmt.Printf("Admin Client request error: %v\n", err)
		os.Exit(1)
	}
	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError && result.Error.Code() != kafka.ErrTopicAlreadyExists {
			fmt.Printf("Failed to create topic: %v\n", result.Error)
			os.Exit(1)
		}
		fmt.Printf("%v\n", result)
	}
	a.Close()

}
