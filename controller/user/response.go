package user

type LoginResponse struct {
	// UserID uint   `json:"userid"`
	Hp    string `json:"hp"`
	Nama  string `json:"nama"`
	Token string `json:"token"`
}
