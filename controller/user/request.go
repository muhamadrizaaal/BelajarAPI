package user

type LoginRequest struct {
	Hp       string `json:"hp" form:"hp" validate:"required,max=13,min=10"`
	Password string `json:"password" form:"password" validate:"required"`
}

type RegisterRequest struct {
	UserID   uint   `json:"userid" form:"userid"`
	Hp       string `json:"hp" form:"hp" validate:"required,max=13,min=10"`
	Nama     string `json:"nama" form:"nama" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type ActivityRequest struct {
	UserID   uint   `json:"userid" form:"userid"`
	Activity string `json:"activity" form:"activity"`
}

type AllActivityRequest struct {
	UserID uint `json:"userid" form:"userid"`
}
