def transform(legacy_data):
    mapped_dict = {}

    for score, letters in legacy_data.items():
        for letter in letters:
            mapped_dict[letter.lower()] = score

    return mapped_dict
