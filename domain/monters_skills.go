package domain

type MonterSkill struct {
	MonterId int `json:"monterId"`
	SkillId  int `json:"skillId"`
}

type MonterSkillRepository interface {
	FindMonterSkillByMonterId(monterId int, m *[]MonterSkill) error
	FindMonterSkillBySkillId(skillId int, m *[]MonterSkill) error
	Insert(body []MonterSkill) error
	DeleteByMonterId(monterId int) error
	DeleteBySkillId(skillId int) error
}
