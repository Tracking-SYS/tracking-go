package server

import (
	"context"
	"encoding/json"
	"fmt"

	"factory/exam/repo"
	"factory/exam/services"

	kafkaLib "github.com/Tracking-SYS/go-lib/kafka"
)

const (
	BOOTSTRAP_SERVERS = "bootstrap.servers"
	SASL_MECHANISMS   = "sasl.mechanisms"
	SECURITY_PROTOCOL = "security.protocol"
	SASL_USERNAME     = "sasl.username"
	SASL_PASSWORD     = "sasl.password"
)

const (
	PRODUCT_KAFKA_TOPIC = "testing_153_product"
	TASK_KAFKA_TOPIC    = "testing_153_task"
)

//KafkaConsumer ...
type KafkaConsumer struct {
	productService services.ProductServiceInterface
	taskService    services.TaskServiceInterface
	repo           repo.ProductRepoInterface
}

//NewKafkaConsumer ...
func NewKafkaConsumer(
	productService services.ProductServiceInterface,
	taskService services.TaskServiceInterface,
	repo repo.ProductRepoInterface,
) (*KafkaConsumer, error) {
	return &KafkaConsumer{
		productService: productService,
		taskService:    taskService,
		repo:           repo,
	}, nil
}

//Close ...
func (kc *KafkaConsumer) Close() error {
	return nil
}

//Start ...
func (kc *KafkaConsumer) Start() error {
	kc.consumeProduct()
	kc.consumeTask()

	return nil
}

func (kc *KafkaConsumer) consumeProduct() {
	consumerProductOutput := make(chan []byte, 1)
	go func() {
		kafkaLib.Start(consumerProductOutput, PRODUCT_KAFKA_TOPIC)
	}()
	go func() {
		for product := range consumerProductOutput {
			fmt.Printf("Consumer Set Cache: %v\n", string(product))
			productModel := &repo.ProductModel{}
			err := json.Unmarshal([]byte(product), productModel)
			if err != nil {
				fmt.Printf("Unmarshal consumed product has error: %v\n", err)
			}
			kc.productService.GetProduct(context.Background(), int(productModel.ID))
		}
	}()

}

func (kc *KafkaConsumer) consumeTask() {
	consumerTaskOutput := make(chan []byte, 1)
	go func() {
		kafkaLib.Start(consumerTaskOutput, TASK_KAFKA_TOPIC)
	}()
	go func() {
		for task := range consumerTaskOutput {
			fmt.Printf("Consumer Set Cache: %v\n", string(task))
			taskModel := &repo.TaskModel{}
			err := json.Unmarshal([]byte(task), taskModel)
			if err != nil {
				fmt.Printf("Unmarshal consumed task has error: %v\n", err)
			}
			kc.taskService.GetSingle(context.Background(), int(taskModel.ID))
		}
	}()
}
