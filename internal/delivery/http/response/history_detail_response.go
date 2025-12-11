package response

type HistoryDetailResponse struct {
	ID            string  `json:"id"`
	Date          string  `json:"date"`
	Day           string  `json:"day"`
	Status        string  `json:"status"`
	CheckInTime   string  `json:"check_in"`
	CheckOutTime  string  `json:"check_out"`
	CheckInLat    float64 `json:"check_in_lat"`
	CheckInLng    float64 `json:"check_in_lng"`
	CheckOutLat   float64 `json:"check_out_lat"`
	CheckOutLng   float64 `json:"check_out_lng"`
	CheckInPhoto  string  `json:"check_in_photo"`
	CheckOutPhoto string  `json:"check_out_photo"`
	Duration      string  `json:"duration"`
}