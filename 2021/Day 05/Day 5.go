package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strconv"
import "strings"

type Coordinated struct {
    x, y int
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() map[Coordinated]int {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    var coordinates map[Coordinated]int = make(map[Coordinated]int)
    var values [4]int
    for scanner.Scan() {
        for i, stringValue := range strings.Split(strings.Replace(scanner.Text(), " -> ", ",", 1), ",") {
            intValue, err := strconv.Atoi(stringValue)
            check(err)
            values[i] = intValue
        }

        // I increase every point where the line passes
        if values[0] == values[2] {
            if values[1] > values[3] {
                values[1], values[3] = values[3], values[1]
            }
            for i := values[1]; i <= values[3]; i++ {
                coordinates[Coordinated{values[0], i}]++
            }
        } else if values[1] == values[3] {
            if values[0] > values[2] {
                values[0], values[2] = values[2], values[0]
            }
            for i := values[0]; i <= values[2]; i++ {
                coordinates[Coordinated{i, values[1]}]++
            }
        } else {
            var m int = (values[1]-values[3]) / (values[0]-values[2])
            var q int = ((values[0]*values[3])-(values[2]*values[1])) / (values[0]-values[2])
            if values[0] > values[2] {
                values[0], values[1], values[2], values[3] = values[2], values[3], values[0], values[1]
            }
            coordinates[Coordinated{values[0], values[1]}]++
            coordinates[Coordinated{values[2], values[3]}]++
            for i := values[0]+1; i < values[2]; i++ {
                coordinates[Coordinated{i, m*i + q}]++
            }
        }
    }
    return coordinates
}

func countsOverlappingLines(coordinates map[Coordinated]int) (count int) {
    for _, coverPoint := range coordinates {
        if coverPoint > 1 {
            count++
        }
    }
    return
}

func main() {
    fmt.Println("The number of points where at least two lines overlap is:", countsOverlappingLines(read()))
}
