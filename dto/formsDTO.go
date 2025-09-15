package dto

import (
	"forms-activities/model"
	"time"
)

type FormsDTO struct {
	RegistrationNumberDTO string    `json:"registrationNumber"`
	NameDTO               string    `json:"name"`
	EmailDTO              string    `json:"email"`
	ModuleDTO             string    `json:"module"`
	QuestionOneDTO        string    `json:"questionOne"`
	QuestionTwoDTO        string    `json:"questionTwo"`
	SentAtDTO             time.Time `json:"sentAt"`
}

func (dto *FormsDTO) ToModel() model.Forms {
	return model.Forms{
		RegistrationNumber: dto.RegistrationNumberDTO,
		Name:               dto.NameDTO,
		Email:              dto.EmailDTO,
		Module:             dto.ModuleDTO,
		QuestionOne:        dto.QuestionOneDTO,
		QuestionTwo:        dto.QuestionTwoDTO,
		SentAt:             dto.SentAtDTO,
	}
}
