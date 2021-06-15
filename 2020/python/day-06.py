with open("input.txt") as input_file:
    _input = input_file.read().splitlines()
    # Add EOF newline
    _input.append("")

    # Format input
    groups = []
    temp_group = []

    for i in _input:
        if (i == ""):
            groups.append(temp_group)
            temp_group = []
        else:
            temp_group.append(i)

# Part 1
def part_1(groups):
    count = 0
    for group in groups:
        answers = ""
        for individual_response in group:
            for question in individual_response:
                if (question not in answers):
                    answers += question
        count += len(answers)
    return count

def part_2(groups):
    count = 0
    for group in groups:
        answers = "".join(group)
        answers_checked = ""
        for question in answers:
            if (answers.count(question) == len(group) and question not in answers_checked):
                answers_checked += question
                count += 1
    return count

# Print answers
print(part_1(groups))
print(part_2(groups))
# [Finished in 301ms]