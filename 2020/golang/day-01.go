package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

// Part 1
func part1(input []int) (result int) {
    for index1 := 0; index1 < len(input); index1++ {
        for index2 := index1 + 1; index2 < len(input); index2++ {
            if (input[index1] + input[index2] == 2020) {
                result = (input[index1]) * (input[index2])
                break
            }
        }
    }
    return
}

// Part 2
func part2(input []int) (result int) {
    for index1 := 0; index1 < len(input); index1++ {
        for index2 := index1 + 1; index2 < len(input); index2++ {
            for index3 := index2 + 1; index3 < len(input); index3++ {
                if (input[index1] + input[index2] + input[index3] == 2020) {
                    result = (input[index1]) * (input[index2]) * (input[index3])
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
    // [Finished in 368.0393ms]
}