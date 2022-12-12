package main

import "os"
import "fmt"
import "log"
import "bufio"

type square struct {
    name rune
    x, y int
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (mapping [][]rune) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        var line []rune
        for _, r := range scanner.Text() {
            line = append(line, r)
        }
        mapping = append(mapping, line)
    }
    return
}

func createGraph(mapping [][]rune) (graph map[square][]square) {
    graph = make(map[square][]square)
    for y := 0; y < len(mapping); y++ {
        for x := 0; x < len(mapping[0]); x++ {
            var posVal rune = mapping[y][x]
            if posVal == 'S' { posVal = 'a' }
            if x == 0 {
                if y == 0 {
                    positionValue := mapping[y+1][x]
                    if positionValue == 'E' { positionValue = 'z' }
                    if posVal >= positionValue-1 { graph[square{mapping[y][x], x,y}] = append(graph[square{mapping[y][x], x,y}], square{mapping[y+1][x], x,y+1}) }
                } else if y == len(mapping)-1 {
                    positionValue := mapping[y-1][x]
                    if positionValue == 'E' { positionValue = 'z' }
                    if posVal >= positionValue-1 { graph[square{mapping[y][x], x,y}] = append(graph[square{mapping[y][x], x,y}], square{mapping[y-1][x], x,y-1}) }
                } else {
                    positionValue := mapping[y-1][x]
                    if positionValue == 'E' { positionValue = 'z' }
                    if posVal >= positionValue { graph[square{mapping[y][x], x,y}] = append(graph[square{mapping[y][x], x,y}], square{mapping[y-1][x], x,y-1}) }
                    positionValue = mapping[y+1][x]-1
                    if positionValue == 'E' { positionValue = 'z' }
                    if posVal >= positionValue-1 { graph[square{mapping[y][x], x,y}] = append(graph[square{mapping[y][x], x,y}], square{mapping[y+1][x], x,y+1}) }
                }
                positionValue := mapping[y][x+1]
                if positionValue == 'E' { positionValue = 'z' }
                if posVal >= positionValue-1 { graph[square{mapping[y][x], x,y}] = append(graph[square{mapping[y][x], x,y}], square{mapping[y][x+1], x+1,y}) }
            } else if x == len(mapping[0])-1 {
                if y == 0 {
                    positionValue := mapping[y+1][x]
                    if positionValue == 'E' { positionValue = 'z' }
                    if posVal >= positionValue-1 { graph[square{mapping[y][x], x,y}] = append(graph[square{mapping[y][x], x,y}], square{mapping[y+1][x], x,y+1}) }
                } else if y == len(mapping)-1 {
                    positionValue := mapping[y-1][x]
                    if positionValue == 'E' { positionValue = 'z' }
                    if posVal >= positionValue-1 { graph[square{mapping[y][x], x,y}] = append(graph[square{mapping[y][x], x,y}], square{mapping[y-1][x], x,y-1}) }
                } else {
                    positionValue := mapping[y+1][x]
                    if positionValue == 'E' { positionValue = 'z' }
                    if posVal >= positionValue-1 { graph[square{mapping[y][x], x,y}] = append(graph[square{mapping[y][x], x,y}], square{mapping[y+1][x], x,y+1}) }
                    positionValue = mapping[y-1][x]
                    if positionValue == 'E' { positionValue = 'z' }
                    if posVal >= positionValue-1 { graph[square{mapping[y][x], x,y}] = append(graph[square{mapping[y][x], x,y}], square{mapping[y-1][x], x,y-1}) }
                }
                positionValue := mapping[y][x-1]
                if positionValue == 'E' { positionValue = 'z' }
                if posVal >= positionValue-1 { graph[square{mapping[y][x], x,y}] = append(graph[square{mapping[y][x], x,y}], square{mapping[y][x-1], x-1,y}) }
            } else {
                if y == 0 {
                    positionValue := mapping[y+1][x]
                    if positionValue == 'E' { positionValue = 'z' }
                    if posVal >= positionValue-1 { graph[square{mapping[y][x], x,y}] = append(graph[square{mapping[y][x], x,y}], square{mapping[y+1][x], x,y+1}) }
                } else if y == len(mapping)-1 {
                    positionValue := mapping[y-1][x]
                    if positionValue == 'E' { positionValue = 'z' }
                    if posVal >= positionValue-1 { graph[square{mapping[y][x], x,y}] = append(graph[square{mapping[y][x], x,y}], square{mapping[y-1][x], x,y-1}) }
                } else {
                    positionValue := mapping[y+1][x]
                    if positionValue == 'E' { positionValue = 'z' }
                    if posVal >= positionValue-1 { graph[square{mapping[y][x], x,y}] = append(graph[square{mapping[y][x], x,y}], square{mapping[y+1][x], x,y+1}) }
                    positionValue = mapping[y-1][x]
                    if positionValue == 'E' { positionValue = 'z' }
                    if posVal >= positionValue-1 { graph[square{mapping[y][x], x,y}] = append(graph[square{mapping[y][x], x,y}], square{mapping[y-1][x], x,y-1}) }
                }
                positionValue := mapping[y][x-1]
                if positionValue == 'E' { positionValue = 'z' }
                if posVal >= positionValue-1 { graph[square{mapping[y][x], x,y}] = append(graph[square{mapping[y][x], x,y}], square{mapping[y][x-1], x-1,y}) }
                positionValue = mapping[y][x+1]
                if positionValue == 'E' { positionValue = 'z' }
                if posVal >= positionValue-1 { graph[square{mapping[y][x], x,y}] = append(graph[square{mapping[y][x], x,y}], square{mapping[y][x+1], x+1,y}) }
            }
        }
    }
    return
}

func visit(graph map[square][]square, start square) int {
    var count int
    var queue []square = []square{start}
    var visited map[square]bool = make(map[square]bool)
    visited[start] = true
    for ; len(queue) > 0 ; {
        var l int = len(queue)
        for i := 0; i < l; i++ {
            if queue[0].name == 'E' { return count }
            for _, element := range graph[queue[0]] {
                if !visited[element] {
                    queue = append(queue, element)
                    visited[element] = true
                }
            }
            queue = queue[1:]
        }
        count++
    }
    return -1
}

func print(t square) {
    fmt.Print(string(t.name),", ", t.x,", ", t.y)
}


func main() {
    graph := createGraph(read())
    var min int
    for key, _ := range graph {
        if key.name == 'S' {
            min = visit(graph, key)
            break
        }
    }
    for key, _ := range graph {
        if key.name == 'a' {
            var v int = visit(graph, key)
            if v < min && v != -1{ min = v }
        }
    }
    fmt.Println(min)
}
