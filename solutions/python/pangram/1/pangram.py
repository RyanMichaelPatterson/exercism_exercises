def is_pangram(sentence):
    if not sentence == '':
        char_list = [s.lower() for s in list(sentence)]
        ascii_letters = 'abcdefghijklmnopqrstuvwxyz'
        for letter in ascii_letters:
            if  letter.isalpha():
                if letter not in char_list:
                    print(letter)
                    return False      
        return True
                
    return False
