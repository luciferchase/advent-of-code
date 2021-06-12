package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

// Part 1
func part1(input []int) (result int) {
    for index1, num1 := range input {
        for index2, num2 := range input {
            if ((index2 > index1) && ((num1 + num2) == 2020)) {
                result = num1 * num2
                break
            }
        }
    }
    return
}

// Part 2
func part2(input []int) (result int) {
    for index1, num1 := range input {
        for index2, num2 := range input {
            for index3, num3 := range input {
                if ((index3 > index2) && (index2 > index1) && ((num1 + num2 + num3) == 2020)) {
                    result = num1 * num2 * num3
                    break
                }
            }
        }
    }
    return
}

func main() {
    // Read data from input
    file, _ := os.Open("input.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var input []int

    for scanner.Scan() {
        num, _ := strconv.Atoi(scanner.Text())
        input = append(input, num)
    }

    // Print answer
    fmt.Println(part1(input))
    fmt.Println(part2(input))
    // [Finished in 1.8s]
}