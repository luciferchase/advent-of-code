package main
import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
)

// Part 1
func checkMembership(array []string, element string) bool {
    for _, value := range array {
        if (value == element) {
            return true
        }
    }
    return false
}

func part1(input [][]string) (validPassports [][]string) {
    requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

    for _, passport := range input {
        var fieldsInPassport []string
        for _, i := range passport {
            fieldsInPassport = append(fieldsInPassport, strings.Split(i, ":")[0])
        }

        isValid := true
        for _, field := range requiredFields {
            if !checkMembership(fieldsInPassport, field) {
                isValid = false
                break
            }
        }
        if isValid {
            validPassports = append(validPassports, passport)
        }
    }
    return
}

// Part 2
// Compile regex beforehand
var byr = regexp.MustCompile("^(19[2-9]\\d)$|^(200[0-2])$")
var iyr = regexp.MustCompile("^(201\\d)$|^(2020)$")
var eyr = regexp.MustCompile("^(202\\d)$|^(2030)$")
var hgt = regexp.MustCompile("^(1[5-8]\\dcm)$|^(19[0-3]cm)$|^(59in)$|^(6\\din)$|^(7[0-6]in)$")
var hcl = regexp.MustCompile("^#([\\da-f]){6}$")
var pid = regexp.MustCompile("^\\d{9}$")

func validate(key, value string) (valid bool) {
    switch key {
    case "cid":
        valid = true
    case "byr":
        valid = byr.MatchString(value)
    case "iyr":
        valid = iyr.MatchString(value)
    case "eyr":
        valid = eyr.MatchString(value)
    case "hgt":
        valid = hgt.MatchString(value)
    case "hcl":
        valid = hcl.MatchString(value)
    case "ecl":
        valid = checkMembership([]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}, value)
    case "pid":
        valid = pid.MatchString(value)
    }
    return
}

func part2(validPassports [][]string) (count int) {
    for _, passport := range validPassports {
        isValid := true
        for _, i := range passport {
            list := strings.Split(i, ":")
            if !validate(list[0], list[1]) {
                isValid = false
                break
            }
        }
        if isValid {
            count += 1
        }
    }
    return
}

func main() {
    // Read data from input
    file, _ := os.Open("input.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var tempInput []string

    for scanner.Scan() {
        tempInput = append(tempInput, scanner.Text())
    }
    // Add the EOF newline
    tempInput = append(tempInput, "")
    // fmt.Println(tempInput)

    // Format input
    var input [][]string
    var passport []string
    for _, i := range tempInput {
        if (i == "") {
            input = append(input, passport)
            passport = nil
        } else {
            passport = append(passport, strings.Split(i, " ")...)
        }
    }

    // Print answer
    validPassports := part1(input)
    fmt.Println(len(validPassports))

    fmt.Println(part2(validPassports))
    // [Finished in 1.5619ms]
}