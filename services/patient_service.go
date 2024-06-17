package services

import (
	"psy-service/models"
	"psy-service/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PatientService struct {
	Repo *repository.PatientRepository
}

func (s *PatientService) CreatePatient(patient *models.Patient) error {
	return s.Repo.Create(patient)
}

func (s *PatientService) UpdatePatient(id string, patient *models.Patient) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.Repo.Update(oid, patient)
}

func (s *PatientService) DeletePatient(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.Repo.Delete(oid)
}

func (s *PatientService) GetAllPatients() ([]models.Patient, error) {
	return s.Repo.GetAll()
}
