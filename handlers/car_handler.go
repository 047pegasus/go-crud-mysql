package handlers

import (
	"github.com/047pegasus/go-crud-mysql/controllers"
	"github.com/047pegasus/go-crud-mysql/models"

	"gofr.dev/pkg/gofr"
)

type CarHandler struct {
	carController *controllers.CarController
}

func NewCarHandler(cc *controllers.CarController) *CarHandler {
	return &CarHandler{carController: cc}
}

// gofr handler to add new car to garage
func (ch *CarHandler) AddCarToGarage(ctx *gofr.Context) (interface{}, error) {
	var car models.Car
	// Parse request body to car struct

	if err := ch.carController.AddCarToGarage(car); err != nil {
		return nil, err
	}

	// Write JSON response
	return nil, nil
}

func (ch *CarHandler) GetCarsInGarage(ctx *gofr.Context) (interface{}, error) {
	// Parse URL parameters
	cars, err := ch.carController.GetCarsInGarage()
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (ch *CarHandler) GetCarInGarage(ctx *gofr.Context) (interface{}, error) {
	// Parse URL parameters
	car, err := ch.carController.GetCarInGarage(0)
	if err != nil {
		return nil, err
	}

	// Write JSON response
	return car, nil
}

func (ch *CarHandler) UpdateCarInGarage(ctx *gofr.Context) (interface{}, error) {
	var car models.Car
	// Parse request body to car struct

	if err := ch.carController.UpdateCarInGarage(car.ID, car); err != nil {
		return nil, err
	}

	// Write JSON response
	return nil, nil
}

func (ch *CarHandler) RemoveCarFromGarage(ctx *gofr.Context) (interface{}, error) {
	// Parse URL parameters
	id := 0

	if err := ch.carController.RemoveCarFromGarage(id); err != nil {
		return nil, err
	}

	// Write JSON response
	return nil, nil
}
