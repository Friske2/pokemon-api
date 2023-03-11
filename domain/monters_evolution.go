package domain

type MonterEvolution struct {
	MonterId        int `json:"monterId"`
	EvolutionRuleId int `json:"evolutionRuleId"`
}

type MonterEvolutionRepository interface {
	FindMonterEvolutionByMonterId(monterId int, m *[]MonterEvolution) error
	FindMonterEvolutionByEvolutionRuleId(evolutionRuleId int, m *[]MonterEvolution) error
	Insert(body []MonterEvolution) error
	DeleteByMonterId(monterId int) error
	DeleteByEvolutionRuleId(evolutionRuleId int) error
}
