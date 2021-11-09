package models

type User struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Job     string `json:"job"`
}
