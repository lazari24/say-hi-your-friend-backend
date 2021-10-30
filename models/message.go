package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	FriendName string             `json:"friendName,omitempty" bson:"friendName,omitempty"`
	Message    string             `json:"message,omitempty" bson:"message,omitempty"`
}
