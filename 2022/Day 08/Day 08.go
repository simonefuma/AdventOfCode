package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strconv"

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (forest [][]int) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        var line []int
        for _, value := range scanner.Text() {
            n, err := strconv.Atoi(string(value))
            check(err)
            line = append(line, n)
        }
        forest = append(forest, line)
    }
    return
}

func visibility(forest [][]int) (num int) {
    for y := 1; y < len(forest)-1; y++ {
        for x := 1; x < len(forest[y])-1; x++ {

            var up bool
            for y2 := 0; y2 < y; y2++ {
                if forest[y2][x] >= forest[y][x] {
                    up = true
                    break
                }
            }
            if !up { continue }

            var down bool
            for y2 := y+1; y2 < len(forest); y2++ {
                if forest[y2][x] >= forest[y][x] {
                    down = true
                    break
                }
            }
            if !down{ continue }

            var left bool
            for x2 := 0; x2 < x; x2++ {
                if forest[y][x2] >= forest[y][x] {
                    left = true
                    break
                }
            }
            if !left { continue }

            var right bool
            for x2 := x+1; x2 < len(forest[y]); x2++ {
                if forest[y][x2] >= forest[y][x] {
                    right = true
                    break
                }
            }
            if !right{ continue }
            num++
        }
    }
    return len(forest)*len(forest[0])-num
}

func calculateScenicScore(forest [][]int) (scenicScore int) {
    for y := 1; y < len(forest)-1; y++ {
        for x := 1; x < len(forest[y])-1; x++ {
            var tempScenicScore int = 1
            var count int

            for y2 := y-1; y2 >= 0; y2-- {
                if forest[y2][x] < forest[y][x] {
                    count++
                } else {
                    count++
                    break
                }
            }
            tempScenicScore *= count
            count = 0

            for y2 := y+1; y2 < len(forest); y2++ {
                if forest[y2][x] < forest[y][x] {
                    count++
                } else {
                    count++
                    break
                }
            }
            tempScenicScore *= count
            count = 0

            for x2 := x-1; x2 >= 0; x2-- {
                if forest[y][x2] < forest[y][x] {
                    count++
                } else {
                    count++
                    break
                }
            }
            tempScenicScore *= count
            count = 0

            for x2 := x+1; x2 < len(forest[y]); x2++ {
                if forest[y][x2] < forest[y][x] {
                    count++
                } else {
                    count++
                    break
                }
            }
            tempScenicScore *= count

            if tempScenicScore > scenicScore { scenicScore = tempScenicScore }
        }
    }
    return
}

func main() {
    forest := read()
    fmt.Println(visibility(forest))
    fmt.Println(calculateScenicScore(forest))
}
