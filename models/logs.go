package models

type LogType string

const (
	Info    LogType = "info"
	Warning LogType = "warning"
	Error   LogType = "error"
)

type Logs struct {
	GormModel
	JobID   uint
	Message string  `json:"message" form:"message" validate:"required"`
	Type    LogType `json:"type" form:"type" validdate:"required"`
}