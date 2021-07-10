package repo

import (
	"context"

	entities_pb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/entities"
)

//TaskRepoInterface ...
type TaskRepoInterface interface {
	Get(context context.Context, limit int, page int, ids []uint64) (taskDAO []*TaskModel, err error)
	Find(context context.Context, id int) (taskDAO *TaskModel, err error)
	Create(context context.Context, id *entities_pb.TaskInfo) (taskDAO *TaskModel, err error)
}
