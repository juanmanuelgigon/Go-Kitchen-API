package dto

import (
	"TPFINAL-GINCITO/clients/responses"
	"TPFINAL-GINCITO/utils"
)

type User struct {
	Codigo   string `json:codigo`
	Email    string `json:email`
	Username string `json:username`
	ID       string `json:id`
}

func NewUser(userInfo *responses.UserInfo) *User {
	return &User{
		Codigo:   userInfo.Codigo,
		Email:    userInfo.Email,
		Username: userInfo.Username,
		ID:       utils.GetStringIDFromObjectID(userInfo.ID),
	}
}
