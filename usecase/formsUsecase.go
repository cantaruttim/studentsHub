package usecase

import (
	"forms-activities/model"
	"forms-activities/repository"
)

type FormsUsecase struct {
	// Repository
	repository repository.FormsRepository
}

func NewFormsUsecase(repo repository.FormsRepository) FormsUsecase {
	return FormsUsecase{
		repository: repo,
	}
}

// Implement business logic here in the future

// GET
func (fu *FormsUsecase) GetForms() ([]model.Forms, error) {
	return fu.repository.GetForms()
}

// POST
func (fu *FormsUsecase) PostForms(form model.Forms) (model.Forms, error) {
	err := fu.repository.CreateForms(form)

	if err != nil {
		return model.Forms{}, err
	}

	return form, nil
}

// FIND BY ID
func (fu *FormsUsecase) FindById(RegistrationNumber int) (*model.Forms, error) {

	form, err := fu.repository.FindById(RegistrationNumber)
	if err != nil {
		return nil, err
	}

	return form, nil
}
