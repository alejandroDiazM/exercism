package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer != 0 {
		return len(layers) * timePerLayer
	} else {
		return len(layers) * 2
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {

	NoodlesLayerName := "noodles"
	SauceLayerName := "sauce"

	for _, layer := range layers {
		switch layer {
		case NoodlesLayerName:
			noodles += 50
		case SauceLayerName:
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) (scaledAmounts []float64) {

	scaledAmounts = make([]float64, len(amounts))
	scalingFactor := float64(portions) / 2.0
	for i, amount := range amounts {
		scaledAmounts[i] = amount * scalingFactor
	}
	return scaledAmounts
}
