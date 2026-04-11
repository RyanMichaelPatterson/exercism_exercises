def steps(number):
    if number > 0:
        itteration = 0
        while True:
            if number == 1:
                return itteration 
            if number % 2 == 0:
                number = int(number / 2)
            else:
                number = int((number * 3) + 1)
            itteration += 1
        
    raise ValueError("Only positive integers are allowed")