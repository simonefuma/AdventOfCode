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

func read() (input [][]int){
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    input = append(input, []int{})
    for scanner.Scan() {
        var temp []int
        temp = append(temp, 9)
        for _, s := range scanner.Text() {
            n, err := strconv.Atoi(string(s))
            check(err)
            temp = append(temp, n)
        }
        temp = append(temp, 9)
        input = append(input, temp)
    }
    input = append(input, []int{})
    for i := 0; i < len(input[1]); i++ {
        input[0] = append(input[0], 9)
        input[len(input)-1] = append(input[len(input)-1], 9)
    }
    return
}

func riskLevel(cave [][]int) (risk int, dimension int) {
    var mostLarge [3]int
    for i := 1; i < len(cave)-1; i++ {
        for j := 1; j < len(cave[1])-1; j++ {
            if cave[i][j] < cave[i+1][j] && cave[i][j] < cave[i-1][j] && cave[i][j] < cave[i][j+1] && cave[i][j] < cave[i][j-1] {
                risk += cave[i][j]+1
                dimensionBasin := basinDimension(cave, make(map[pair]bool), i, j)
                if dimensionBasin > mostLarge[0] {
                    mostLarge[2] = mostLarge[1]
                    mostLarge[1] = mostLarge[0]
                    mostLarge[0] = dimensionBasin
                } else if dimensionBasin > mostLarge[1] {
                    mostLarge[2] = mostLarge[1]
                    mostLarge[1] = dimensionBasin
                } else if dimensionBasin > mostLarge[2] {
                    mostLarge[2] = dimensionBasin
                }
            }
        }
    }
    return risk, mostLarge[0] * mostLarge[1] * mostLarge[2]
}

func basinDimension(cave [][]int, checked map[pair]bool, i int, j int) (dimension int) {
    if cave[i][j] == 9 { return 0 }
    checked[pair{i,j}] = true
    if cave[i][j] < cave[i+1][j] && !checked[pair{i+1,j}] {dimension += basinDimension(cave, checked, i+1, j)}
    if cave[i][j] < cave[i-1][j] && !checked[pair{i-1,j}] {dimension += basinDimension(cave, checked, i-1, j)}
    if cave[i][j] < cave[i][j+1] && !checked[pair{i,j+1}] {dimension += basinDimension(cave, checked, i, j+1)}
    if cave[i][j] < cave[i][j-1] && !checked[pair{i,j-1}] {dimension += basinDimension(cave, checked, i, j-1)}
    return 1+dimension
}


func main() {
    fmt.Println(riskLevel(read()))
}
