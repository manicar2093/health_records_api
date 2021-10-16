package entities

import "time"

type BoneDensity struct {
	ID          int32     `db:",primary" json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

func (b BoneDensity) Table() string {
	return "BoneDensity"
}
