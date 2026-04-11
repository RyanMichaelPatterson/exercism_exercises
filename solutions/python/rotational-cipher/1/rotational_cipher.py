def rotate(text, key):
    result = []
    for char in text:
        if 'a' <= char <= 'z':
            shifted = (ord(char) - ord('a') + key) % 26 + ord('a')
            result.append(chr(shifted))
        elif 'A' <= char <= 'Z':
            shifted = (ord(char) - ord('A') + key) % 26 + ord('A')
            result.append(chr(shifted))
        else:
            result.append(char)
    return ''.join(result)