package models

type Scheduler struct {
	GormModel
	UserID      uint
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description"`
	Interval    string `json:"interval" form:"interval" validate:"required"`
	IsActice    bool   `json:"is_active" form:"is_active" validate:"required"`

	Jobs []Jobs `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"jobs"`
}