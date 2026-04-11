package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var totalBirdCount int = 0
	for i := 0; i < len(birdsPerDay); i++{
        totalBirdCount = totalBirdCount + birdsPerDay[i]
    }
    return totalBirdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var totalBirdCount int = 0
    var weekIndex int = (week - 1) * 7

	for i := weekIndex; i < (weekIndex +7); i ++ {
        totalBirdCount = totalBirdCount + birdsPerDay[i]
    }
    return totalBirdCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
        birdsPerDay[i] = birdsPerDay[i] + 1
    }
    return birdsPerDay
}
