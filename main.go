package main

import (
    "errors"
    "fmt"
    "strings"
)

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

// findSpot finds an available parking spot in the provided slice
func (p *ParkingLot) findSpot(spots []ParkingSpot) (int, error) {
    for i, spot := range spots {
        if !spot.occupied {
            return i, nil
        }
    }
    return -1, errors.New("no available spots")
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

// displayParkingLotStatus shows the available spots in the parking lot
func (p *ParkingLot) displayParkingLotStatus() {
    smallAvailable := p.countAvailableSpots(p.smallspot)
    mediumAvailable := p.countAvailableSpots(p.mediumspot)
    largeAvailable := p.countAvailableSpots(p.largespot)

    fmt.Println("Parking lot status:")
    fmt.Printf("Small spots available: %d\n", smallAvailable)
    fmt.Printf("Medium spots available: %d\n", mediumAvailable)
    fmt.Printf("Large spots available: %d\n", largeAvailable)
}

// countAvailableSpots counts the available parking spots
func (p *ParkingLot) countAvailableSpots(spots []ParkingSpot) int {
    count := 0
    for _, spot := range spots {
        if !spot.occupied {
            count++
        }
    }
    return count
}

func main() {
    var small, medium, large int
    fmt.Println("Welcome to the Parking Lot System!")
    fmt.Println("Enter the number of small, medium, and large parking spots:")

    fmt.Print("Small spots: ")
    fmt.Scanln(&small)

    fmt.Print("Medium spots: ")
    fmt.Scanln(&medium)

    fmt.Print("Large spots: ")
    fmt.Scanln(&large)

    // Create a new parking lot with user-specified number of spots
    parkingLot := newParkingLot(small, medium, large)

    for {
        fmt.Println("\nChoose an action: park | retrieve | status | exit")
        var command string
        fmt.Scanln(&command)

        switch strings.ToLower(command) {
        case "park":
            fmt.Println("Enter vehicle size (small, medium, large):")
            var vehicleSize string
            fmt.Scanln(&vehicleSize)

            if vehicleSize != "small" && vehicleSize != "medium" && vehicleSize != "large" {
                fmt.Println("Invalid vehicle size. Please enter small, medium, or large.")
                continue
            }

            vehicle := Vehicle{size: vehicleSize}
            ticket, spotType, err := parkingLot.parkVehicle(vehicle)
            if err != nil {
                fmt.Println("Error parking vehicle:", err)
            } else {
                fmt.Printf("Vehicle parked with ticket: %d in a %s spot.\n", ticket, spotType)
            }

        case "retrieve":
            fmt.Println("Enter ticket number:")
            var ticket int
            fmt.Scanln(&ticket)

            vehicle, spotType, err := parkingLot.retrieveVehicle(ticket)
            if err != nil {
                fmt.Println("Error retrieving vehicle:", err)
            } else {
                fmt.Printf("Vehicle retrieved from a %s spot with ticket: %d (Vehicle size: %s).\n", spotType, ticket, vehicle.size)
            }

        case "status":
            parkingLot.displayParkingLotStatus()

        case "exit":
            var confirm string
            fmt.Println("Are you sure you want to exit? (yes/no)")
            fmt.Scanln(&confirm)
            if strings.ToLower(confirm) == "yes" {
                fmt.Println("Exiting the system. Goodbye!")
                return
            }

        default:
            fmt.Println("Invalid command. Please choose park, retrieve, status, or exit.")
        }
    }
}
