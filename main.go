package main

import (
    "fmt"
    "strings"
)

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
            commandPark(parkingLot)
        case "retrieve":
            commandRetrieve(parkingLot)
        case "status":
            parkingLot.commandStatus()
        case "exit":
            if commandExit() {
                fmt.Println("Exiting the system. Goodbye!")
                return
            }
        default:
            fmt.Println("Invalid command. Please choose park, retrieve, status, or exit.")
        }
    }
}

