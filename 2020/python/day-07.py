import re
import json

with open("input.txt") as input_file:
    data = input_file.read().splitlines()
    _input = [i.split() for i in data]

# Part 1
# Make relation between all the colours
def colours_relation(_input):
    relation = {}

    for rule in _input:
        base_colour = " ".join([rule[0], rule[1]])
        relation[base_colour] = {}
        
        for i in range(len(rule)):
            match = re.match(r"\d", rule[i])
            if match:
                relation[base_colour][" ".join([rule[i + 1], rule[i + 2]])] = int(rule[i])
    return relation

def part_1(_input):
    count = 0
    relation = _input.copy()

    # Strip relation of extra information
    for base_colour in relation:
        relation[base_colour] = list(relation[base_colour].keys())

    # Add indirectly related colours
    for base_colour in relation:
        for sub_colour in relation[base_colour]:
            relation[base_colour] += [i for i in relation[sub_colour] if i not in relation[base_colour]]

    for base_colour in relation:
        if ("shiny gold" in relation[base_colour]):
            count += 1
    return count

# Part 2
def add_no_of_bag(relation, colour):
    count = 0
    if not relation[colour]:
        return count
    else:
        count += sum(list(relation[colour].values()))
        for sub_colour in relation[colour]:
            count += relation[colour][sub_colour] * add_no_of_bag(relation, sub_colour)
    return count

def part_2(_input):
    relation = _input.copy()
    count = sum(list(relation["shiny gold"].values()))
    for colour in relation["shiny gold"]:
        count += relation["shiny gold"][colour] * add_no_of_bag(relation, colour)
    return count

# Print answers
relation = colours_relation(_input)
print(part_1(relation))
print(part_2(relation))
