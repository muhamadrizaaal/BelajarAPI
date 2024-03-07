package todo

type ActivityRequest struct {
	Hp       string   `json:"hp" form:"hp"`
	Activity string `json:"activity" form:"activity"`
}

type AllActivityRequest struct {
	Hp string `json:"hp" form:"hp"`
}
