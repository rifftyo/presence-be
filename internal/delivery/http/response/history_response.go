package response

type AttendanceSummary struct {
	TotalDays    int `json:"total_days"`
	PresentCount int `json:"present"`
	LateCount    int `json:"late"`
}

type AttendanceDetail struct {
	ID           string `json:"id"`
	Date         string `json:"date"`
	Day          string `json:"day"`
	Status       string `json:"status"`
	CheckInTime  string `json:"check_in"`
	CheckOutTime string `json:"check_out"`
	Duration     string `json:"check_duration"`
}

type AttendanceHistoryResponse struct {
	Summary AttendanceSummary  `json:"summary"`
	History []AttendanceDetail `json:"history"`
}