package repository

import (
	"time"

	"github.com/rifftyo/presence-be/internal/delivery/http/request"
	"github.com/rifftyo/presence-be/internal/entity"
	"gorm.io/gorm"
)

type absenceRepository struct {
	DB *gorm.DB
}

func NewAbsenceRepository(db *gorm.DB) AbsenceRepository {
	return &absenceRepository{db}
}

func (a *absenceRepository) Create(data *entity.AttendanceHistory) error {
	return a.DB.Create(data).Error
}

func (a *absenceRepository) Update(id string, data *entity.AttendanceHistory) error {
	return a.DB.Model(data).
		Where("id = ?", id).
		Updates(data).Error
}

func (a *absenceRepository) Get(filter *request.HistoryFilter) ([]*entity.AttendanceHistory, error) {
	var histories []*entity.AttendanceHistory
	db := a.DB.Model(&entity.AttendanceHistory{})

	db = db.Where("user_id = ?", filter.UserID)

	if filter.Status != nil {
		db = db.Where("status = ?", *filter.Status)
	}

	now := time.Now()
	switch filter.Period {
	case "this_month":
		start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		end := start.AddDate(0, 1, 0).Add(-time.Nanosecond)
		db = db.Where("attendance_date BETWEEN ? AND ?", start, end)

	case "last_month":
		start := time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, now.Location())
		end := start.AddDate(0, 1, 0).Add(-time.Nanosecond)
		db = db.Where("attendance_date BETWEEN ? AND ?", start, end)

	case "last_3_months":
		start := now.AddDate(0, -3, 0)
		end := now
		db = db.Where("attendance_date BETWEEN ? AND ?", start, end)

	}

	err := db.Order("attendance_date DESC").Find(&histories).Error
	if err != nil {
		return nil, err
	}

	return histories, nil
}

func (a *absenceRepository) GetById(userId string, historyId string) (*entity.AttendanceHistory, error) {
	var history entity.AttendanceHistory

	err := a.DB.Where("user_id = ?", userId).Where("id = ?", historyId).First(&history).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &history, nil
}

func (a *absenceRepository) FindToday(userID string, date time.Time) (*entity.AttendanceHistory, error) {
	var record entity.AttendanceHistory

	err := a.DB.Where("user_id = ?", userID).Where("DATE(attendance_date)= DATE(?)", date).First(&record).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &record, nil
}
