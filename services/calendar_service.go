package services

import (
	"psy-service/models"
	"psy-service/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CalendarService struct {
	Repo *repository.CalendarRepository
}

func (s *CalendarService) CreateCalendarEvent(event *models.CalendarEvent) error {
	return s.Repo.Create(event)
}

func (s *CalendarService) UpdateCalendarEvent(id string, event *models.CalendarEvent) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.Repo.Update(oid, event)
}

func (s *CalendarService) DeleteCalendarEvent(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.Repo.Delete(oid)
}

func (s *CalendarService) GetAllCalendarEvents() ([]models.CalendarEvent, error) {
	return s.Repo.GetAll()
}
