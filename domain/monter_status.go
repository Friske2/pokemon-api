package domain

type MonterStats struct {
	ID       int `json:"id"`
	Level    int `json:"level"`
	StatusId int `json:"statusId"`
	MonterId int `json:"monterId"`
}

type MonterStatsRepository interface {
	FindMonterStatsByMonterId(monterId int, m *[]MonterStats) error
	Insert(body []MonterStats) error
	DeleteByMonterId(monterId int) error
	Update(body MonterStats) error
	Delete(id int) error
}
