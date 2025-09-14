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

func (fr *FormsRepository) GetForms() ([]model.Forms, error) {
	query := "select registrationNumber, name, email, module, questionOne, questionTwo, SentAt  from activities"

	rows, err := fr.connection.Query(query)
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

func (fr *FormsRepository) CreateForms(form model.Forms) error {

	_, err := fr.connection.Exec(`
		INSERT INTO activities (
			RegistrationNumber,
			Name,
			Email,
			Module,
			QuestionOne,
			QuestionTwo,
			SentAt
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`,
		form.RegistrationNumber,
		form.Name,
		form.Email,
		form.Module,
		form.QuestionOne,
		form.QuestionTwo,
		form.SentAt,
	)

	if err != nil {
		fmt.Println("Erro ao inserir no banco:", err)
		return err
	}

	fmt.Println("✅ Formulário inserido com sucesso!")
	return nil
}

func (fr *FormsRepository) FindById(RegistrationNumber string) (*model.Forms, error) {

	query, err := fr.connection.Prepare(
		"select * from activities" +
			"where RegistrationNumber = $1",
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var form model.Forms

	err = query.QueryRow(RegistrationNumber).Scan(
		&form.RegistrationNumber,
		&form.Name,
		&form.Email,
		&form.Module,
		&form.QuestionOne,
		&form.QuestionTwo,
		&form.SentAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	query.Close()
	return &form, nil
}
