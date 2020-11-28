package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type User struct {
	ID             primitive.ObjectID       `bson:"_id,omitempty"`
	UserName       string    `json:"username"`
	FirstName      string    `json:"firstname"`
	LastName       string    `json:"lastname"`
	Email          string    `json:"email"`
	Password []byte    `json:"hashedPassword"`
	Created        time.Time `json:"created"`
	Active         bool      `json:"active"`
	Address        Address   `json:"address"`
}

//Address data type for the address object
type Address struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Street string             `json:"street"`
	City   string             `json:"city"`
	State  string             `json:"state"`
	Zip    string             `json:"zip"`
}

//Store data type for the store object
type Store struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `json:"name"`
	PhoneNumber string             `json:"phoneNumber"`
	Owner       string             `json:"owner"`
	Address     Address            `json:"address"`
	Items       []Item             `json:"items"`
	// Services Rendered because services have different values and aren't products that can just be
	// picked up
}

type Item struct {
	ItemID          int    `json:"itemID"`
	ItemName        string `json:"itemName"`
	ItemPrice       string `json:"itemPrice"`
	ItemDescription string `json:"itemDescription"`
	InStock         bool   `json:"inStock"`
	Visibility      bool   `json:"visibility`
	// Type defines whether it is a product or service
	Type string `json:"type"`
}

// type Service struct {
// 	ServiceId int `json:"serviceId"`
// 	ServiceName string `json:"serviceName"`
// 	ServicePrice float64 `json:"servicePrice"`
// 	ServiceDescription string `json:"serviceDescription"`
// }

// type Product struct{
// 	ProductId int `json:"productId"`
// 	ProductName string `json:"productName"`
// 	ProductPrice float64 `json:"productPrice"`
// 	ProductDescription string `json:"productDescription"`
// }

type Driver struct {
	User      User      `json:"user"`
	Address   Address   `json:"address"`
	HashedSSN string    `json:"hashedSSN"`
	Insurance Insurance `json:"insurance"`
	License   License   `json:"license"`
	Vehicle   Vehicle   `json:"vehicle"`
}

//Image data type for the image object
type Image struct {
	photo map[int]int
}

//Vehicle data type for the vehicle object
type Vehicle struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	VehicleMake  string             `json:"vehicleMake"`
	VehicleModel string             `json:"vehicleModel"`
	VehicleYear  string             `json:"vehicleYear"`
	VehicleColor string             `json:"vehicleColor"`
	VinNumber    string             `json:"vinNumber"`
}

//Insurance data type for the insurance object
type Insurance struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Insured           *User              `json:"insured"`
	InsuranceProvider string             `json:"insuranceProvider"`
	PolicyNumber      string             `json:"policyNumber"`
	ExpirationDate    time.Time          `json:"expirationDate"`
	Vehicle           Vehicle            `json:"vehicle"`
}

//License data type for the license object
type License struct {
	Proof Image `json:"proof"`
}

//Students data type for the students object
type Students struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	StudentID int32              `bson:"student_id,omitempty"`
	Type      string             `bson:"type,omitempty"`
	Score     float64            `bson:"score,omitempty"`
}
