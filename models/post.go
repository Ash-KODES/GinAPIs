package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name        string             `json:"name"`
    Description string             `json:"description"`
    UserID      primitive.ObjectID `json:"user_id"`
}
