package handler

type TodoRequest struct {
	Kegiatan string `json:"kegiatan" form:"kegiatan"`
}
