package server

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Tracking-SYS/tracking-go/repo"
	"github.com/Tracking-SYS/tracking-go/services"

	kafkaLib "github.com/Tracking-SYS/go-lib/kafka"
	"github.com/Tracking-SYS/go-lib/kafka/ccloud"
)

const (
	//ProductKafkaTopic ...
	ProductKafkaTopic = "yuuxxq8y-product"

	//TaskKafkaTopic ...
	TaskKafkaTopic = "yuuxxq8y-task"
)

//KafkaConsumer ...
type KafkaConsumer struct {
	productService services.ProductServiceInterface
	taskService    services.TaskServiceInterface
	repo           repo.ProductRepoInterface
	consumer       *kafkaLib.Consumer
}

//NewKafkaConsumer ...
func NewKafkaConsumer(
	productService services.ProductServiceInterface,
	taskService services.TaskServiceInterface,
	repo repo.ProductRepoInterface,
) (*KafkaConsumer, error) {
	configPath := ccloud.ParseArgs()
	consumer := &kafkaLib.Consumer{
		ConfigFile: configPath,
	}
	return &KafkaConsumer{
		productService: productService,
		taskService:    taskService,
		repo:           repo,
		consumer:       consumer,
	}, nil
}

//Close ...
func (kc *KafkaConsumer) Close() error {
	return nil
}

//Start ...
func (kc *KafkaConsumer) Start() error {
	err := kc.consumer.InitConfig()
	if err != nil {
		fmt.Println("Init consumer config has error")
		os.Exit(1)
	}

	err = kc.consumer.CreateConsumerInstance()
	if err != nil {
		fmt.Println("create consumer has error")
		os.Exit(1)
	}

	kc.process()
	return nil
}

func (kc *KafkaConsumer) process() {
	consumerProductOutput := make(chan []byte, 1)
	consumerTaskOutput := make(chan []byte, 1)
	consumerChan := map[string]chan []byte{
		ProductKafkaTopic: consumerProductOutput,
		TaskKafkaTopic:    consumerTaskOutput,
	}
	topics := []string{ProductKafkaTopic, TaskKafkaTopic}

	go func() {
		kc.consumer.Start(consumerChan, topics)
	}()

	go func() {
		kc.handleCacheProduct(consumerProductOutput)
	}()

	go func() {
		kc.handleCacheTask(consumerTaskOutput)
	}()
}

func (kc *KafkaConsumer) handleCacheProduct(consumerProductOutput chan []byte) {
	for product := range consumerProductOutput {
		fmt.Printf("Consumer Set Product Cache: %v\n", string(product))
		productModel := &repo.ProductModel{}
		err := json.Unmarshal([]byte(product), productModel)
		if err != nil {
			fmt.Printf("Unmarshal consumed product has error: %v\n", err)
		}
		kc.productService.GetProduct(context.Background(), int(productModel.ID))
	}
}

func (kc *KafkaConsumer) handleCacheTask(consumerTaskOutput chan []byte) {
	for task := range consumerTaskOutput {
		fmt.Printf("Consumer Set Task Cache: %v\n", string(task))
		taskModel := &repo.TaskModel{}
		err := json.Unmarshal([]byte(task), taskModel)
		if err != nil {
			fmt.Printf("Unmarshal consumed task has error: %v\n", err)
		}
		kc.taskService.GetSingle(context.Background(), int(taskModel.ID))
	}
}
