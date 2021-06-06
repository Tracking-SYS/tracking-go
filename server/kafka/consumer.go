// =============================================================================
//
// Consume messages from Confluent Cloud
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
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	"factory/exam/server/kafka/ccloud"
)

const (
	BOOTSTRAP_SERVERS = "bootstrap.servers"
	SASL_MECHANISMS   = "sasl.mechanisms"
	SECURITY_PROTOCOL = "security.protocol"
	SASL_USERNAME     = "sasl.username"
	SASL_PASSWORD     = "sasl.password"
)

// ConsumeRecordValue represents the struct of the value in a Kafka message
type ConsumeRecordValue ccloud.RecordValue

//KafkaConsumer ...
type KafkaConsumer struct {
}

//NewKafkaConsumer ...
func NewKafkaConsumer() (*KafkaConsumer, error) {
	return &KafkaConsumer{}, nil
}

//Close ...
func (kc *KafkaConsumer) Close() error {
	return nil
}

//Start ...
func (kc *KafkaConsumer) Start() error {

	// Initialization
	configFile, topic := ccloud.ParseArgs()
	conf := ccloud.ReadCCloudConfig(*configFile)

	fmt.Print(conf[BOOTSTRAP_SERVERS])
	// Create Consumer instance
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		BOOTSTRAP_SERVERS:   conf[BOOTSTRAP_SERVERS],
		SASL_MECHANISMS:     conf[SASL_MECHANISMS],
		SECURITY_PROTOCOL:   conf[SECURITY_PROTOCOL],
		SASL_USERNAME:       conf[SASL_USERNAME],
		SASL_PASSWORD:       conf[SASL_PASSWORD],
		"group.id":          "go_example_group_1",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	// Subscribe to topic
	err = c.SubscribeTopics([]string{*topic}, nil)
	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	totalCount := 0
	run := true
	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			msg, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			recordKey := string(msg.Key)
			recordValue := msg.Value
			data := ConsumeRecordValue{}
			err = json.Unmarshal(recordValue, &data)
			if err != nil {
				fmt.Printf("Failed to decode JSON at offset %d: %v", msg.TopicPartition.Offset, err)
				continue
			}
			count := data.Count
			totalCount += count
			fmt.Printf("Consumed record with key %s and value %s, and updated total count to %d\n", recordKey, recordValue, totalCount)
		}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()

	return nil
}
