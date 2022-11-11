package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strconv"
import "strings"

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (positions []int, min, max int) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    scanner.Scan()
    for i, stringPosition := range strings.Split(scanner.Text(), ",") {
        intPosition, err := strconv.Atoi(stringPosition)
        check(err)
        positions = append(positions, intPosition)

        if i == 0 {
            min, max = intPosition, intPosition
        } else {
            if intPosition < min {
                min = intPosition
            } else if intPosition > max {
                max = intPosition
            }
        }
    }
    return
}

func distance(a, b int) int {
    if a < b { return b-a }
    return a-b
}

func align(positions []int, min, max int) (fuel int) {
    for x := min; x < max; x++ {
        var potentialFuel int
        for i := 0; i < len(positions); i++ {
            var distance int = distance(x, positions[i])
            potentialFuel += (distance*(distance+1)) / 2
        }
        if x == min {
            fuel = potentialFuel
        } else if fuel > potentialFuel {
            fuel = potentialFuel
        } else {
            break
        }
    }
    return fuel
}

func main() {
    fmt.Println("Fuel:", align(read()))

}
