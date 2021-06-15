with open("input.txt", "r") as input_file:
    _input = input_file.read().splitlines()

# Both Parts
def both_parts(_input):
    seat_ids = []

    for seat in _input:
        rows = [i for i in range(128)]
        colums = [i for i in range(8)]
        _id = 0

        for i in seat[:7]:
            if (i == "F"):
                rows = rows[:len(rows) // 2]
            else:
                rows = rows[len(rows) // 2:]

        for i in seat[7:]:
            if (i == "L"):
                colums = colums[:len(colums) // 2]
            else:
                colums = colums[len(colums) // 2:]

        _id = rows[0] * 8 + colums[0]
        seat_ids.append(_id)

    return seat_ids

# Print answer
seat_ids = both_parts(_input)

# Part 1
print(max(seat_ids))

# Part 2
possible_seat_ids = [i for i in range(85, 891)]
my_id = [i for i in possible_seat_ids if i not in seat_ids]
print(my_id)
# [Finished in 307ms]