package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"

}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var finalChoose string = option1
	if option1 > option2 {
		finalChoose = option2
	} 

	return fmt.Sprintf("%s is clearly the better choice.", finalChoose)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var cutPrice float64

	if age < 3  {
		cutPrice = .8
	} else if age < 10 {
		cutPrice = .7
	} else {
		cutPrice = .5
	}

	return originalPrice * cutPrice

}
