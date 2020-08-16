package database

import (
	"testTask/internal/pkg/models"
)

type IDatabase interface {
	Get([]string) ([]models.Line, error)
	Set(models.Line) error
	Ping() error
}
