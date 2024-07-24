package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    PostID      primitive.ObjectID `json:"post_id"`
    Description string             `json:"description"`
    UserID      primitive.ObjectID `json:"user_id"`
}
