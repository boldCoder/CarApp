package services

import (
	"encoding/json"
	"errors"

	"github.com/CarApp/internal/model"
)

type CarService interface {
	AddCarDetails() error
	ListAll() []model.Car
	ListCar() *model.Car
	UpdateCarDetails(model.Car) error
}

type CarRepo struct {
	Details     []byte
	DataStore map[string][]model.Car
}

func NewService() *CarRepo {
	return &CarRepo{}
}

// Add new car details.
func (c *CarRepo) AddCarDetails() error {
	// Check if it is an empty map. If yes, initialize the map.
	if len(c.DataStore) == 0 {
		c.DataStore = map[string][]model.Car{}
	}

	var store []model.Car
	err := json.Unmarshal(c.Details, &store)
	if err != nil {
		return err
	}

	for _, value := range store {
		c.DataStore[value.Make] = append(c.DataStore[value.Make], value)
	}

	return err
}

// List all cars stored in the DB
func (c *CarRepo) ListAll() []model.Car {
	var output []model.Car

	for _, value := range c.DataStore {
		output = append(output, value...)
	}

	return output
}

// List car details with the unique ID associated with it.
func (c *CarRepo) ListCar(key string) *model.Car {
	for _, value := range c.DataStore {
		for _, check := range value {
			if check.Id == key {
				return &check
			}
		}
	}

	return nil
}

// Update car details for a specific car
func (c *CarRepo) UpdateCarDetails(data model.Car) error {
	for key, value := range c.DataStore {
		if key == data.Make {
			for index, check := range value {
				if check.Id == data.Id {
					check.Category = data.Category
					check.Mileage = data.Mileage
					check.Price = data.Price
					check.Year = data.Year
					c.DataStore[check.Make][index] = check
					return nil
				}
			}
		}
	}

	return errors.New("unable to find record")
}
