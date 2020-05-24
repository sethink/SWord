package models

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	En        string `json:"en"`
	GroupId   string `json:"group_id"`
	AuthLevel string `json:"auth_level"`
}
