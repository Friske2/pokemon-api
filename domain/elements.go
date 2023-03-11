package domain

type Elements struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type ElementRepository interface {
	FindAll(m *[]Elements) error
	GetById(id int, m *Elements) error
	Insert(m *Elements) error
	Update(body Elements) error
	Delete(id int) error
}

type ElementService interface {
	FindAll() ([]Elements, error)
	GetById(id int) (Elements, error)
	Insert(body Elements) (int, error)
	Update(id int, body Elements) error
	Delete(id int) error
}
