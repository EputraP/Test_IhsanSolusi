package dto

type TransactionBody struct {
	NoRekening string `json:"no_rekening" binding:"required"`
	Nominal    string `json:"nominal" binding:"required"`
}
type CurrentBalanceResponse struct {
	SaldoSaatIni string `json:"saldo_saat_ini" binding:"required"`
}
