package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CalendarEvent struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty" validate:"required"`
	PatientID primitive.ObjectID `json:"patient_id,omitempty" bson:"patient_id,omitempty" validate:"required"`
	Date      string             `json:"date,omitempty" bson:"date,omitempty" validate:"required"`
}
