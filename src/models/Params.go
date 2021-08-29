package models

type Params struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}
