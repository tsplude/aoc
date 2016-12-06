f = open("input.txt").readlines()
sequence = [l.strip() for l in f]
print sequence

convert = {
    "U": [0,1], "D": [0,-1], "R": [1,0], "L": [-1,0]
}

bounds = {
    "U": 2, "D": 0, "L": 0, "R": 2
}

keypad = [[7,8,9],
          [4,5,6],
          [1,2,3]]

x = 0
y = 1

def next_step(prev, cur):
    if cur == convert["U"]: return [prev[x], min(prev[y] + cur[y], bounds["U"])]
    if cur == convert["D"]: return [prev[x], max(prev[y] + cur[y], bounds["D"])]
    if cur == convert["L"]: return [max(prev[x] + cur[x], bounds["L"]), prev[y]]
    if cur == convert["R"]: return [min(prev[x] + cur[x], bounds["R"]), prev[y]]

start = [1,1]
for step in sequence:
    steps = map(lambda x: convert[x], [s for s in step])
    start = reduce(next_step, steps, start)
    print keypad[start[y]][start[x]]

