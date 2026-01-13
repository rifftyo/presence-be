package repository

import (
	"time"

	"github.com/rifftyo/presence-be/pkg/delivery/http/request"
	"github.com/rifftyo/presence-be/pkg/entity"
)

type AbsenceRepository interface {
	Create(data *entity.AttendanceHistory) error
	Update(id string, data *entity.AttendanceHistory) error
	Get(filter *request.HistoryFilter) ([]*entity.AttendanceHistory, error)
	GetById(userId, historyId string) (*entity.AttendanceHistory, error)
	FindToday(userID string, date time.Time) (*entity.AttendanceHistory, error)
}