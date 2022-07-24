package sdkcm

import "time"

type SQLModel struct {
	ID        uint32    `json:"id" gorm:"id,PRIMARY_KEY"`
	CreatedAt *JSONTime `json:"created_at,omitempty"`
	UpdatedAt *JSONTime `json:"updated_at,omitempty"`
}

func NewSQLModel() *SQLModel {
	t := JSONTime(time.Now())

	return &SQLModel{
		CreatedAt: &t,
		UpdatedAt: &t,
	}
}

func (m *SQLModel) FullFill() {
	t := JSONTime(time.Now())

	if m.UpdatedAt == nil {
		m.UpdatedAt = &t
	}
}
