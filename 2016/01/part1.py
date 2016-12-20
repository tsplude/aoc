# This is 01.1

f = open("input1.txt")
paths = f.readlines()
p = paths[0].strip()
steps = p.split(', ')

attitude = [0,1]
position = [0,0]

x = 0
y = 1

for step in steps:
    direction = step[:1]
    if direction == "R":
        if attitude == [0,1]:
            attitude = [1,0]
            position[x] += int(step[1:])
        elif attitude == [1,0]:
            attitude = [0,-1]
            position[y] -= int(step[1:])
        elif attitude == [0,-1]:
            attitude = [-1,0]
            position[x] -= int(step[1:])
        else: # attitude == [-1,0]
            attitude = [0,1]
            position[y] += int(step[1:])
    else: # step == "L":
        if attitude == [0,1]:
            attitude = [-1,0]
            position[x] -= int(step[1:])
        elif attitude == [-1,0]:
            attitude = [0,-1]
            position[y] -= int(step[1:])
        elif attitude == [0,-1]:
            attitude = [1,0]
            position[x] += int(step[1:])
        else: # attitude == [1,0]
            attitude = [0,1]
            position[y] += int(step[1:])

print abs(position[0]) + abs(position[1])
