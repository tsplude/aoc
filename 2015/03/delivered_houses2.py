f = open("input2.txt")
paths = [l.strip() for l in f.readlines()]

santa_x = 0
santa_y = 0

robot_x = 0
robot_y = 0

houses = set()
houses.add(str(santa_x) + str(santa_y))

count = 0
for p in paths:
    for step in p:
        if (step == "^"):
            if count % 2:
                santa_y += 1
            else:
                robot_y += 1
        elif (step == "v"):
            if count % 2:
                santa_y -= 1
            else:
                robot_y -= 1
        elif (step == ">"):
            if count % 2:
                santa_x += 1
            else:
                robot_x += 1
        else:
            if count % 2:
                santa_x -= 1
            else:
                robot_x -=1
        houses.add(str(santa_x) + str(santa_y))
        houses.add(str(robot_x) + str(robot_y))

    print len(p)
    print len(houses)

