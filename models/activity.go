package models

type Activity struct {
	GormModel
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description"`
	ServiceUrl  string `json:"service_url" form:"service_url" validdate:"required"`
}