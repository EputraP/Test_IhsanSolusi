package dto

type CreateUserBody struct {
	Nama string `json:"nama" binding:"required"`
	NIK  string `json:"nik" binding:"required"`
	NoHP string `json:"no_hp" binding:"required"`
}
type CreateUserResponse struct {
	NoRekening string `json:"no_rekening" binding:"required"`
}
