package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strconv"

type Pair struct {
    x,y int
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (istruction []string) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        istruction = append(istruction, scanner.Text())
    }
    return
}

func run(istruction []string) (intensity int, draws map[Pair]bool) {
    var register, cycle int = 1, 1
    var line int
    var temp int
    var isRunning bool
    draws = make(map[Pair]bool)
    for i := 0; i < len(istruction); i++ {
        if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
            intensity += cycle * register
        }
        var t int = cycle%40
        if t == 0 { t = 40 }
        if t == register {
            draws[Pair{line, register}] = true
        } else if t == register+1 {
            draws[Pair{line, register+1}] = true
        } else if t == register+2 {
            draws[Pair{line, register+2}] = true
        }

        if isRunning {
            register += temp
            isRunning = false
        } else {
            if istruction[i][:4] == "addx" {
                isRunning = true
                n, err := strconv.Atoi(istruction[i][5:])
                check(err)
                temp = n
                i--
            }
        }
        if t == 40 { line++ }
        cycle++
    }
    return
}

func drawMonitor(monitor map[Pair]bool) {
    for i := 0; i < 6; i++ {
        for j := 1; j < 41; j++ {
            if monitor[Pair{i,j}] {
                fmt.Print("#")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
}

func main() {
    intensity, monitor := run(read())
    fmt.Println("intensity:", intensity)
    fmt.Println()
    drawMonitor(monitor)
}
