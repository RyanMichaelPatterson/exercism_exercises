package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var workingCarsPerHour float64 = float64(productionRate) * (float64(successRate) / 100)
    return workingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var workingCarsPerMinute int = int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
    return workingCarsPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var bundleCar int = carsCount / 10
    var individualCar int = carsCount % 10
    return uint((bundleCar * 95000) + (individualCar * 10000))
}
