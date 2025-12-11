package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rifftyo/presence-be/internal/delivery/http/request"
	"github.com/rifftyo/presence-be/internal/delivery/http/response"
	"github.com/rifftyo/presence-be/internal/entity"
	"github.com/rifftyo/presence-be/internal/repository"
)

type absenceUseCase struct {
	absenceRepo repository.AbsenceRepository
}

func NewAbsenceUseCase(absenceRepo repository.AbsenceRepository) AbsenceUseCase {
	return &absenceUseCase{absenceRepo}
}

func (a *absenceUseCase) CheckIn(req *request.AbsenceRequest, userId string) error {
	now := time.Now()

	attendanceDate := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)

	checkInTime := now

	todayRecord, _ := a.absenceRepo.FindToday(userId, now)
	if todayRecord != nil && !todayRecord.CheckInTime.IsZero() {
		return errors.New("already checked in today")
	}

	cutOff := time.Date(
		now.Year(), now.Month(), now.Day(),
		7, 30, 0, 0,
		now.Location(),
	)

	status := "Present"
	if checkInTime.After(cutOff) {
		status = "Late"
	}

	data := &entity.AttendanceHistory{
		ID:             uuid.New().String(),
		UserId:         userId,
		AttendanceDate: attendanceDate,
		CheckInTime:    checkInTime,
		CheckInLat:     req.CheckInLat,
		CheckInLng:     req.CheckInLng,
		CheckInPhoto:   req.CheckInPhoto,
		Status:         status,
	}

	return a.absenceRepo.Create(data)
}

func (a *absenceUseCase) CheckOut(req *request.AbsenceRequest, userId string) error {
	now := time.Now()

	todayRecord, err := a.absenceRepo.FindToday(userId, now)
	if err != nil {
		return err
	}

	if todayRecord == nil || todayRecord.CheckInTime.IsZero() {
		return errors.New("you haven't checked in today")
	}

	if !todayRecord.CheckOutTime.IsZero() {
		return errors.New("you have already checked out today")
	}

	durationHours := now.Sub(todayRecord.CheckInTime).Hours()
	durationStr := fmt.Sprintf("%.2f hours", durationHours)

	data := &entity.AttendanceHistory{
		CheckOutTime:  now,
		Duration:      durationStr,
		CheckOutLat:   req.CheckOutLat,
		CheckOutLng:   req.CheckOutLng,
		CheckOutPhoto: req.CheckOutPhoto,
	}

	return a.absenceRepo.Update(todayRecord.ID, data)
}

func (a *absenceUseCase) GetHistory(req *request.HistoryFilter) (*response.AttendanceHistoryResponse, error) {
	history, err := a.absenceRepo.Get(req)
	if err != nil {
		return nil, err
	}

	summary := response.AttendanceSummary{}
	details := make([]response.AttendanceDetail, 0, len(history))

	for _, h := range history {
		switch h.Status {
		case "Present":
			summary.PresentCount++
		case "Late":
			summary.LateCount++
		}

		checkIn := ""
		checkOut := ""
		duration := ""
		if !h.CheckInTime.IsZero() {
			checkIn = h.CheckInTime.Format("15:04")
		}
		if !h.CheckOutTime.IsZero() {
			checkOut = h.CheckOutTime.Format("15:04")
			duration = h.Duration
		}

		details = append(details, response.AttendanceDetail{
			ID: h.ID,
			Date:         h.CheckInTime.Format("2006-01-02"),
			Day:          h.CheckInTime.Weekday().String(),
			Status:       h.Status,
			CheckInTime:  checkIn,
			CheckOutTime: checkOut,
			Duration:     duration,
		})
	}

	summary.TotalDays = len(history)

	return &response.AttendanceHistoryResponse{
		Summary: summary,
		History: details,
	}, nil
}

func (a *absenceUseCase) GetHistoryById(userId string, historyId string) (*entity.AttendanceHistory, error) {
	return a.absenceRepo.GetById(userId, historyId)
}