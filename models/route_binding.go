package models

type IdBinding struct {
	Id string `uri:"id" binding:"required"`
}
