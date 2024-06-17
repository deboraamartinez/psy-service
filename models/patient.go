package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Patient struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"first_name,omitempty" bson:"first_name,omitempty" validate:"required"`
	LastName  string             `json:"last_name,omitempty" bson:"last_name,omitempty" validate:"required"`
	Age       int                `json:"age,omitempty" bson:"age,omitempty" validate:"required"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Phone     string             `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
}
