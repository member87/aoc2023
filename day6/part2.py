from math import floor

with open('input.txt') as f:
    [time, distance] = f.read().splitlines()
    time = int(time.split(': ')[1].replace(' ', ''))
    distance = int(distance.split(': ')[1].replace(' ', ''))

    count = 0
    for pointer in range(floor(time / 2)+1):
        end_pointer = time-pointer
        if((pointer * end_pointer) > int(distance)):
            count += 1 if (pointer == end_pointer) else 2
    
    print(count)

