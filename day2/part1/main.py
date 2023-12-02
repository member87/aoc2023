

LIMITS = {
        'RED': 12,
        'GREEN': 13,
        'BLUE': 14
}




def check_cubes(cubes):
    for cube in cubes.split(', '):
        [count, color] = cube.split(' ')
        if(int(count) > LIMITS[color.upper()]):
            return False
    return True

def check_game(game):
    for cubes in game.split('; '):
        if not check_cubes(cubes):
            return False
    return True


def process_line(line):
    line = line[:-1]
    split = line.split(': ')
    game_id = split[0].split(' ')[1]
    return game_id if check_game(split[1]) else 0

with open('input.txt', 'r') as f:
    count = 0;
    for line in f:
        count += int(process_line(line))

    print(count)


