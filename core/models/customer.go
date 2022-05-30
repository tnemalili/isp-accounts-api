package models

import "time"

type CustomerType struct {
	ID          int    `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type Customer struct {
	ID           string            `json:"id" bson:"id"`
	Name         string            `json:"name" bson:"name"`
	Surname      string            `json:"surname" bson:"surname"`
	Mobile       string            `json:"mobile" bson:"mobile"`
	Telephone    string            `json:"telephone" bson:"telephone"`
	Email        string            `json:"email" bson:"email"`
	Address      *Address          `json:"address" bson:"address"`
	Status       *Status           `json:"status" bson:"status"`
	CustomerType *CustomerType     `json:"customer_type" bson:"customer_type"`
	Identity     *NationalIdentity `json:"identity" bson:"identity"`
	Created      time.Time         `json:"created" bson:"created"`
	Modified     time.Time         `json:"modified" bson:"modified"`
}

type NationalIdentity struct {
	ID      string   `json:"id" bson:"id"`
	Country *Country `json:"country" bson:"country"`
}

func CreateCustomerType(id int) *CustomerType {
	switch id {
	case 2:
		return &CustomerType{Name: "organization", Description: "Organization Customer", ID: id}
	default:
		return &CustomerType{Name: "individual", Description: "House Hold Customer", ID: id}
	}
}
