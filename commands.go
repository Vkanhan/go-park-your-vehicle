package main

import (
	"fmt"
	"strings"
)

func commandPark(parkingLot *ParkingLot) {
	fmt.Println("Enter vehicle size (small, medium, large):")
	var vehicleSize string
	fmt.Scanln(&vehicleSize)

	// Validate vehicle size inside the function
	if vehicleSize != "small" && vehicleSize != "medium" && vehicleSize != "large" {
		fmt.Println("Invalid vehicle size. Please enter small, medium, or large.")
		return
	}

	vehicle := Vehicle{size: vehicleSize}
	ticket, spotType, err := parkingLot.parkVehicle(vehicle)
	if err != nil {
		fmt.Println("Error parking vehicle:", err)
	} else {
		fmt.Printf("Vehicle parked with ticket: %d in a %s spot.\n", ticket, spotType)
	}
}

func commandRetrieve(parkingLot *ParkingLot) {
	fmt.Println("Enter ticket number:")
	var ticket int
	fmt.Scanln(&ticket)

	vehicle, spotType, err := parkingLot.retrieveVehicle(ticket)
	if err != nil {
		fmt.Println("Error retrieving vehicle:", err)
	} else {
		fmt.Printf("Vehicle retrieved from a %s spot with ticket: %d (Vehicle size: %s).\n", spotType, ticket, vehicle.size)
	}
}

// displayParkingLotStatus shows the available spots in the parking lot
func (p *ParkingLot) commandStatus() {
	// Define a helper function within commandStatus to count available spots for each size
	countAvailableSpots := func(spots []ParkingSpot) int {
		count := 0
		for _, spot := range spots {
			if !spot.occupied {
				count++
			}
		}
		return count
	}

	// Use the helper function to get the available spots for each type
	smallAvailable := countAvailableSpots(p.smallspot)
	mediumAvailable := countAvailableSpots(p.mediumspot)
	largeAvailable := countAvailableSpots(p.largespot)

	fmt.Println("Parking lot status:")
	fmt.Printf("Small spots available: %d\n", smallAvailable)
	fmt.Printf("Medium spots available: %d\n", mediumAvailable)
	fmt.Printf("Large spots available: %d\n", largeAvailable)
}

func commandExit() bool {
	var confirm string
	fmt.Println("Are you sure you want to exit? (yes/no)")
	fmt.Scanln(&confirm)
	return strings.ToLower(confirm) == "yes"
}
