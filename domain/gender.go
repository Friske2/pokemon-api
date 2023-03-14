package domain

type Gender struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GenderRepository interface {
	FindAll(m *[]Gender) error
	GetById(id int, m *Gender) error
	Insert(name string) error
	Update(id int, name string) error
	Delete(id int) error
}

type GenderService interface {
	FindAll() ([]Gender, error)
	GetById(id int) (Gender, error)
	Insert(name string) error
	Update(id int, name string) error
	Delete(id int) error
}
