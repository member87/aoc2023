bad_chars = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.']
def get_symbols_indexs(string):
    arr = [*string[:-1]]
    output = []
    for i in range(len(arr)):
        if arr[i] not in bad_chars:
            output.append(i)
    return output

good_chars = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9']
def get_numbers_indexes(string):
    output = []
    start = -1
    end = 0
    searching = False
    for i in range(len(string)):
        if(string[i] in good_chars):
            searching = True
        else:
            end = i - 1
            searching = False
            if(start != -1):
                output.append([start, end, (int)(string[start:end+1])])
                start = -1

        if (searching and start == -1):
            start = i
    return output

def find_neighbours(array, num):
    output = []
    for sub_array in array:
        check_array = [i for i in range(sub_array[0]-1, sub_array[1]+2)]
        if num in check_array:
            output.append(sub_array[2])
    return output

def number_adjacent(before, current, after, symbols):
    if(len(symbols) == 0):
        return 0

    output = 0
    for i in range(len(symbols)):
        matches = []
        if (found := find_neighbours(before, symbols[i])):
            matches += found
        if (found := find_neighbours(current, symbols[i])):
            matches += found
        if (found := find_neighbours(after, symbols[i])):
            matches += found
        if(len(matches) == 2):
            total = 1
            for i in matches:
                total *= i
            output += total

    return output

def process_lines(before, current, after):
    before_num_indexes = get_numbers_indexes(before)
    current_num_indexes = get_numbers_indexes(current)
    after_num_indexes = get_numbers_indexes(after)

    symbol_indexes = get_symbols_indexs(current)

    return number_adjacent(before_num_indexes, current_num_indexes, after_num_indexes, symbol_indexes)


with open('input.txt', 'r') as f:
    lines = f.readlines()
    output = 0
    for i in range(len(lines)):
        if i == 0:
            output += process_lines("", lines[i], lines[i+1])
        elif i == len(lines) - 1:
            output += process_lines(lines[i-1], lines[i], "")
        else:
            output += process_lines(lines[i-1], lines[i], lines[i+1])
    print(output)
