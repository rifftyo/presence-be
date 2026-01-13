package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rifftyo/presence-be/internal/delivery/http/request"
	"github.com/rifftyo/presence-be/internal/delivery/http/response"
	"github.com/rifftyo/presence-be/internal/usecase"
	"github.com/rifftyo/presence-be/utils"
)

type AbsenceHandler struct {
	absenceUseCase usecase.AbsenceUseCase
}

func NewAbsenceHandler(absenceUseCase usecase.AbsenceUseCase) *AbsenceHandler {
	return &AbsenceHandler{absenceUseCase}
}

func (a *AbsenceHandler) CheckIn(c *fiber.Ctx) error {
	var req request.AbsenceRequest

	latStr := c.FormValue("check_in_lat")
	lngStr := c.FormValue("check_in_lng")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid check_in_lat"})
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid check_in_lng"})
	}

	req.CheckInLat = lat
	req.CheckInLng = lng

	fileHeader, err := c.FormFile("check_in_photo")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		defer file.Close()

		fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), fileHeader.Filename)

		filePath, err := utils.SaveFileToSupabase(file, fileName, fileHeader)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to upload to cloud: " + err.Error()})
    }

    req.CheckInPhoto = filePath
	}

	userID := c.Locals("userID").(string)

	err = a.absenceUseCase.CheckIn(&req, userID)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "check in success"})
}

func (a *AbsenceHandler) CheckOut(c *fiber.Ctx) error {
	var req request.AbsenceRequest

	latStr := c.FormValue("check_out_lat")
	lngStr := c.FormValue("check_out_lng")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid check_out_lat"})
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid check_out_lng"})
	}

	req.CheckOutLat = lat
	req.CheckOutLng = lng

	fileHeader, err := c.FormFile("check_out_photo")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		defer file.Close()

		fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), fileHeader.Filename)

		filePath, err := utils.SaveFileToSupabase(file, fileName, fileHeader)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		req.CheckOutPhoto = filePath
	}

	userID := c.Locals("userID").(string)

	err = a.absenceUseCase.CheckOut(&req, userID)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "check out success"})
}

func (a *AbsenceHandler) GetHistory(c *fiber.Ctx) error {
	filter := new(request.HistoryFilter)
	
	userID := c.Locals("userID").(string)
	filter.UserID = userID

    if err := c.QueryParser(filter); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "invalid query parameters",
            "error":   err.Error(),
        })
    }

    resp, err := a.absenceUseCase.GetHistory(filter)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "failed to get history",
            "error":   err.Error(),
        })
    }

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (a *AbsenceHandler) GetHistoryById(c *fiber.Ctx) error {
	userId := c.Locals("userID").(string)

	historyId := c.Params("historyId")
	if historyId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "historyId is required",
        })
	}

	history, err := a.absenceUseCase.GetHistoryById(userId, historyId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "failed to get history",
            "error":   err.Error(),
        })
	}

	resp := response.HistoryDetailResponse{
        ID:            history.ID,
        Date:          history.AttendanceDate.Format("2006-01-02"),
        Day:           history.AttendanceDate.Weekday().String(),
        Status:        history.Status,
        CheckInTime:   history.CheckInTime.Format("15:04"),
        CheckOutTime:  history.CheckOutTime.Format("15:04"),
        CheckInLat:    history.CheckInLat,
        CheckInLng:    history.CheckInLng,
        CheckOutLat:   history.CheckOutLat,
        CheckOutLng:   history.CheckOutLng,
        CheckInPhoto:  history.CheckInPhoto,
        CheckOutPhoto: history.CheckOutPhoto,
        Duration:      history.Duration,
    }

	return c.Status(fiber.StatusOK).JSON(resp)
}