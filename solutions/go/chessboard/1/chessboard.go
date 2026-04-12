package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    var count int
	for _, value := range cb[file]{
        if value == true {
             count ++
        }
    }
    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    var count int
	if rank > 0 && rank < 9 {
        for _, value := range cb{
            if value[rank -1] == true {
                count ++
            }
        }
        return count
    }
    return 0
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    return len(cb) * len(cb)
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    var count int
    for _, values := range cb{
        for _, value := range values {
            if value == true {
                count ++
            }
        }
    }
    return count
}
