import sys

input_file = sys.argv[1]
instructions = open(input_file, 'r').read()

level = 0
count = 0
for step in instructions:
    count += 1
    if step == '(':
        level += 1
    else:
        level -= 1
    if level < 0:
        print count
        break
