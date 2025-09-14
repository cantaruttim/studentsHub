package repository

import (
	"database/sql"
	"fmt"
	"forms-activities/model"
)

type FormsRepository struct {
	connection *sql.DB
}

func NewFormsRepository(
	connection *sql.DB,
) FormsRepository {
	return FormsRepository{
		connection: connection,
	}
}

func (pr *FormsRepository) GetForms() ([]model.Forms, error) {
	query := "select registrationNumber, name, email, module, questionOne, questionTwo, SentAt  from activities"

	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Forms{}, err
	}

	var activitiesList []model.Forms
	var activitiesObj model.Forms

	for rows.Next() {
		err = rows.Scan(
			&activitiesObj.RegistrationNumber,
			&activitiesObj.Name,
			&activitiesObj.Email,
			&activitiesObj.Module,
			&activitiesObj.QuestionOne,
			&activitiesObj.QuestionTwo,
			&activitiesObj.SentAt,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Forms{}, err
		}

		activitiesList = append(activitiesList, activitiesObj)
	}
	rows.Close()
	return activitiesList, nil
}
