package bu

import (
	"encoding/json"
	"errors"
	"fmt"
	kafkaconfig "go-company/kafkaConfig"
	"go-company/models"
	"time"

	
)

func CreateCompany(company *models.Company) (*models.Company, error) {

	dbconn := models.GetDbHandle()
	if dbconn == nil {
		return nil, errors.New(
			"failed to get db connectionn")
	}

	//check if company with same name exists
	dbCompany, err := dbconn.GetCompany(company.Name)

	if err != nil {
		//return nil,err
	}

	if dbCompany != nil && dbCompany.Name == company.Name {
		return nil, errors.New(fmt.Sprintf("company %s already exists", company.Name))
	}


	dbCompany, err = dbconn.CreateCompany(company)
	if err != nil {
		return nil, err
	}

	//write to the kafka
	evt := kafkaconfig.MessagePayload{
		EventType: kafkaconfig.CompanyCreated,
		 Payload: *dbCompany,
		 Time: time.Now(),
	}
	payload, _ := json.Marshal(evt)
	kafkaconfig.WriteToProducer(payload)
	return dbCompany, nil

}

func GetCompany(companyName string) (*models.Company, error) {

	dbconn := models.GetDbHandle()
	if dbconn == nil {
		return nil, errors.New(
			"failed to get db connectionn")
	}

	//check if company with same name exists

	dbCompany, err := dbconn.GetCompany(companyName)
	if err != nil {
		return nil, err
	}

	return dbCompany, nil

}

func DeleteCompany(companyName string) error {

	dbconn := models.GetDbHandle()
	if dbconn == nil {
		return errors.New(
			"failed to get db connectionn")
	}

	//check if company with same name exists

	err := dbconn.DeleteCompany(companyName)
	if err != nil {
		return err
	}
	//write to the kafka
	evt := kafkaconfig.MessagePayload{
		EventType: kafkaconfig.CompanyDeleted,
		 Payload: *&models.Company{Name: companyName},
		 Time: time.Now(),
	}
	payload, _ := json.Marshal(evt)
	kafkaconfig.WriteToProducer(payload)

	return nil

}

func UpadteCompany(company *models.Company) (*models.Company, error) {

	dbconn := models.GetDbHandle()
	if dbconn == nil {
		return nil, errors.New(
			"failed to get db connectionn")
	}

	//check if company with same name exists
	dbCompany, err := dbconn.GetCompany(company.Name)

	if err != nil {
		return nil, err
	}

	if dbCompany != nil && dbCompany.Name == "" {
		return nil, errors.New(fmt.Sprintf("company %s does not exists", company.Name))
	}

	company.ID = dbCompany.ID

	dbCompany, err = dbconn.UpdateCompany(company)
	if err != nil {
		return nil, err
	}
	//write to the kafka
	evt := kafkaconfig.MessagePayload{
		EventType: kafkaconfig.CompanyUpdated,
		 Payload: *dbCompany,
		 Time: time.Now(),
	}
	payload, _ := json.Marshal(evt)
	kafkaconfig.WriteToProducer(payload)
	return dbCompany, nil

}
