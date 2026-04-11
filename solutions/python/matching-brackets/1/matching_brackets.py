def is_paired(input_string):
    opening_bracket = {'(', '[', '{'}
    matching_bracket = {')': '(', ']': '[', '}': '{'}
    char_stack = []

    for char in input_string:
        if char in opening_bracket:
            char_stack.append(char)
        elif char in matching_bracket:
            if not char_stack or char_stack.pop() != matching_bracket[char]:
                return False
    return not char_stack