package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Part 1
func checkMembership(input, element string) bool {
    for _, i := range input {
        value := string(i)
        if (value == element) { return true }
    }
    return false
}

func part1(groups [][]string) (count int) {
    for _, group := range groups {
        var answers string
        for _, individualResponse := range group {
            for _, question := range individualResponse {
                if !checkMembership(answers, string(question)) {
                    answers += string(question)
                }
            }
        }
        count += len(answers)
    }
    return
}

// Part 2
func part2(groups [][]string) (count int) {
    for _, group := range groups {
        answers := strings.Join(group, "")
        var answersChecked string
        for _, i := range answers {
            question := string(i)
            if (strings.Count(answers, question) == len(group) && !checkMembership(answersChecked, question)) {
                answersChecked += question
                count += 1
            }
        }
    }
    return
}

func main() {
    // Read input
    file, _ := os.Open("input.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var input []string

    for scanner.Scan() {
        input = append(input, scanner.Text())
    }
    // Add EOF newline
    input = append(input, "")

    var groups [][]string
    var tempGroup []string
    for _, i := range input {
        if (i == "") {
            groups = append(groups, tempGroup)
            tempGroup = nil
        } else {
            tempGroup = append(tempGroup, i)
        }
    }

    // Print answers
    fmt.Println(part1(groups))
    fmt.Println(part2(groups))
}