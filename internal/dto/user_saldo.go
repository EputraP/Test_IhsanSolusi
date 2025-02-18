package dto

type TransactionBody struct {
	NoRekening string `json:"no_rekening" binding:"required"`
	Nominal    string `json:"nominal" binding:"required"`
}
