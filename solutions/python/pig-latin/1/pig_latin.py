def translate(text):
    answer = ''

    for word in text.split():
        print(parse_word(word))
        if answer == '':
            answer = parse_word(word)
        else:
            answer = answer + ' ' + parse_word(word)
    return answer


def parse_word(word):  
    if word[:1] in ('a', 'e', 'i', 'o', 'u') or word[:2] in ('xr', 'yt'):
        return word + 'ay'
        
    if word[:1] not in ('a', 'e', 'i', 'o', 'u'):
        if word[0] == 'y':
            word = word[1:] + word[0]
        for char in word:
            if word[:2] == 'qu':
                return word[2:] + 'quay'
            if char not in ('a', 'e', 'i', 'o', 'u'):
                if char == 'y': 
                    return word + 'ay'
                word = word[1:] + char
            else:
                return word + 'ay'
            
                

