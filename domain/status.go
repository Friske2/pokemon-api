package domain

type Status struct {
	ID      int `json:"id"`
	Hp      int `json:"Hp"`
	Attack  int `json:"Attack"`
	Defense int `json:"Defense"`
	SpAtk   int `json:"SpAtk"`
	SpDef   int `json:"SpDef"`
}

type StatusRepository interface {
	FindAll(m *[]Status) error
	GetById(id int, m *Status) error
	Insert(m Status) error
	Update(body Status) error
	Delete(id int) error
}
