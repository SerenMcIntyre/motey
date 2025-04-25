package services

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func StringToPgUuid(id string) (*pgtype.UUID, error) {
	parseID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	pgUUID := pgtype.UUID{
		Bytes: parseID,
		Valid: true,
	}

	return &pgUUID, nil
}