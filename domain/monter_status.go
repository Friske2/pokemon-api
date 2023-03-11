package domain

type MonterStatus struct {
	ID       int `json:"id"`
	Level    int `json:"level"`
	StatusId int `json:"statusId"`
	MonterId int `json:"monterId"`
}
