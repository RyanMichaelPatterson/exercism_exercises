def find(search_list, value):
    left_index = 0
    right_index = len(search_list) -1

    while left_index <= right_index:
        middle_index = (left_index + right_index) // 2
        if search_list[middle_index] == value:
            return middle_index
        if search_list[middle_index] < value:
            left_index = middle_index + 1
        else:
            right_index = middle_index - 1

    else:
          raise ValueError("value not in array")
    
