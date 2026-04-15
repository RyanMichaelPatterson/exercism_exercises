package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	colors := []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
    return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	if color == "black"{
        return 0
    }
    if color == "brown"{
        return 1
    }
    if color == "red"{
        return 2
    }
    if color == "orange"{
        return 3
    }
    if color == "yellow"{
        return 4
    }
    if color == "green"{
        return 5
    }
    if color == "blue"{
        return 6
    }
    if color == "violet"{
        return 7
    }
    if color == "grey"{
        return 8
    } else {
        return 9
    }
}
