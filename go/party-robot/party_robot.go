package partyrobot

import (
	"fmt"
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name) 
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %s years old!", name, strconv.Itoa(age)) 
}

func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	directionMessage := fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table, direction, distance )
	neighborMessage := fmt.Sprintf("You will be sitting next to %s.", neighbor)

	return fmt.Sprintf("%s\n%s\n%s", Welcome(name),directionMessage,neighborMessage)
}
