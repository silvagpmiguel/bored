package model

import (
	"strconv"

	"gorm.io/gorm"
)

// Activity structure model
type Activity struct {
	gorm.Model
	Activity      string  `json:"activity,omitempty"`
	Accessibility float64 `json:"accessibility,omitempty"`
	Type          string  `json:"type,omitempty"`
	Participants  int     `json:"participants,omitempty"`
	Price         float64 `json:"price,omitempty"`
	Link          string  `json:"link,omitempty"`
	Key           string  `json:"key,omitempty"`
}

// NewActivity returns an empty Activity struct Reference
func NewActivity() *Activity {
	return &Activity{}
}

// WithID set Activity ID
func (act *Activity) WithID(_id string) *Activity {
	id, err := strconv.ParseUint(_id, 10, 64)

	if err == nil {
		act.ID = uint(id)
	}

	return act
}

// WithType set Activity Type
func (act *Activity) WithType(actType string) *Activity {
	act.Type = actType
	return act
}
