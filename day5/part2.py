maps = {}
def get_mapping(value, map_key):
    for i in maps[map_key]:
        if(len(i) == 0): return value
        if(value >= i[1] and value < i[1] + i[2]):
            return i[0] + (value - i[1])
    return value

order = ['seed-to-soil', 'soil-to-fertilizer', 'fertilizer-to-water', 'water-to-light', 'light-to-temperature', 'temperature-to-humidity', 'humidity-to-location']
def get_location(value):
    current_value = value
    for key in order:
        current_value = get_mapping(current_value, key)
    return current_value

def get_seeds(seed_list):
    current_location = -1
    for i in range(0, len(seed_list), 2):
        for x in range(seed_list[i], seed_list[i] + seed_list[i + 1]):
            if((location := get_location(x)) <= current_location or (current_location == -1)):
                current_location = location
                print("updated location: " + str(current_location))
    return current_location

with open('input.txt') as f:
    for mappings in (f.read().split("\n\n")):
        [key, values] = mappings.split(':')
        key = key.split(' ')[0]
        values = values[1:]

        maps[key] = []
        for value in (values.split('\n')):
            maps[key].append([int(x) for x in value.split()])
    print(get_seeds(maps['seeds'][0]))
