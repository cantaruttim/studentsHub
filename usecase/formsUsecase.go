package usecase

import (
	"forms-activities/model"
)

type FormsUsecase struct {
	// Repository
}

func NewFormsUsecase() FormsUsecase {
	return FormsUsecase{}
}

// Implement business logic here in the future
func (u *FormsUsecase) GetForms() ([]model.Forms, error) {
	return []model.Forms{}, nil
}
