package main

// ParkingSpot represents a parking spot in the parking lot
type ParkingSpot struct {
	size     string
	occupied bool
}

// Vehicle represents a vehicle with a certain size
type Vehicle struct {
	size string
}

// ParkingLot represents the parking lot with various spot sizes and ticket management
type ParkingLot struct {
	smallspot  []ParkingSpot
	mediumspot []ParkingSpot
	largespot  []ParkingSpot
	tickets    map[int]Vehicle
	ticketID   int
}
