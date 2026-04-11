def convert(number):
    return_string = ""
    
    if number % 3 == 0:
        return_string += "Pling"
    if number % 5 == 0:
        return_string += "Plang"
    if number % 7 == 0:
        return_string += "Plong"
    if return_string is not "":
        return return_string
    return str(number)
