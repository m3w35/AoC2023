// AoC2023
// internal utilities & helper functions
package utils

import (
	"fmt"
	"os"
)

func Check(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func OpenFile(args []string) string {
    if len(args) == 2 {
        return os.Args[1]
    }
    fmt.Println("provide a file path for the input, dickhead.")
    return ""
}
