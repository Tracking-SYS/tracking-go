package handler

import (
	"context"

	"github.com/Tracking-SYS/tracking-go/services"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	servicesPb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/services"
)

//TaskPBHandler ...
type TaskPBHandler struct {
	servicesPb.UnimplementedTaskServiceServer
	taskService services.TaskServiceInterface
	tracer      trace.Tracer
}

//NewTaskPBHandler ...
func NewTaskPBHandler(
	taskService services.TaskServiceInterface,
) *TaskPBHandler {
	tracer := otel.Tracer("NewTaskPBHandler")
	return &TaskPBHandler{
		tracer:      tracer,
		taskService: taskService,
	}
}

//Get List of Task
func (s *TaskPBHandler) Get(
	ctx context.Context,
	req *servicesPb.TaskServiceGetRequest,
) (*servicesPb.TaskServiceGetResponse, error) {
	ctx, span := s.tracer.Start(ctx, "Get")
	defer span.End()

	limit := req.GetLimit()
	page := req.GetPage()
	ids := req.GetIds()
	tasks, err := s.taskService.GetList(ctx, int(limit), int(page), ids)
	if err != nil {
		return nil, err
	}

	data := s.taskService.Transform(tasks)
	return &servicesPb.TaskServiceGetResponse{
		Data: data,
	}, nil
}

//GetSingle Task
func (s *TaskPBHandler) GetSingle(
	ctx context.Context,
	req *servicesPb.TaskServiceGetSingleRequest,
) (*servicesPb.TaskServiceGetSingleResponse, error) {
	ctx, span := s.tracer.Start(ctx, "GetSingle")
	defer span.End()

	ID := req.GetId()
	task, err := s.taskService.GetSingle(ctx, int(ID))
	if err != nil {
		return nil, err
	}

	data := s.taskService.TransformSingle(task)
	return &servicesPb.TaskServiceGetSingleResponse{
		Data: data,
	}, nil
}

//Create Task
func (s *TaskPBHandler) Create(
	ctx context.Context,
	req *servicesPb.TaskServiceCreateRequest,
) (*servicesPb.TaskServiceCreateResponse, error) {
	ctx, span := s.tracer.Start(ctx, "Create")
	defer span.End()

	data := req.GetData()
	task, err := s.taskService.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	data = s.taskService.TransformSingle(task)
	return &servicesPb.TaskServiceCreateResponse{
		Data: data,
	}, nil
}
