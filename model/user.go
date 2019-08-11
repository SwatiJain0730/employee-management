package model

//User ...
type User struct {
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	AuthType string `form:"authType" json:"authType" xml:"authType" binding:"required"`
}
