def square_root(number):
    answer_not_found = True
    squared_number = 1

    while answer_not_found:
        if (squared_number * squared_number) == number:
            return squared_number
        squared_number += 1