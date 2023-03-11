package domain

type Skills struct {
	ID              int `json:"id"`
	Level           int `json:"level"`
	ElementId       int `json:"elementId"`
	SkillCategoryId int `json:"skillCategoryId"`
	Power           int `json:"power"`
	Accuracy        int `json:"accuracy"`
	SkillLearnId    int `json:"skillLearnId"`
}
