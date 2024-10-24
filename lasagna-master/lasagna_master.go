package lasagna

func PreparationTime(layers []string, averagePreparationTime int) int {
	preparationTime := 2

	if averagePreparationTime > 0 {
		preparationTime = averagePreparationTime
	}

	return preparationTime * len(layers)
}

func Quantities(layers []string) (int, float64) {
	noodleLayers := 0
	sauceLayers := 0

	for _, layer := range layers {
		if layer == "noodles" {
			noodleLayers++
		}

		if layer == "sauce" {
			sauceLayers++
		}
	}

	return noodleLayers * 50, float64(sauceLayers) * 0.2
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	lastFriendsListIngredient := friendsList[len(friendsList)-1]

	myList[len(myList)-1] = lastFriendsListIngredient
	return myList
}

func ScaleRecipe(quantities []float64, recipes int) (finalRecipe []float64) {
	finalRecipe = make([]float64, len(quantities))

	for i, quantity := range quantities {
		finalRecipe[i] = quantity / 2 * float64(recipes)
	}

	return finalRecipe
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
