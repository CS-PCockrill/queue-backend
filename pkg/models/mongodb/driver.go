package mongodb

import (
	"github.com/CS-PCockrill/queue/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type DriverFunctions struct{
	CLIENT *mongo.Client
}

func (d *DriverFunctions) RegisterDriver(driver *models.Driver) (int, error) {
	// Register a driver, verify background check, and verify state id & insurance
	// newDriver := d.CLIENT.Database("queue")
	// driverCollection := newDriver.Collection("drivers")
	// var driver models.Driver

	// Get the current user in session
	// driver.User = current session user


	return 0, nil
}

func (d *DriverFunctions) Validate(email, password string) (int, error) {
	// Validate the drivers login credentials
	return 0, nil
}

func (d *DriverFunctions) getDriver(id int) *models.Driver {
	// Get a driver with parameter id
	return nil
}
