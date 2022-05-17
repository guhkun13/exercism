package lasagna

// import "fmt"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePrepTime int) int {
	if averagePrepTime == 0 {
		averagePrepTime = 2
	}
	return len(layers) * averagePrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	const noodlesEachLayer int = 50 //gram
	const sauceEachLayer float64 = 0.2  //liters

	var countNoodles int = 0
	var countSauce float64 = 0

	for _, layer := range layers {
		if layer == "noodles" {
			countNoodles++
		} else if layer == "sauce" {
			countSauce++
		}
	}
	
	noodles = countNoodles * noodlesEachLayer
	sauce = countSauce * sauceEachLayer
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	// fmt.Println(amounts, portions)
	var scaledRecipe []float64

	for _, amount := range amounts {
		scaledRecipe = append(scaledRecipe, amount * (float64(portions)/2)) 
	}
	return scaledRecipe
}
