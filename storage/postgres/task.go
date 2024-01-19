package postgres

import (
	"time"

	pbt "student/student_task_service/genproto/task_service"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type taskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) *taskRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) CreateTask(req *pbt.CreateTaskReq) (*pbt.CreateTaskRes, error) {
	res := pbt.CreateTaskRes{}
	query := `INSERT INTO student_tasks(
		id, 
		title,
		deadline,
		summary,
		type,
		student_id,
		status,
		created_at)
        VALUES($1,$2,$3,$4,$5,$6,$7,$8)
        RETURNING 
			id,
			title,
			deadline,
			summary,
			student_id,
			status,
			created_at,
			type`
	uuid := uuid.New().String()
	time := time.Now().Format(time.RFC3339)
	err := r.db.QueryRow(query,
		uuid,
		req.Title,
		req.Deadline,
		req.Summary,
		req.Type,
		req.StudentId,
		"incomplete",
		time).
		Scan(
			&res.Id,
			&res.Title,
			&res.Deadline,
			&res.Summary,
			&res.StudentId,
			&res.Status,
			&res.CreatedAt,
			&res.Type,
		)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *taskRepo) GetTask(req *pbt.ById) (*pbt.GetTaskRes, error) {
	res := pbt.GetTaskRes{}
	query := `SELECT 
			id,
			title,
			deadline,
			summary,
			student_id,
			status,
			created_at,
			type
		FROM student_tasks WHERE id=$1 AND deleted_at IS NULL`
	err := r.db.QueryRow(query, req.TaskId).Scan(
		&res.Id,
		&res.Title,
		&res.Deadline,
		&res.Summary,
		&res.StudentId,
		&res.Status,
		&res.CreatedAt,
		&res.Type,
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *taskRepo) UpdateTask(req *pbt.UpdateTaskReq) (*pbt.Success, error) {

	time := time.Now().Format(time.RFC3339)

	query := `UPDATE student_tasks SET 
	title = $1,
	deadline = $2,
	summary = $3,
	student_id = $4,
	status = $5,
	type = $6,
	updated_at = $7
	WHERE id=$8 AND
	deleted_at IS NULL`
	_, err := r.db.Exec(query,
		req.Title,
		req.Deadline,
		req.Summary,
		req.StudentId,
		req.Status,
		req.Type,
		time,
		req.Id)
	if err != nil {
		return nil, err
	}
	return &pbt.Success{Message: "ok"}, nil
}

func (r *taskRepo) DeleteTask(req *pbt.ById) (*pbt.Success, error) {
	query := `UPDATE student_tasks SET deleted_at=$1 WHERE id=$2 and deleted_at IS NULL`
	time := time.Now().Format(time.RFC3339)
	_, err := r.db.Exec(query, time, req.TaskId)
	if err != nil {
		return nil, err
	}
	return &pbt.Success{Message: "ok"}, nil
}

func (r *taskRepo) ListOverDue(req *pbt.Empty) (*pbt.ListTasksRes, error) {
	var count int64

	query := `SELECT 
		id,
		title,
		student_id,
		deadline,
		summary,
		status,
		created_at,
		type
	FROM student_tasks WHERE deadline<$1  AND deleted_at IS NULL`
	current_time := time.Now().Format(time.RFC3339)
	rows, err := r.db.Query(query, current_time)
	if err != nil {
		return nil, err
	}
	res := pbt.ListTasksRes{}
	for rows.Next() {
		task := pbt.GetTaskRes{}
		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.StudentId,
			&task.Deadline,
			&task.Summary,
			&task.Status,
			&task.CreatedAt,
			&task.Type,
		)
		if err != nil {
			return nil, err
		}

		res.Tasks = append(res.Tasks, &task)
	}
	queryForCount := `SELECT COUNT(*) FROM student_tasks WHERE deleted_at IS NULL AND deadline<$1`
	row := r.db.QueryRow(queryForCount, current_time)
	err = row.Scan(&count)
	if err != nil {
		return nil, err
	}
	res.Count = count
	return &res, err
}

func (r *taskRepo) ListTasks(req *pbt.ListTasksReq) (*pbt.ListTasksRes, error) {
	var count int64
	offset := (req.Page - 1) * req.Limit
	res := pbt.ListTasksRes{}

	query := `SELECT 
	 		id,
			title,
			deadline,
			summary,
			student_id,
			status,
			created_at,
			type
			FROM student_tasks WHERE deleted_at IS NULL
			LIMIT $1 OFFSET $2`
			
	rows, err := r.db.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		task := pbt.GetTaskRes{}
		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Deadline,
			&task.Summary,
			&task.StudentId,
			&task.Status,
			&task.CreatedAt,
			&task.Type,
		)

		if err != nil {
			return nil, err
		}

		res.Tasks = append(res.Tasks, &task)
	}
	queryForCount := `SELECT COUNT(*) FROM student_tasks WHERE deleted_at IS NULL`
	row := r.db.QueryRow(queryForCount)
	err = row.Scan(&count)
	if err != nil {
		return nil, err
	}
	res.Count = count
	return &res, nil
}