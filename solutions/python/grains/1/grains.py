def square(number):
    grains_on_square = 1

    if number > 0 and number < 65:
        for i in range(1, number):
            grains_on_square = (grains_on_square * 2)
    
        return grains_on_square
    raise ValueError("square must be between 1 and 64")


def total():
    total_grains = 0

    for i in range(1, 65):
        total_grains = total_grains + square(i)

    return total_grains
        
