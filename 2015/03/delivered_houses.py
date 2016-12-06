f = open("input.txt")
paths = [l.strip() for l in f.readlines()]

for p in paths:
    x = 0
    y = 0
    house_coordinates = set()
    house_coordinates.add(str(x) + str(y))
    for step in p:
        if (step == "^"):
            y += 1
        elif (step == "v"):
            y -= 1
        elif (step == ">"):
            x += 1
        elif (step == "<"):
            x -= 1
        else:
            print "oops"
        house_coordinates.add(str(x) + str(y))

    print len(p)
    print len(house_coordinates)

