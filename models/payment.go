package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PatientID primitive.ObjectID `json:"patient_id,omitempty" bson:"patient_id,omitempty" validate:"required"`
	Amount    float64            `json:"amount,omitempty" bson:"amount,omitempty" validate:"required"`
	Paid      bool               `json:"paid,omitempty" bson:"paid,omitempty"`
	Date      string             `json:"date,omitempty" bson:"date,omitempty" validate:"required"`
}
