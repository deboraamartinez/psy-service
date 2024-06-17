package repository

import (
	"context"
	"psy-service/config"
	"psy-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentRepository struct{}

func (r *PaymentRepository) Create(payment *models.Payment) error {
	collection := config.DB.Collection("payments")
	_, err := collection.InsertOne(context.Background(), payment)
	return err
}

func (r *PaymentRepository) Update(id primitive.ObjectID, payment *models.Payment) error {
	collection := config.DB.Collection("payments")
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": payment})
	return err
}

func (r *PaymentRepository) Delete(id primitive.ObjectID) error {
	collection := config.DB.Collection("payments")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (r *PaymentRepository) GetAll() ([]models.Payment, error) {
	var payments []models.Payment
	collection := config.DB.Collection("payments")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return payments, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var payment models.Payment
		cursor.Decode(&payment)
		payments = append(payments, payment)
	}

	return payments, cursor.Err()
}
