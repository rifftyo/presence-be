package entity

import "time"

type AttendanceHistory struct {
	ID             string `gorm:"primaryKey" json:"id"`
	UserId	string `json:"user_id"`
	User User `gorm:"foreignKey:UserId;references:ID"`
	AttendanceDate time.Time `json:"attendance_date"`
	CheckInTime    time.Time `json:"check_in_time"`
	CheckOutTime   time.Time `json:"check_out_time"`
	Duration string `json:"duration"`
	Status string `json:"status"`
	CheckInLat float64 `json:"check_in_lat"`
	CheckInLng float64 `json:"check_in_lng"`
	CheckOutLat float64 `json:"check_out_lat"`
	CheckOutLng float64 `json:"check_out_lng"`
	CheckInPhoto string `json:"check_in_photo"`
	CheckOutPhoto string `json:"check_out_photo"`
}