# Read from input
with open("input.txt", "r") as input_file:
    _input = input_file.read().splitlines()

# Both Parts
def both_parts(right_shift, down_shift, _input):
    count = 0
    index = 0

    for i in range(0, len(_input), down_shift):
        if (index > 30):
            index -= 31
        if (_input[i][index] == "#"):
            count += 1
        index += right_shift
    return count

# Print answer
# Part 1
print(both_parts(3, 1, _input))

# Part 2
result = 1
for i in [[1, 1], [3, 1], [5, 1], [7, 1], [1, 2]]:
    result *= both_parts(i[0], i[1], _input)
print(result)