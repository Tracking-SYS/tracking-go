package mysql

import (
	kafkaLib "github.com/Tracking-SYS/go-lib/kafka"
)

var (
	configPath  *string
	producerLib *kafkaLib.Producer
)

const (
	//ProductKafkaTopic ...
	ProductKafkaTopic = "yuuxxq8y-product"

	//TaskKafkaTopic ...
	TaskKafkaTopic = "yuuxxq8y-task"
)
