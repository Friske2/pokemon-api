package domain

import "time"

type Monters struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CreatedDate time.Time `json:"created_date"`
	CreatedBy   string    `json:"created_by"`
	Enabled     bool      `json:"enabled"`
}
