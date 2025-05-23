// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Frequency struct {
	ID               pgtype.UUID
	Name             string
	FrequencyGroupID pgtype.UUID
}

type FrequencyGroup struct {
	ID   pgtype.UUID
	Name string
}

type Task struct {
	ID              pgtype.UUID
	TitleName       string
	Background      string
	Sticker         []byte
	IsMeasured      bool
	MeasurementUnit pgtype.Text
	StickerValue    pgtype.Int4
	UserID          pgtype.UUID
}

type TaskNotification struct {
	ID          pgtype.UUID
	Text        string
	Time        pgtype.Time
	TaskID      pgtype.UUID
	FrequencyID pgtype.UUID
}

type User struct {
	ID         pgtype.UUID
	Name       string
	Background string
}
