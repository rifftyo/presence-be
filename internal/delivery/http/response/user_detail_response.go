package response

type UserDetailResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Photo      string `json:"photo"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	Telephone  string `json:"telephone"`
	Department string `json:"department"`
}