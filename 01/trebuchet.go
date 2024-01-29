// AoC2023
// Day 01 - Trebuchet
package main

import (
	"bufio"
	"fmt"
	"m3w35/AoC2023/internal/utils"
	"os"
)

var (
    first_digit, second_digit, answer int
)

func calibrate(line string) {
    // IDEA:
    // start from line[0]; assign first digit found to first_digit.
    // start from line[len(line)] and move backwards; assign first
    // digit found to second_digit.
    // answer += first_digit + second_digit
    answer += 1
}

func main() {
    file_path := utils.OpenFile(os.Args)

    f, err := os.Open(file_path)
    utils.Check(err)

    s := bufio.NewScanner(f)
    s.Split(bufio.ScanLines)

    var line string
    for s.Scan() {
        line = s.Text()
        calibrate(line)
    }
    utils.Check(s.Err())

    fmt.Println(answer)
}
