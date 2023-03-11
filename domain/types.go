package domain

type Types struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TypesRepository interface {
	FindAll(m *[]Types) error
	GetById(id int, m *Types) error
	Insert(name string) error
	Update(id int, name string) error
	Delete(id int) error
}
