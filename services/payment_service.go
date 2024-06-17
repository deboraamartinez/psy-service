package services

import (
	"psy-service/models"
	"psy-service/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentService struct {
	Repo *repository.PaymentRepository
}

func (s *PaymentService) CreatePayment(payment *models.Payment) error {
	return s.Repo.Create(payment)
}

func (s *PaymentService) UpdatePayment(id string, payment *models.Payment) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.Repo.Update(oid, payment)
}

func (s *PaymentService) DeletePayment(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.Repo.Delete(oid)
}

func (s *PaymentService) GetAllPayments() ([]models.Payment, error) {
	return s.Repo.GetAll()
}
