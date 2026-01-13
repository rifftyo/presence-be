package usecase

import (
	"github.com/rifftyo/presence-be/pkg/delivery/http/request"
	"github.com/rifftyo/presence-be/pkg/delivery/http/response"
	"github.com/rifftyo/presence-be/pkg/entity"
)

type AbsenceUseCase interface {
	CheckIn(req *request.AbsenceRequest, userId string) error
	CheckOut(req *request.AbsenceRequest, userId string) error
	GetHistory(req *request.HistoryFilter) (*response.AttendanceHistoryResponse, error)
	GetHistoryById(userId string, historyId string) (*entity.AttendanceHistory, error) 
}