package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Jobdesk struct {
    Id        primitive.ObjectID `json:"id,omitempty"`
    Name      string             `json:"name,omitempty" validate:"required"`
    Desc      string             `json:"desc,omitempty" validate:"required"`
    Duration  int                `json:"duration,omitempty" validate:"required"`
}