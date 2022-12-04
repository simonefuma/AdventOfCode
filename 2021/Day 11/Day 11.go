package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strconv"

type pair struct {
    x, y int
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (octopuses [][]int) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    octopuses = append(octopuses, []int{})
    for scanner.Scan() {
        var row []int
        row = append(row, 0)
        for _, octopusString := range scanner.Text() {
            octopusInt, err := strconv.Atoi(string(octopusString))
            check(err)
            row = append(row, octopusInt)
        }
        row = append(row, 0)
        octopuses = append(octopuses, row)
    }
    octopuses = append(octopuses, []int{})
    for i := 0; i < len(octopuses[1]); i++ {
        octopuses[0] = append(octopuses[0], 0)
        octopuses[len(octopuses)-1] = append(octopuses[len(octopuses)-1], 0)
    }
    return
}

func countFleshes(octopuses [][]int) (step int){
    for count := 1; true; count++ {
        for i := 1; i < len(octopuses)-1; i++ {
            for j := 1; j < len(octopuses[i])-1; j++ {
                octopuses[i][j]++
            }
        }
        var check map[pair]bool = make(map[pair]bool)
        for i := 1; i < len(octopuses)-1; i++ {
            for j := 1; j < len(octopuses[i])-1; j++ {
                if octopuses[i][j] > 9 {
                    if fleshes(octopuses, check, i, j) == (len(octopuses)-2) * (len(octopuses[0])-2) { return count }
                }
            }
        }

        for i := 1; i < len(octopuses)-1; i++ {
            for j := 1; j < len(octopuses[i])-1; j++ {
                if octopuses[i][j] > 9 { octopuses[i][j] = 0 }
            }
        }
    }
    return
}

func fleshes(octopuses [][]int, check map[pair]bool, i int, j int) (countFlesh int) {
    if !check[pair{i,j}] && i > 0 && i < len(octopuses)-1 && j > 0 && j < len(octopuses[i])-1 {
        octopuses[i][j]++
        if octopuses[i][j] > 9 {
            countFlesh++
            check[pair{i,j}] = true
            for k := -1; k < 2; k++ {
                for z := -1; z < 2; z++ {
                    if !(k == 0 && z == 0) { countFlesh += fleshes(octopuses, check, i+k, j+z) }
                }
            }
        }
    }
    return countFlesh
}

func print(octopuses [][]int) {
    for i := 1; i < len(octopuses)-1; i++ {
        for j := 1; j < len(octopuses[i])-2; j++ {
            fmt.Print(octopuses[i][j], ",")
        }
        fmt.Println(octopuses[i][len(octopuses[i])-2])
    }
}

func main() {
    fmt.Println(countFleshes(read()))
}
