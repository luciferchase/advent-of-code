# Read from input
with open("input.txt", "r") as input_file:
    _input = input_file.read().splitlines()
    _input = [int(i) for i in _input]

# Part 1
def part_1(_input):
    for i in range(len(_input)):
        for j in range(i + 1, len(_input)):
            if ((_input[i] + _input[j]) == 2020):
                return (_input[i] * _input[j])

# Part 2
def part_2(_input):
    for i in range(len(_input)):
        for j in range(i + 1, len(_input)):
            for k in range(j + 1, len(_input)):
                if ((_input[i] + _input[j] + _input[k]) == 2020):
                    return (_input[i] * _input[j] * _input[k])

# Print answer
print(part_1(_input))
print(part_2(_input))
# [Finished in 461ms]