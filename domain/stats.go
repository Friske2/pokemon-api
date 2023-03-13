package domain

type Stats struct {
	ID      int `json:"id"`
	Hp      int `json:"Hp"`
	Attack  int `json:"Attack"`
	Defense int `json:"Defense"`
	SpAtk   int `json:"SpAtk"`
	SpDef   int `json:"SpDef"`
}

type StatsRepository interface {
	FindAll(m *[]Stats) error
	GetById(id int, m *Stats) error
	Insert(m Stats) error
	Update(body Stats) error
	Delete(id int) error
}
