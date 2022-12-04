package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strings"

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() int {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    var points int
    for scanner.Scan() {
        var round []string = strings.Split(scanner.Text(), " ")
        switch  round[1] {
        case "X":
            switch round[0] {
            case "A":
                points += 3
            case "B":
                points += 1
            case "C":
                points += 2
            }
        case "Y":
            switch round[0] {
            case "A":
                points += 4
            case "B":
                points += 5
            case "C":
                points += 6
            }
        case "Z":
            switch round[0] {
            case "A":
                points += 8
            case "B":
                points += 9
            case "C":
                points += 7
            }
        }
    }
    return points
}

func main() {
    fmt.Println(read())
}
