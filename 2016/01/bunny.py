f = open("input1.txt")
paths = f.readlines()
p = paths[0].strip()
steps = p.split(', ')

attitude = [0,1]
position = [0,0]

x = 0
y = 1

position_list = []

for step in steps:
    direction = step[:1]
    if direction == "R":
        if attitude == [0,1]:
            attitude = [1,0]
            for i in range(int(step[1:])):
                position[x] += 1
                if position in position_list:
                    print position
                position_list.append(list(position))

        elif attitude == [1,0]:
            attitude = [0,-1]
            for i in range(int(step[1:])):
                position[y] -= 1
                if position in position_list:
                    print position
                position_list.append(list(position))

        elif attitude == [0,-1]:
            attitude = [-1,0]
            for i in range(int(step[1:])):
                position[x] -= 1
                if position in position_list:
                    print position
                position_list.append(list(position))

        else: # attitude == [-1,0]
            attitude = [0,1]
            for i in range(int(step[1:])):
                position[y] += 1
                if position in position_list:
                    print position
                position_list.append(list(position))

    else: # step == "L":
        if attitude == [0,1]:
            attitude = [-1,0]
            for i in range(int(step[1:])):
                position[x] -= 1
                if position in position_list:
                    print position
                position_list.append(list(position))

        elif attitude == [-1,0]:
            attitude = [0,-1]
            for i in range(int(step[1:])):
                position[y] -= 1
                if position in position_list:
                    print position
                    break
                position_list.append(list(position))

        elif attitude == [0,-1]:
            attitude = [1,0]
            for i in range(int(step[1:])):
                position[x] += 1
                if position in position_list:
                    print position
                position_list.append(list(position))

        else: # attitude == [1,0]
            attitude = [0,1]
            for i in range(int(step[1:])):
                position[y] += 1
                if position in position_list:
                    print position
                position_list.append(list(position))


print abs(position[0]) + abs(position[1])
