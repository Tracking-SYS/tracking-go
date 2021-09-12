package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Tracking-SYS/tracking-go/repo"
	"github.com/Tracking-SYS/tracking-go/repo/cache"

	entities_pb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/entities"
)

var _ TaskServiceInterface = &TaskService{}

//TaskProvider ...
func TaskProvider(
	taskRepo repo.TaskRepoInterface,
	cacheRepo cache.CacheInteface,
) *TaskService {
	return &TaskService{
		taskRepo:  taskRepo,
		cacheRepo: cacheRepo,
	}
}

//TaskService ...
type TaskService struct {
	taskRepo  repo.TaskRepoInterface
	cacheRepo cache.CacheInteface
}

//GetList ...
func (ts *TaskService) GetList(ctx context.Context, limit int, page int, ids []uint64) ([]*repo.TaskModel, error) {
	tasks, err := ts.taskRepo.Get(ctx, limit, page, ids)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

//GetSingle ...
func (ts *TaskService) GetSingle(ctx context.Context, id int) (*repo.TaskModel, error) {
	task, err := ts.cacheRepo.Get(ctx, strconv.Itoa(id))

	if err != nil {
		return nil, err
	}

	if task != nil {
		fmt.Printf("GetCache: %v\n", task)
		return ts.parseData(task.(map[string]interface{}))
	}

	task, err = ts.taskRepo.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	err = ts.cacheRepo.Set(ctx, fmt.Sprintf("task_%s", strconv.Itoa(id)), task)
	if err != nil {
		return nil, err
	}
	return task.(*repo.TaskModel), nil
}

func (ts *TaskService) parseData(data map[string]interface{}) (task *repo.TaskModel, err error) {
	jsonbody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonbody, &task); err != nil {
		return nil, err
	}

	return task, nil
}

//Create ...
func (ts *TaskService) Create(ctx context.Context, data *entities_pb.TaskInfo) (*repo.TaskModel, error) {
	task, err := ts.taskRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return task, nil
}

//Transform ...
func (ts *TaskService) Transform(input []*repo.TaskModel) []*entities_pb.TaskInfo {
	result := []*entities_pb.TaskInfo{}
	for _, task := range input {
		item := &entities_pb.TaskInfo{
			Id:      uint32(task.ID),
			Name:    task.Name,
			StartAt: task.StartAt,
			EndAt:   task.EndAt,
			Status:  uint32(task.Status),
		}
		result = append(result, item)
	}

	return result
}

//TransformSingle ...
func (ts *TaskService) TransformSingle(task *repo.TaskModel) *entities_pb.TaskInfo {
	if task == nil {
		return nil
	}

	result := &entities_pb.TaskInfo{
		Id:      uint32(task.ID),
		Name:    task.Name,
		StartAt: task.StartAt,
		EndAt:   task.EndAt,
		Status:  uint32(task.Status),
	}

	return result
}
