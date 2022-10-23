package models

type Groups struct {
	Name    string `json:"name" db:"name"`
	RCreate bool   `json:"r_create" db:"r_create"`
	RRead   bool   `json:"r_read" db:"r_read"`
	RUpdate bool   `json:"r_update" db:"r_update"`
	RDelete bool   `json:"r_delete" db:"r_delete"`
}

type User struct {
	Name    string `json:"name" db:"name"`
	GroupId int    `json:"group_id" db:"group_id"`
	Pass    string `json:"pass" db:"pass"`
	Email   string `json:"email" db:"email"`
}
