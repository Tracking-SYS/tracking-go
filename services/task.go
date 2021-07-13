package services

import (
	"context"
	"encoding/json"
	"factory/exam/repo"
	"factory/exam/repo/cache"
	"fmt"
	"strconv"

	entities_pb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/entities"
)

var _ TaskServiceInterface = &TaskService{}

//TaskProvider
func TaskProvider(
	taskRepo repo.TaskRepoInterface,
	cacheRepo cache.CacheInteface,
) *TaskService {
	return &TaskService{
		taskRepo:  taskRepo,
		cacheRepo: cacheRepo,
	}
}

//TaskService
type TaskService struct {
	taskRepo  repo.TaskRepoInterface
	cacheRepo cache.CacheInteface
}

//GetList Task
func (ts *TaskService) GetList(ctx context.Context, limit int, page int, ids []uint64) []*repo.TaskModel {
	tasks, err := ts.taskRepo.Get(ctx, limit, page, ids)
	if err != nil {
		return nil
	}

	return tasks
}

//GetSingle Task
func (ts *TaskService) GetSingle(ctx context.Context, id int) *repo.TaskModel {
	task, err := ts.cacheRepo.Get(ctx, strconv.Itoa(id))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	if task != nil {
		fmt.Printf("GetCache: %v\n", task)
		return ts.parseData(task.(map[string]interface{}))
	}

	task, err = ts.taskRepo.Find(ctx, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = ts.cacheRepo.Set(ctx, fmt.Sprintf("task_%s", strconv.Itoa(id)), task)
	if err != nil {
		fmt.Println(err)
	}
	return task.(*repo.TaskModel)
}

func (ts *TaskService) parseData(data map[string]interface{}) (task *repo.TaskModel) {
	jsonbody, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := json.Unmarshal(jsonbody, &task); err != nil {
		fmt.Println(err)
		return
	}

	return task
}

//Create Task
func (ts *TaskService) Create(ctx context.Context, data *entities_pb.TaskInfo) *repo.TaskModel {
	task, err := ts.taskRepo.Create(ctx, data)
	if err != nil {
		return nil
	}

	return task
}

//Transform Task
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

//TransformSingle Task
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
