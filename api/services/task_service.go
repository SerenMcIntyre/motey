package services

import (
	"context"
	db "motey-api/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskService struct {
	DB *db.Queries
}

type CreateTaskRequest struct {
	Name      string `json:"name"`
	Background     string `json:"background"`
	Sticker        string `json:"sticker"`
	IsMeasured     bool   `json:"is_measured"`
	MeasurementUnit string `json:"measurement_unit"`
	StickerValue   int    `json:"sticker_value"`
	UserID         string `json:"user_id"`
}

func NewTaskService(db *db.Queries) *TaskService {
	return &TaskService{DB: db}
}

func (s *TaskService) GetTaskByID(ctx context.Context, id string) (*db.Task, error) {
	uuid, err := StringToPgUuid(id)
	if err != nil {
		return nil, nil
	}

	task, err := s.DB.GetTaskByID(ctx, *uuid)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *TaskService) GetTasks(ctx context.Context, userid string) ([]db.Task, error) {
	uuid, err := StringToPgUuid(userid)
	if err != nil {
		return nil, err
	}

	tasks, err := s.DB.GetUserTasks(ctx, *uuid)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) CreateTask(ctx context.Context, insertRequest CreateTaskRequest) (*db.Task, error) {
	insert, err := CreateTaskRequestToInsertTaskParams(insertRequest)
	if err != nil {
		return nil, err
	}

	newTask, err := s.DB.InsertTask(ctx, *insert)
	if err != nil {
		return nil, err
	}

	return &newTask, nil
}

func CreateTaskRequestToInsertTaskParams(request CreateTaskRequest) (*db.InsertTaskParams, error){
	userid, err := StringToPgUuid(request.UserID)
	if err != nil {
		return nil, err
	}

	return &db.InsertTaskParams{
		ID:             pgtype.UUID{Bytes: uuid.New(), Valid: true},
		TitleName:      request.Name,
		Background:     request.Background,
		Sticker:        []byte(request.Sticker),
		IsMeasured:     request.IsMeasured,
		MeasurementUnit: pgtype.Text{String: request.MeasurementUnit, Valid: true},
		StickerValue:   pgtype.Int4{Int32: int32(request.StickerValue), Valid: true},
		UserID:         *userid,
	}, nil
}