package services

import (
	"context"

	entities_pb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/entities"

	"github.com/Tracking-SYS/tracking-go/repo"
)

//TaskServiceInterface ...
type TaskServiceInterface interface {
	GetList(ctx context.Context, limit int, page int, ids []uint64) ([]*repo.TaskModel, error)
	GetSingle(ctx context.Context, id int) (*repo.TaskModel, error)
	Create(ctx context.Context, data *entities_pb.TaskInfo) (*repo.TaskModel, error)
	Transform(input []*repo.TaskModel) []*entities_pb.TaskInfo
	TransformSingle(prod *repo.TaskModel) *entities_pb.TaskInfo
}
