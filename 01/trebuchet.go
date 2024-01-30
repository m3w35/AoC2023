// AoC2023
// Day 01 - Trebuchet
package main

import (
	"bufio"
	"fmt"
	"m3w35/AoC2023/internal/utils"
	"os"
	"strconv"
	"strings"
)

const (
    digits string = "0123456789"
)

var (
    answer int
    first_d, last_d string
    first_i, last_i int
)


func calibrate(line string) {
    // Get indices of first and last digits in line.
    first_i = strings.IndexAny(line, digits)
    last_i = strings.LastIndexAny(line, digits)

    if first_i != -1 {
        first_d = string(line[first_i])
    } else { // If no digit found, set first_d to 0.
        first_d = "0"
    }

    if last_i != -1 {
        last_d = string(line[last_i])
    } else { // If no digit found, set last_d to 0.
        last_d = "0"
    }

    // Join first_d to last_d and convert to int
    s := first_d+last_d
    c_value, err := strconv.Atoi(s)
    utils.Check(err)
    answer += c_value
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
