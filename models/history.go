package models

type History struct {
	Id uint `json:"id"`
	Topup string `json:"topup"`
	Transfer string `json:"transfer"`
	Received string `json:"received"`
	To string `json:"to"`
	From string `json:"from"`
	UserId string `json:"user_id"`
}

type CheckSaldo struct {
    TotalTopup int
    TotalTransfer int
    TotalReceived int
}