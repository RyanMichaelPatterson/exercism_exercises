package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averageTime int) int{
    if averageTime == 0 {
        return len(layers) * 2
    } else {
        return len(layers) * averageTime
    }
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64){
    
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles"{
            noodles += 50
        }
        if layers[i] == "sauce"{
            sauce += .2
        }
    }
    return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, personalList []string) {
    
    personalList[len(personalList)-1] = friendsList [len(friendsList)-1]
}


// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(basePortion []float64, totalPortions int) []float64{
	var scaledRecipe []float64
    
    for i := 0; i < len(basePortion); i++ {
        scaledRecipe = append(scaledRecipe, (basePortion[i] * (float64(totalPortions) / 2)))
    }
    return scaledRecipe
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
