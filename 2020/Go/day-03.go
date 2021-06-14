package main
import (
    "bufio"
    "fmt"
    "os"
)

// Both Parts
func bothParts(right_shift, down_shift int, input []string) (count int) {
    index := 0
    for i := 0; i < len(input); i += down_shift {
        if (index > 30) {
            index -= 31
        }
        if (string(input[i][index]) == "#") {
            count += 1
        }
        index += right_shift
    }
    return
} 

func main() {
    // Read data from input
    file, _ := os.Open("input.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var input []string

    for scanner.Scan() {
        input = append(input, scanner.Text())
    }

    // Print answer
    // Part 1
    fmt.Println(bothParts(3, 1, input))

    // Part 2
    result := 1
    for _, args := range [][]int{[]int{1, 1}, []int{3, 1}, []int{5, 1}, []int{7, 1}, []int{1, 2}} {
        result *= bothParts(args[0], args[1], input)
    }
    fmt.Println(result)
    // [Finished in 532.3µs]
}