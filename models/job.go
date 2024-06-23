package models

type Jobs struct {
	GormModel
	SchedulerID  uint
	ActivityID   uint
	ActivityData map[string]interface{} `gorm:"serializer:json" json:"activity_data" form:"activity_data" validate:"required"`
	Name         string                 `json:"name" form:"name" validate:"required"`
	Description  string                 `json:"description" form:"description"`
	Url          string                 `json:"url" form:"url" validate:"required"`
	Order        int64                  `json:"order" form:"order" validate:"required"`
	IsActice     bool                   `json:"is_active" form:"is_active" validate:"required"`

	// parent
	Scheduler *Scheduler
	Activity  *Activity

	// child
	Logs []Logs `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"logs"`
}