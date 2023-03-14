package domain

type SkillLearn struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SkillLearnRepository interface {
	FindAll(m *[]SkillLearn) error
	GetById(id int, m *SkillLearn) error
	Insert(name string) error
	Update(id int, name string) error
	Delete(id int) error
}
