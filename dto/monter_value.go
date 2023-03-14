package dto

type MonterValue struct {
	Name    string `json:"name"`
	Enabled *bool  `json:"enabled"`
}

type GetMonters struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}
