package domain

type MonterElement struct {
	MonterId  int `json:"monterId"`
	ElementId int `json:"elementId"`
}

type MonterElementRepository interface {
	FindMonterElementByMonterId(monterId int, m *[]MonterElement) error
	FindMonterElementByElementId(elementId int, m *[]MonterElement) error
	Insert(body []MonterElement) error
	DeleteByMonterId(monterId int) error
	DeleteByElementId(elementId int) error
}
