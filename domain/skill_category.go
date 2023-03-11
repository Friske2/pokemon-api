package domain

type SkillCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SkillCategoryRepository interface {
	FindAll(m *[]SkillCategory) error
	GetById(id int, m *SkillCategory) error
	Insert(name string) error
	Update(id int, name string) error
	Delete(id int) error
}
