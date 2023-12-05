maps = {}
def get_mapping(value, map_key):
    for i in maps[map_key]:
        if(len(i) == 0): return value
        if(value >= i[1] and value < i[1] + i[2]):
            return i[0] + (value - i[1])
    return value

with open('input.txt') as f:
    for mappings in (f.read().split("\n\n")):
        [key, values] = mappings.split(':')
        key = key.split(' ')[0]
        values = values[1:]

        maps[key] = []
        for value in (values.split('\n')):
            maps[key].append([int(x) for x in value.split()])
    
    order = ['seed-to-soil', 'soil-to-fertilizer', 'fertilizer-to-water', 'water-to-light', 'light-to-temperature', 'temperature-to-humidity', 'humidity-to-location']
    localtion_list = []
    
    for seed in maps['seeds'][0]:
        current_value = seed
        for key in order:
            current_value = get_mapping(current_value, key)
            if(key == 'humidity-to-location'):
                localtion_list.append(current_value)

    print(min(localtion_list))
