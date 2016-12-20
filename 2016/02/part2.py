import functools

f = open("input.txt").readlines()
sequence = [l.strip() for l in f]

convert = {
    "U": [0,1], "D": [0,-1], "R": [1,0], "L": [-1,0]
}

keypad = [["-","-","1","-","-"],
          ["-","2","3","4","-"],
          ["5","6","7","8","9"],
          ["-","A","B","C","-"],
          ["-","-","D","-","-"]]

x = 0
y = 1

start = [2,2]



