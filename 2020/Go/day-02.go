package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// Part 1
func part1(input [][]string) (count int) {
    for _, value := range input {
        limit, char, password := strings.Split(value[0], "-"), string(value[1][0]), value[2]
        lower_limit, _ := strconv.Atoi(limit[0]); upper_limit, _ := strconv.Atoi(limit[1])
        char_count := strings.Count(password, char)

        if (char_count >= lower_limit && char_count <= upper_limit) {
            count++
        }
    }
    return
}

// Part 2
func part2(input [][]string) (count int) {
    for _, value := range input {
        indexes, char, password := strings.Split(value[0], "-"), string(value[1][0]), value[2]
        lower_index, _ := strconv.Atoi(indexes[0]); upper_index, _ := strconv.Atoi(indexes[1])
        char_string := string(password[lower_index - 1]) + string(password[upper_index - 1])

        if (strings.Count(char_string, char) == 1) {
            count++
        }
    }
    return
}


func main() {
    // Read data from input
    file, _ := os.Open("input.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var input [][]string

    for scanner.Scan() {
        value := strings.Split(scanner.Text(), " ")
        input = append(input, value)
    }

    // Print answer
    fmt.Println(part1(input))
    fmt.Println(part2(input))
    // [Finished in 518.8µs]
}