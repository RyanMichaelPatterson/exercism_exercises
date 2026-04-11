def square_of_sum(number):
    square_sum = 0

    for value in range(number + 1):
        square_sum = square_sum + value

    return square_sum * square_sum

def sum_of_squares(number):
    sum_squares = 0
    
    for value in range(number + 1):
        sum_squares = sum_squares + (value * value)

    return sum_squares
        


def difference_of_squares(number):
    return square_of_sum(number) - sum_of_squares(number)
