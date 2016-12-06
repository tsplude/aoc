route = open("input.txt").readlines()[0].strip().split(', ')
atts = [x[:1] for x in route]                                                                      # attitudes, ['R', 'L', ...]
mags = [int(x[1:]) for x in route]                                                                 # magnitudes, [1, 4, 3, ...]
dvs = [[0,1], [1,0], [0,-1], [-1,0]]                                                               # unit vectors, facing North followed by 'R' = [0,1]
stepa = map(lambda x: -1 if x == 'L' else 1, atts)                                                 # map attitudess to indices
stepb = [dvs[i] for i in [e % 4 for e in [sum(stepa[0:x]) for x in range(1,len(stepa)+1)]]]        # map attitudess to unit vectors
stepc = map(lambda x: [stepb[x][0]*mags[x], stepb[x][1]*mags[x]], range(len(stepb)))               # direction vector * [magnitude, magnitude]
stepd = reduce(lambda x, y: [x[0] + y[0], x[1] + y[1]], stepc, [0,0])                              # reduce
print abs(stepd[0]) + abs(stepd[1])
