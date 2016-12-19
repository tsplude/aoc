f = [l.strip() for l in open("part1_input.txt").readlines()]
foo = [s.split(' ') for s in f]
dims = [[int(d) for d in tri if d.isdigit()] for tri in foo]

count = 0
for d in dims:
    if sum(d) - max(d) > max(d):
        count += 1

print("Number of valid triangles:", count)
