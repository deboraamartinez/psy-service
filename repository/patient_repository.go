package repository

import (
	"context"
	"psy-service/config"
	"psy-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PatientRepository struct{}

func (r *PatientRepository) Create(patient *models.Patient) error {
	collection := config.DB.Collection("patients")
	_, err := collection.InsertOne(context.Background(), patient)
	return err
}

func (r *PatientRepository) Update(id primitive.ObjectID, patient *models.Patient) error {
	collection := config.DB.Collection("patients")
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": patient})
	return err
}

func (r *PatientRepository) Delete(id primitive.ObjectID) error {
	collection := config.DB.Collection("patients")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (r *PatientRepository) GetAll() ([]models.Patient, error) {
	var patients []models.Patient
	collection := config.DB.Collection("patients")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return patients, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var patient models.Patient
		cursor.Decode(&patient)
		patients = append(patients, patient)
	}

	return patients, cursor.Err()
}
