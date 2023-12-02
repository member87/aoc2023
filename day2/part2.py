def check_game(game):
    count = {}
    for cubes in game.split('; '):
        for cube in cubes.split(', '):
            [num, color] = cube.split(' ')
            num = int(num)
            if color not in count:
                count[color] = num
            elif count[color] < num:  
                count[color] = num
    output = 1
    for color in count:
        output *= int(count[color])
    return output

def process_line(line):
    line = line[:-1]
    split = line.split(': ')
    return check_game(split[1])

with open('input.txt', 'r') as f:
    count = 0;
    for line in f:
        count += int(process_line(line))
    print(count)
