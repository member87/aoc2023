import numpy as np

card_count = {}

with open('input.txt') as f:
    count = 0
    for line in f.readlines():
        [card_num, cards] = line[:-1].split(': ')
        [win, have] = (cards.split('|'))
        print(card_num)
        card_num = card_num.split()[1]
        win = list(filter(None, win.split(' ')))
        have = list(filter(None, have.split(' ')))
        intersect = np.intersect1d(win, have)
        count = len(intersect)

        
        print('CARDS ' + card_num + ': ' + str(count) + ' matching numbers')
        card_count[str(card_num)] = card_count.get(str(card_num), 0) + 1
        for x in range(card_count.get(str(card_num), 0)):
            for i in range(count):
                card_count[str(i + int(card_num) + 1)] = (card_count.get(str(i + int(card_num) + 1), 0) + 1)
    print(sum(card_count.values()))
