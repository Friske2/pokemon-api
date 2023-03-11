package domain

type Status struct {
	ID      int `json:"id"`
	Hp      int `json:"Hp"`
	Attack  int `json:"Attack"`
	Defense int `json:"Defense"`
	SpAtk   int `json:"SpAtk"`
	SpDef   int `json:"SpDef"`
}
