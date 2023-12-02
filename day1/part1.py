import re

def process_line(line):
    nums = re.findall("[0-9]", line)
    if(len(nums) <= 0):
        return 0
    return (int)(nums[0] + nums[len(nums)-1])

with open('input.txt') as f:
    total = 0
    for line in f.readlines():
        total += process_line(line)
    print(total)
