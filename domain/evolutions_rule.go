package domain

type EvolutionsRule struct {
	ID    int    `json:"id"`
	Level int    `json:"level"`
	Item  string `json:"item"`
}

type EvolutionsRuleRepository interface {
	FindAll(m *[]EvolutionsRule) error
	GetById(id int, m *EvolutionsRule) error
	Insert(m *EvolutionsRule) error
	Update(body EvolutionsRule) error
	Delete(id int) error
}
