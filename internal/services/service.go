package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/CarApp/internal/model"
)

type CarService interface {
	AddCarDetails() error
	ListAll() []model.Car
	ListCar() *model.Car
	UpdateCarDetails(model.Car) error
}

type CarRepo struct {
	Dataa     []byte
	StoreData map[string][]model.Car
}

func NewService() *CarRepo {
	return &CarRepo{}
}

func (c *CarRepo) AddCarDetails() error {
	var err error
	if len(c.StoreData) == 0 {
		c.StoreData = map[string][]model.Car{}
	}

	var store []model.Car
	_ = json.Unmarshal(c.Dataa, &store)

	for _, value := range store {
		c.StoreData[value.Make] = append(c.StoreData[value.Make], value)
	}

	return err
}

func (c *CarRepo) ListAll() []model.Car {
	var output []model.Car

	for _, value := range c.StoreData {
		output = append(output, value...)
	}

	return output
}

func (c *CarRepo) ListCar(key string) *model.Car {
	for _, value := range c.StoreData {
		for _, check := range value {
			if check.Id == key {
				return &check
			}
		}
	}

	return nil
}

func (c *CarRepo) UpdateCarDetails(data model.Car) error {
	for key, value := range c.StoreData {
		if key == data.Make {

			for index, check := range value {
				fmt.Println("[DEBUG]:", check)
				if check.Id == data.Id {
					check.Category = data.Category
					check.Mileage = data.Mileage
					check.Price = data.Price
					check.Year = data.Year
					c.StoreData[check.Make][index] = check
					return nil
				}
			}
		}
	}

	return errors.New("unable to find record")
}
