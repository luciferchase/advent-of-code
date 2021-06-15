package main
import (
    "bufio"
    "fmt"
    "os"
    "sort"
)

// Both parts
func bothParts(input []string) (seat_ids []int) {
    for _, seat := range input {
        var rows []int
        var colums []int

        for i := 0; i < 128; i++ {
            rows = append(rows, i)
        }
        for i := 0; i < 8; i++ {
            colums = append(colums, i)
        }

        for index, value := range seat {
            if (index < 7) {
                if (string(value) == "F") {
                    rows = rows[:len(rows) / 2]
                } else {
                    rows = rows[len(rows) / 2:]
                }
            } else {
                if (string(value) == "L") {
                    colums = colums[:len(colums) / 2]
                } else {
                    colums = colums[len(colums) / 2:]
                }
            }
        }
        id := rows[0] * 8 + colums[0]
        seat_ids = append(seat_ids, id)
    }

    sort.Ints(seat_ids)
    return
}

func checkMembership(array []int, element int) bool {
    for _, i := range array {
        if (i == element) { return true }
    }
    return false
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

    // Print answer
    seat_ids := bothParts(input)
    min_id, max_id := seat_ids[0], seat_ids[len(seat_ids) - 1]

    // Part 1
    fmt.Println(max_id)

    // Part 2
    var possibleIds []int
    var my_id int

    for i := min_id; i <= max_id; i++ {
        possibleIds = append(possibleIds, i)
    }
    for _, id := range(possibleIds) {
        if !checkMembership(seat_ids, id) { my_id = id }
    }

    fmt.Println(my_id)
    // [Finished in 4.2254ms]
}