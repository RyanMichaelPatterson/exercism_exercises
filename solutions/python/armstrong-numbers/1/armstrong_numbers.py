def is_armstrong_number(number):

    total_sum = 0

    for digit in str(number):
        raised_digit = int(digit) ** len(str(number))
        total_sum = total_sum + raised_digit

    if number == total_sum:
        return True
    return False