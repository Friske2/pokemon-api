package domain

type MonterStatus struct {
	ID       int `json:"id"`
	Level    int `json:"level"`
	StatusId int `json:"statusId"`
	MonterId int `json:"monterId"`
}

type MonterStatusRepository interface {
	FindMonterStatusByMonterId(monterId int, m *[]MonterStatus) error
	Insert(body []MonterStatus) error
	DeleteByMonterId(monterId int) error
	Update(body MonterStatus) error
	Delete(id int) error
}
