# Read from input
with open("input.txt", "r") as input_file:
    _input = input_file.read().splitlines()

# Part 1
def part_1(_input):
    count = 0
    for i in _input:
        limits, char, password = i.split()

        lower_limit, upper_limit = [int(i) for i in limits.split("-")]
        char = char[0]

        if (password.count(char) in [i for i in range(lower_limit, upper_limit + 1)]):
            count += 1

    return count

# Part 2
def part_2(_input):
    count = 0
    for i in _input:
        indexes, char, password = i.split()

        index1, index2 = [(int(i) - 1) for i in indexes.split("-")]
        char = char[0]
        char_string = password[index1] + password[index2]

        if (char_string.count(char) == 1):
            count += 1

    return count

# Print answer
print(part_1(_input))
print(part_2(_input))
# [Finished in 434ms]