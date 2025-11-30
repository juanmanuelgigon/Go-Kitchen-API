package responses

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserInfo struct {
	Codigo   string             `json:codigo`
	Email    string             `json:email`
	Username string             `json:username`
	ID       primitive.ObjectID `json:id`
}
