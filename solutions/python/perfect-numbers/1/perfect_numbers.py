def classify(number):
    """ A perfect number equals the sum of its positive divisors.

    :param number: int a positive integer
    :return: str the classification of the input integer
    """
    if (number <= 0) or number == "":
        raise ValueError("Classification is only possible for positive integers.")

    if number == 1:
        return "deficient"
    else:
        aliquot_sum = 1
        digit = 2
        while digit * digit <= number:
            if number % digit == 0:
                quotiant = number // digit
                aliquot_sum += digit
                if quotiant != digit: 
                    aliquot_sum += quotiant
            digit += 1
    
        if aliquot_sum == number:
            return "perfect"
        elif aliquot_sum > number:
            return "abundant"
        else:
            return "deficient"
        
        
    

