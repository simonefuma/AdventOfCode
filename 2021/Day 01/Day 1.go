package main

import "os"
import "log"
import "fmt"
import "bufio"
import "strconv"

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (input []int) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        n, err := strconv.Atoi(scanner.Text())
        check(err)
        input = append(input, n)
    }
    return
}

func sumSlidingWindow(measures []int) (newMeasures []int) {
    for i := 0; i < len(measures); i++ {
        if i > 1 {
            var sum int = measures[i] + measures[i-1] + measures[i-2]
            newMeasures = append(newMeasures, sum)
        }
    }
    return
}

func increased(measures []int) (n int) {
    for i := 1; i < len(measures); i++ {
        if measures[i] > measures[i-1] {
            n++
        }
    }
    return
}

func main() {
    var measures []int = sumSlidingWindow(read())
    fmt.Println("The number of increments is:", increased(measures))
}
