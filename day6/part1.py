from math import floor
from operator import mul
from functools import reduce

with open('input.txt') as f:
    [times, distances] = f.read().splitlines()
    times = times.split(': ')[1].split()
    distances = distances.split(': ')[1].split()

    solutions = []
    for i, time in enumerate(times):
        time = int(time)
        count = 0
        for pointer in range(floor(time / 2)+1):
            end_pointer = time-pointer
            if((pointer * end_pointer) > int(distances[i])):
                count += 1 if (pointer == end_pointer) else 2
        solutions.append(count)

    print(reduce(mul, solutions, 1))


