package darts

func Score(x, y float64) int {
	var distanceOfSquares float64 = (x * x) + (y * y)

    if distanceOfSquares <= 1 {
        return 10
    } else if distanceOfSquares <= 25 {
        return 5
    } else if distanceOfSquares <= 100 {
        return 1
    } else {
        return 0
    }
}
