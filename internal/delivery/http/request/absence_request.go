package request

type AbsenceRequest struct {
	Duration       string   `json:"duration"`
	Status         string    `json:"status"`
	CheckInLat     float64   `json:"check_in_lat" validate:"required"`
	CheckInLng     float64   `json:"check_in_lng" validate:"required"`
	CheckOutLat    float64   `json:"check_out_lat"`
	CheckOutLng    float64   `json:"check_out_lng"`
	CheckInPhoto   string    `json:"check_in_photo" validate:"required"`
	CheckOutPhoto  string    `json:"check_out_photo"`
}