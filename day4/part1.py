import numpy as np

with open('input.txt') as f:
    count = 0
    for line in f.readlines():
        [win, have] = (line[:-1].split(': ')[1].split('|'))
        win = list(filter(None, win.split(' ')))
        have = list(filter(None, have.split(' ')))
        intersect = np.intersect1d(win, have)
        if(len(intersect) == 0):
            continue
        elif (len(intersect) == 1):
            count += 1
        else:
            count += 2**(len(intersect)-1)
    print(count)
