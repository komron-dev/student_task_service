package grpcClient

import (
	"student/student_task_service/config"
	pbt "student/student_task_service/genproto/task_service"
)

type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

type IService interface {
	TaskService() pbt.TaskServiceClient
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {
	return &GrpcClient{
		cfg:         cfg,
		connections: map[string]interface{}{},
	}, nil
}

func (g *GrpcClient) TaskService() pbt.TaskServiceClient {
	return g.connections["student_task_service"].(pbt.TaskServiceClient)
}
