package mysql

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Tracking-SYS/tracking-go/infra"
	"github.com/Tracking-SYS/tracking-go/repo"

	kafkaLib "github.com/Tracking-SYS/go-lib/kafka"
	"github.com/Tracking-SYS/go-lib/kafka/ccloud"
	entities_pb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/entities"
)

var _ repo.TaskRepoInterface = &TaskMySQLRepo{}

//TaskMySQLRepo
type TaskMySQLRepo struct {
	db       *infra.ConnPool
	producer *kafkaLib.KafkaProducer
	topic    *string
}

//NewTaskMySQLRepo
func NewTaskMySQLRepo(
	db *infra.ConnPool,
) *TaskMySQLRepo {
	configPath = ccloud.ParseArgs()
	producerLib := &kafkaLib.KafkaProducer{
		ConfigFile: configPath,
	}
	producerLib.InitConfig()
	err := producerLib.CreateProducerInstance()
	if err != nil {
		fmt.Println("create producer has error")
		os.Exit(1)
	}
	producerLib.CreateTopic(TASK_KAFKA_TOPIC)
	topic := TASK_KAFKA_TOPIC
	return &TaskMySQLRepo{
		db:       db,
		producer: producerLib,
		topic:    &topic,
	}
}

//Get Task
func (t *TaskMySQLRepo) Get(ctx context.Context, limit int, page int, ids []uint64) (taskDAO []*repo.TaskModel, err error) {
	tx := t.db.Conn.WithContext(ctx)
	if limit != 0 {
		tx = tx.Limit(limit)
	}

	if page != 0 {
		tx = tx.Offset(page * limit)
	}

	tx = tx.Order("startAt")
	if ids != nil {
		tx = tx.Find(&taskDAO, ids)
	} else {
		tx = tx.Find(&taskDAO)
	}

	if err = tx.Error; err != nil {
		return nil, err
	}

	return taskDAO, nil
}

func (t *TaskMySQLRepo) Find(ctx context.Context, id int) (taskDAO *repo.TaskModel, err error) {
	if err = t.db.Conn.WithContext(ctx).First(&taskDAO, id).Error; err != nil {
		return nil, err
	}

	return taskDAO, nil
}

//Create Task
func (t *TaskMySQLRepo) Create(ctx context.Context, data *entities_pb.TaskInfo) (taskDAO *repo.TaskModel, err error) {
	taskDAO = &repo.TaskModel{}
	taskDAO.ID = uint64(data.Id)
	taskDAO.Name = data.Name
	taskDAO.StartAt = data.StartAt
	taskDAO.EndAt = data.EndAt
	taskDAO.Status = uint8(data.Status)

	result := t.db.Conn.WithContext(ctx).Create(&taskDAO)
	if result.Error != nil {
		return nil, result.Error
	}

	raw, err := json.Marshal(taskDAO)
	if err != nil {
		fmt.Println("parse data has error")
	}
	t.producer.ProduceMessage(t.topic, string(raw))

	return taskDAO, nil
}

//Update Task
func (t *TaskMySQLRepo) Update(ctx context.Context, task *repo.TaskModel) (err error) {
	err = t.db.Conn.WithContext(ctx).Model(task).Updates(task).Error
	return err
}
