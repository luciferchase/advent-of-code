import re

with open("input.txt", "r") as input_file:
    _input = input_file.read().splitlines()
    # Add the EOF newline
    _input.append("")

    # Format input
    temp_input = []
    passport = []
    for i in _input:
        if not i:
            temp_input.append(passport)
            passport = []
        else:
            passport += i.split()
    _input = temp_input

# Part 1
def part_1(_input):
    required_fields = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]
    valid_passports = []

    for passport in _input:
        fields_in_passport = [field.split(":")[0] for field in passport]

        for field in required_fields:
            if (field not in fields_in_passport):
                break
        else:
            valid_passports.append(passport)
    return valid_passports

# Part 2
# Compile regex beforehand
byr = re.compile("^(19[2-9]\\d)$|^(200[0-2])$")
iyr = re.compile("^(201\\d)$|^(2020)$")
eyr = re.compile("^(202\\d)$|^(2030)$")
hgt = re.compile("^(1[5-8]\\dcm)$|^(19[0-3]cm)$|^(59in)$|^(6\\din)$|^(7[0-6]in)$")
hcl = re.compile("^#([\\da-f]){6}$")
pid = re.compile("^\\d{9}$")

def validate(key, value):
    valid = True
    
    if (key == "byr"):
        valid =  bool(byr.match(value))
    elif (key == "ecl" and value not in ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]):
        valid =  False
    elif (key == "eyr"):
        valid =  bool(eyr.match(value))
    elif (key == "hcl"):
        valid =  bool(hcl.match(value))
    elif (key == "hgt"):
        valid =  bool(hgt.match(value))
    elif (key == "iyr"):
        valid =  bool(iyr.match(value))
    elif (key == "pid"):
        valid =  bool(pid.match(value))
    return valid

def part_2(valid_passports):
    count = 0
    for passport in valid_passports:
        for i in passport:
            key, value = i.split(":")
            if not validate(key, value):
                break
        else:
            count += 1
    return count

# Print answers
valid_passports = part_1(_input)

print(len(valid_passports))
print(part_2(valid_passports))
# [Finished in 301ms]