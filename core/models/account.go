package models

import (
	"fmt"
	"os"
	"time"
)

type CreateAccountModel struct {
	CustomerId string `json:"customer_id" bson:"customer_id"`
}

type Account struct {
	ID            string    `json:"id" bson:"id"`
	AccountHolder interface{} `json:"account_holder" bson:"account_holder"`
	Status        *Status   `json:"status" bson:"status"`
	CustomerID    string    `json:"customer_id" bson:"customer_id"`
	Balance       *Amount   `json:"balance" bson:"balance"`
	Created       time.Time `json:"created" bson:"created"`
	Modified      time.Time `json:"modified" bson:"modified"`
}

func AddAmount(amount float64, balance float64) float64 {
	return balance + amount
}

func WithdrawAmount(amount float64, balance float64) float64 {
	return balance - amount
}

func Overdraft(amount float64, balance float64) bool {
	return balance-amount < 0
}

func NewDisplayName(amount float64) string {

	curr := os.Getenv("CURRENCY")
	return fmt.Sprintf("%v%0.2f", curr, amount)
}
