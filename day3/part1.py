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
    if(num in array):
        return True
    elif(num + 1 in array):
        return True
    elif(num - 1 in array):
        return True
    return False

def number_adjacent(before, current, after, nums):
    output = 0
    for i in range(len(nums)):
        for k in range((int)(nums[i][0]), (int)(nums[i][1])+1):
            if find_neighbours(before, k) or find_neighbours(after, k):
                output += nums[i][2]
                break
        if(nums[i][0] - 1 in current or nums[i][1] + 1 in current):
            output += nums[i][2]
    return output

def process_lines(before, current, after):
    before_symbols_indexes = get_symbols_indexs(before)
    current_symbols_indexes = get_symbols_indexs(current)
    after_symbols_indexes = get_symbols_indexs(after)

    number_index = get_numbers_indexes(current)
    return number_adjacent(before_symbols_indexes, current_symbols_indexes, after_symbols_indexes, number_index)

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
