package repo

import (
	pbt "student/student_task_service/genproto/task_service"
)

//UserStorageI ...
type TaskStorageI interface {
	CreateTask(*pbt.CreateTaskReq) (*pbt.CreateTaskRes, error)
	GetTask(*pbt.ById) (*pbt.GetTaskRes, error)
	UpdateTask(*pbt.UpdateTaskReq) (*pbt.Success, error)
	DeleteTask(*pbt.ById) (*pbt.Success, error)
	ListOverDue(*pbt.Empty) (*pbt.ListTasksRes, error)
	ListTasks(*pbt.ListTasksReq) (*pbt.ListTasksRes, error)
}
