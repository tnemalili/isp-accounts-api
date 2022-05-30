package models

import (
	"isp.accounts.api/core/config"
	"strings"
	"time"
)

type Country struct {
	ISOCode     string `json:"iso_code" bson:"iso_code"`
	Name        string `json:"name" bson:"name"`
	CountryCode string `json:"country_code" bson:"country_code"`
}

type Address struct {
	Province    string `json:"province" bson:"province"`
	City        string `json:"city" bson:"city"`
	Township    string `json:"township" bson:"township"`
	HouseNumber string `json:"house_number" bson:"house_number"`
	StreetName  string `json:"street_name" bson:"street_name"`
	PostalCode  string `json:"postal_code" bson:"postal_code"`
}

type Terms struct {
	BillingCycle        string        `json:"billing_cycle" bson:"billing_cycle"`
	Name                float64       `json:"value" bson:"value"`
	Description         string        `json:"description" bson:"description"`
	PaymentDueDate      time.Time     `json:"payment_due_date" bson:"payment_due_date"`
	FinalPaymentDueDate time.Time     `json:"final_payment_due_date" json:"final_payment_due_date"`
	Duration            time.Duration `json:"duration" bson:"duration"`
}

type Amount struct {
	Value       float64 `json:"value" bson:"value"`
	Currency    string  `json:"currency" bson:"currency"`
	DisplayName string  `json:"display_name" bson:"display_name"`
}

type Size struct {
	Value       float64 `json:"value" bson:"value"`
	DisplayName string  `json:"display_name" bson:"display_name"`
}

type Validity struct {
	Value       int    `json:"value" bson:"value"`
	DisplayName string `json:"display_name" bson:"display_name"`
}

type Bundle struct {
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Size        *Size     `json:"size" bson:"size"`
	Price       *Amount   `json:"price" bson:"price"`
	Validity    *Validity `json:"validity" bson:"validity"`
}

type Device struct {
	Name  string `json:"name" bson:"name"`
	Color string `json:"color" bson:"color"`
	Image string `json:"image" bson:"image"`
}

type Provider struct {
	ID          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type SIMCard struct {
	MSISDN          string    `json:"msisdn" bson:"msisdn"`
	Color           string    `json:"color" bson:"color"`
	Image           string    `json:"image" bson:"image"`
	ServiceProvider *Provider `json:"provider" bson:"provider"`
}

type GenericMessage struct {
	Message string `json:"message"`
	Success bool   `json:"success" bson:"success"`
}

type Package struct {
	Meta          map[string]interface{} `json:"meta" bson:"meta"`
	ID            string                 `json:"id" bson:"id"`
	Name          string                 `json:"name" bson:"name"`
	Description   string                 `json:"description" bson:"description"`
	Device        *Device                `json:"device" bson:"device"`
	Bundle        *Bundle                `json:"bundle" bson:"bundle"`
	AddOns        []*Bundle              `json:"add_ons" bson:"add_ons"`
	Price         *Amount                `json:"price" bson:"price"`
	Providers     []*Provider            `json:"providers" bson:"providers"`
	Created       time.Time              `json:"created" bson:"created"`
	Modified      time.Time              `json:"modified" bson:"modified"`
	ActivatedDate time.Time              `json:"activated_date" bson:"activated_date"`
	ExpiryDate    time.Time              `json:"expiry_date" bson:"expiry_date"`
}

type Status struct {
	ID          int       `json:"id" bson:"id"`
	Name        string    `json:"name" json:"name"`
	Description string    `json:"description" bson:"description"`
	Color       string    `json:"color" bson:"color"`
	Created     time.Time `json:"created" bson:"created"`
}

func NewStatus(name string) *Status {

	now := time.Now()

	switch strings.ToLower(name) {
	case strings.ToLower(config.APPROVED):
		return &Status{ID: config.ONE, Name: config.APPROVED, Description: config.APPROVED, Created: now, Color: config.BLUE}
	case strings.ToLower(config.ACTIVE):
		return &Status{ID: config.TWO, Name: config.ACTIVE, Description: config.ACTIVE, Created: now, Color: config.GREEN}
	case strings.ToLower(config.SUSPENDED):
		return &Status{ID: config.THREE, Name: config.SUSPENDED, Description: config.SUSPENDED, Created: now, Color: config.GRAY}
	case strings.ToLower(config.BLOCKED):
		return &Status{ID: config.FOUR, Name: config.BLOCKED, Description: config.BLOCKED, Created: now, Color: config.RED}
	default:
		return &Status{ID: config.ZERO, Name: config.PENDING, Description: config.PENDING, Created: now, Color: config.GRAY}
	}
}
