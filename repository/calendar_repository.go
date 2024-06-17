package repository

import (
	"context"
	"psy-service/config"
	"psy-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CalendarRepository struct{}

func (r *CalendarRepository) Create(event *models.CalendarEvent) error {
	collection := config.DB.Collection("calendar_events")
	_, err := collection.InsertOne(context.Background(), event)
	return err
}

func (r *CalendarRepository) Update(id primitive.ObjectID, event *models.CalendarEvent) error {
	collection := config.DB.Collection("calendar_events")
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": event})
	return err
}

func (r *CalendarRepository) Delete(id primitive.ObjectID) error {
	collection := config.DB.Collection("calendar_events")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (r *CalendarRepository) GetAll() ([]models.CalendarEvent, error) {
	var events []models.CalendarEvent
	collection := config.DB.Collection("calendar_events")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return events, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var event models.CalendarEvent
		cursor.Decode(&event)
		events = append(events, event)
	}

	return events, cursor.Err()
}
