package elon

import (
	"strconv"
)

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if car.battery < car.batteryDrain {
		return
	}
	car.distance += car.speed
	car.battery -= car.batteryDrain
}

// TODO: define the 'DisplayDistance() string' method
func (car *Car) DisplayDistance() string {
	display := "Driven " + strconv.Itoa(car.distance) + " meters"
	return display
}

// TODO: define the 'DisplayBattery() string' method
func (car *Car) DisplayBattery() string {
	display := "Battery at " + strconv.Itoa(car.battery) + "%"
	return display
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car *Car) CanFinish(trackDistance int) bool {
	batteryBefore := 101

	for batteryBefore > car.battery {
		batteryBefore = car.battery
		car.Drive()
	}
	return car.distance >= trackDistance
}
