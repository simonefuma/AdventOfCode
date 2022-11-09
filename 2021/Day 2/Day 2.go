package main

import "os"
import "log"
import "fmt"
import "bufio"
import "strconv"
import "strings"

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (input []string) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        input = append(input, scanner.Text())
    }
    return
}

func calculatePosition(commands []string) int {
    var horizontalPosition, depth, aim int
    for _, command := range commands {
        c := strings.Split(command, " ")
        value, err := strconv.Atoi(c[1])
        check(err)

        switch c[0] {
        case "forward":
            horizontalPosition += value
            depth += aim * value
        case "down":
            aim += value
        case "up":
            aim -= value
        }
    }
    return horizontalPosition * depth
}

func main() {
    var commands []string = read()
    fmt.Println("The result is:", calculatePosition(commands))
}
