def equilateral(sides):
    if 0 not in {sides[0], sides[1], sides[2]}:
        if sides[0] == sides[1] and sides[0] == sides[2]:
            return True
    return False


def isosceles(sides):
    if 0 not in {sides[0], sides[1], sides[2]}:
        if sides[0] + sides[1] >= sides[2] and sides[0] + sides[2] >= sides[1] and sides[1] + sides[2] >= sides[0]:
            if sides[0] == sides[1]:
                return True
            if sides[0] == sides[2]:
                return True
            if sides [1] == sides[2]:
                return True
    return False


def scalene(sides):
    if 0 not in {sides[0], sides[1], sides[2]}:
        if sides[0] + sides[1] >= sides[2] and sides[0] + sides[2] >= sides[1] and sides[1] + sides[2] >= sides[0]:
            if sides[0] not in (sides[1], sides[2]) and sides[1] not in (sides[0], sides[2]):
                return True
    return False
