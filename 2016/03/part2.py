f = [l.strip() for l in open("part2_input.txt").readlines()]
foo = [s.split(' ') for s in f]
dims = [[int(d) for d in tri if d.isdigit()] for tri in foo]

count = 0
for i in range(0,len(dims),3):
    tri1 = [dims[i+0][0], dims[i+1][0], dims[i+2][0]]
    tri2 = [dims[i+0][1], dims[i+1][1], dims[i+2][1]]
    tri3 = [dims[i+0][2], dims[i+1][2], dims[i+2][2]]
    tris = [tri1, tri2, tri3]
    for tri in tris:
        if sum(tri) - max(tri) > max(tri):
            count += 1

print("Number of valid tris:", count)
