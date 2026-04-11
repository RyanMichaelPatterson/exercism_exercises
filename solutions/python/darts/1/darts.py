def score(x, y):
    distance_sq = (x * x) + (y * y)

    if distance_sq <= 1:
        return 10
    elif distance_sq <= 25:
        return 5
    elif distance_sq <= 100:
        return 1
    else:
        return 0