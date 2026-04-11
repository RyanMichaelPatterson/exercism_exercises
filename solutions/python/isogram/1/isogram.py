def is_isogram(string):
    if string != '':
        string_list = list(string)

        for index, char in enumerate(string_list):
            string_list[index] = char.lower()

        for char in string_list:
            occurences = string_list.count(char)
            if occurences > 1 and char not in (" ", "-"):
                return False
    return True
