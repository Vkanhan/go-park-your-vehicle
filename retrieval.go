package main

import (
	"errors"
)

// retrieveVehicle retrieves a vehicle using its ticket and frees the spot
func (p *ParkingLot) retrieveVehicle(ticket int) (Vehicle, string, error) {
	v, ok := p.tickets[ticket]
	if !ok {
		return Vehicle{}, "", errors.New("invalid ticket")
	}

	spotIndex, spotType, err := p.getSpotIndex(v)
	if err != nil {
		return Vehicle{}, "", err
	}

	p.freeSpot(spotType, spotIndex)
	delete(p.tickets, ticket)
	return v, spotType, nil
}

// getSpotIndex finds the index of the parking spot for the given vehicle
func (p *ParkingLot) getSpotIndex(v Vehicle) (int, string, error) {
	var spotIndex int
	var err error
	var spotType string

	switch v.size {
	case "small":
		spotIndex, err = p.findSpotIndex(p.smallspot, v)
		spotType = "small"
		if err != nil {
			spotIndex, err = p.findSpotIndex(p.mediumspot, v)
			spotType = "medium"
			if err != nil {
				spotIndex, err = p.findSpotIndex(p.largespot, v)
				spotType = "large"
			}
		}
	case "medium":
		spotIndex, err = p.findSpotIndex(p.mediumspot, v)
		spotType = "medium"
		if err != nil {
			spotIndex, err = p.findSpotIndex(p.largespot, v)
			spotType = "large"
		}
	case "large":
		spotIndex, err = p.findSpotIndex(p.largespot, v)
		spotType = "large"
	default:
		return -1, "", errors.New("invalid vehicle size")
	}

	if err != nil {
		return -1, "", errors.New("no spot found for vehicle")
	}

	return spotIndex, spotType, nil
}

// freeSpot marks the spot as available
func (p *ParkingLot) freeSpot(spotType string, index int) {
	switch spotType {
	case "small":
		p.smallspot[index].occupied = false
	case "medium":
		p.mediumspot[index].occupied = false
	case "large":
		p.largespot[index].occupied = false
	}
}

// findSpotIndex finds the index of the parking spot for the given vehicle
func (p *ParkingLot) findSpotIndex(spots []ParkingSpot, v Vehicle) (int, error) {
	for i, spot := range spots {
		if spot.size == v.size && spot.occupied {
			return i, nil
		}
	}
	return -1, errors.New("spot not found")
}
