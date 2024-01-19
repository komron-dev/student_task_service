package services

import (
	"context"

	pbt "student/student_task_service/genproto/task_service"
	l "student/student_task_service/pkg/logger"
	"student/student_task_service/services/grpcClient"
	"student/student_task_service/storage"

	// "github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

type taskService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewTaskService(db *sqlx.DB, log l.Logger, client grpcClient.IService) *taskService {
	return &taskService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *taskService) CreateTask(ctx context.Context, req *pbt.CreateTaskReq) (*pbt.CreateTaskRes, error) {
	res, err := s.storage.Task().CreateTask(req)
	if err != nil {
		s.logger.Error("error while creating a task", l.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *taskService) GetTask(ctx context.Context, req *pbt.ById) (*pbt.GetTaskRes, error) {
	res, err := s.storage.Task().GetTask(req)
	if err != nil {
		s.logger.Error("error while getting a task", l.Error(err))
		return nil, err
	}

	return res, nil

}

func (s *taskService) UpdateTask(ctx context.Context, req *pbt.UpdateTaskReq) (*pbt.Success, error) {
	res, err := s.storage.Task().UpdateTask(req)
	if err != nil {
		s.logger.Error("error while updating a task", l.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *taskService) DeleteTask(ctx context.Context, req *pbt.ById) (*pbt.Success, error) {
	res, err := s.storage.Task().DeleteTask(req)
	if err != nil {
		s.logger.Error("error while deleting a task", l.Error(err))
		return nil, err
	}
	return res, nil
}
func (s *taskService) ListOverDue(ctx context.Context, req *pbt.Empty) (*pbt.ListTasksRes, error) {
	res, err := s.storage.Task().ListOverDue(req)
	if err != nil {
		s.logger.Error("error while getting a list of tasks which were not completed in time", l.Error(err))
		return nil, err
	}
	return res, nil
}
func (s *taskService) ListTasks(ctx context.Context, req *pbt.ListTasksReq) (*pbt.ListTasksRes, error) {
	res, err := s.storage.Task().ListTasks(req)
	if err != nil {
		s.logger.Error("error while getting a list of all tasks")
		return nil, err
	}
	return res, nil
}
