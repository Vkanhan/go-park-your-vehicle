package main

import "errors"

// newParkingLot creates a new ParkingLot with specified numbers of spots
func newParkingLot(small, medium, large int) *ParkingLot {
	p := &ParkingLot{
		smallspot:  make([]ParkingSpot, small),
		mediumspot: make([]ParkingSpot, medium),
		largespot:  make([]ParkingSpot, large),
		tickets:    make(map[int]Vehicle),
	}

	// Initialize parking spots
	for i := range p.smallspot {
		p.smallspot[i] = ParkingSpot{size: "small", occupied: false}
	}
	for i := range p.mediumspot {
		p.mediumspot[i] = ParkingSpot{size: "medium", occupied: false}
	}
	for i := range p.largespot {
		p.largespot[i] = ParkingSpot{size: "large", occupied: false}
	}

	return p
}

// parkVehicle parks a vehicle and returns a ticket ID or an error
func (p *ParkingLot) parkVehicle(v Vehicle) (int, string, error) {
	spotIndex, spotType, err := p.findAvailableSpot(v.size)
	if err != nil {
		return -1, "", err
	}

	p.occupySpot(spotType, spotIndex)
	p.ticketID++
	p.tickets[p.ticketID] = v

	return p.ticketID, spotType, nil
}

// findAvailableSpot finds an available parking spot based on vehicle size
func (p *ParkingLot) findAvailableSpot(vehicleSize string) (int, string, error) {
	var spotIndex int
	var err error
	var spotType string

	switch vehicleSize {
	case "small":
		spotIndex, err = p.findSpot(p.smallspot)
		spotType = "small"
		if err != nil {
			spotIndex, err = p.findSpot(p.mediumspot)
			spotType = "medium"
			if err != nil {
				spotIndex, err = p.findSpot(p.largespot)
				spotType = "large"
			}
		}
	case "medium":
		spotIndex, err = p.findSpot(p.mediumspot)
		spotType = "medium"
		if err != nil {
			spotIndex, err = p.findSpot(p.largespot)
			spotType = "large"
		}
	case "large":
		spotIndex, err = p.findSpot(p.largespot)
		spotType = "large"
	default:
		return -1, "", errors.New("invalid vehicle size")
	}

	if err != nil {
		return -1, "", errors.New("no parking spot available for the vehicle")
	}

	return spotIndex, spotType, nil
}

// occupySpot marks the spot as occupied
func (p *ParkingLot) occupySpot(spotType string, index int) {
	switch spotType {
	case "small":
		p.smallspot[index].occupied = true
	case "medium":
		p.mediumspot[index].occupied = true
	case "large":
		p.largespot[index].occupied = true
	}
}

// findSpot finds an available parking spot in the provided slice
func (p *ParkingLot) findSpot(spots []ParkingSpot) (int, error) {
	for i, spot := range spots {
		if !spot.occupied {
			return i, nil
		}
	}
	return -1, errors.New("no available spots")
}