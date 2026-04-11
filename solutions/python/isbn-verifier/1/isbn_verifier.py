def is_valid(isbn):
    isbn_total = 0
    isbn = isbn.replace("-", "")

    if len(isbn) != 10:
        return False
    
    if not isbn[:9].isdigit():
        return False
    
    if not (isbn[9].isdigit() or isbn[9] == "X"):
        return False
    
    for index, char in enumerate(isbn):
        if char == "X":
            value = 10
        else:
            value = int(char)
        weight = 10 - index
        isbn_total += value * weight
    
    return isbn_total % 11 == 0