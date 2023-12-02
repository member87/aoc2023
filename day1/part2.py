import re

words = {
        'one': '1',
        'two': '2',
        'three': '3',
        'four': '4',
        'five': '5',
        'six': '6',
        'seven': '7',
        'eight': '8',
        'nine': '9',
}

regex = f"{('|').join(words.keys())}"
def process_line(line):
    nums = re.findall(f"(?=({regex}|[1-9]))", line)
    if(len(nums) <= 0):
        return 0

    start_num = get_number(nums[0])
    end_num = get_number(nums[len(nums)-1])
    return (int)(f"{start_num}{end_num}")

def get_number(num):
    if num in words:
        return words[num]
    return num


with open('input.txt') as f:
    total = 0
    for line in f.readlines():
        total += process_line(line)
    print(total)
