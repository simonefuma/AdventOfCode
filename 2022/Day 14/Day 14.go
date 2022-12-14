package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strings"
import "strconv"

type Pair struct {
    x,y int
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (pairs map[Pair]bool, maxY int) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    pairs = make(map[Pair]bool)
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), " -> ")
        for i := 0; i < len(line)-1; i++ {
            strx1, stry1 := split(line[i])
            strx2, stry2 := split(line[i+1])
            if strx1 == strx2 {
                x, err := strconv.Atoi(strx1)
                check(err)
                y1, err := strconv.Atoi(stry1)
                check(err)
                y2, err := strconv.Atoi(stry2)
                check(err)
                if y1 > y2 { y1, y2 = y2, y1 }
                if maxY == 0 {
                    maxY = y2
                } else if y1 > maxY { maxY = y1 }
                for j := y1; j <= y2; j++ { pairs[Pair{x,j}] = true }
            } else {
                y, err := strconv.Atoi(stry1)
                check(err)
                if maxY == 0 {
                    maxY = y
                } else if y > maxY { maxY = y }
                x1, err := strconv.Atoi(strx1)
                check(err)
                x2, err := strconv.Atoi(strx2)
                check(err)
                if x1 > x2 { x1, x2 = x2, x1 }
                for j := x1; j <= x2; j++ { pairs[Pair{j,y}] = true }
            }
        }
    }
    return pairs, maxY+1
}

func split(str string) (string, string) {
    value := strings.Split(str, ",")
    return value[0], value[1]
}

func run(pairs map[Pair]bool, maxY int) (sands map[Pair]bool) {
    sands = make(map[Pair]bool)
    for true {
        sandPosition := Pair{500, 0}
        for true {
            if sandPosition.y == maxY {
                sands[sandPosition] = true
                break
            }
            if !(pairs[Pair{sandPosition.x, sandPosition.y+1}] || sands[Pair{sandPosition.x, sandPosition.y+1}]) {
                sandPosition.y++
            } else if !(pairs[Pair{sandPosition.x-1, sandPosition.y+1}] || sands[Pair{sandPosition.x-1, sandPosition.y+1}]) {
                sandPosition.x--
                sandPosition.y++
            } else if !(pairs[Pair{sandPosition.x+1, sandPosition.y+1}] || sands[Pair{sandPosition.x+1, sandPosition.y+1}]) {
                sandPosition.x++
                sandPosition.y++
            } else if sandPosition.x == 500 && sandPosition.y == 0 {
                sands[sandPosition] = true
                return
            } else {
                sands[sandPosition] = true
                break
            }
        }
    }
    return
}

func main() {
    pairs, maxY := read()
    sands := run(pairs, maxY)
    fmt.Println(len(sands))
}
