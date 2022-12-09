package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strconv"

type Command struct {
    direction string
    move int
}

type Pair struct {
    x,y int
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (commands []Command) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        text := scanner.Text()
        n, err := strconv.Atoi(text[2:])
        check(err)
        commands = append(commands, Command{string(text[0]), n})
    }
    return
}

func run(commands []Command) map[Pair]bool {
    visited := make(map[Pair]bool)
    var cord []Pair
    for i := 0; i < 10; i++ {
        cord = append(cord, Pair{0,0})
    }
    visited[cord[9]] = true
    for _, command := range commands {
        for i := 0; i < command.move; i++ {
            switch command.direction {
            case "R":
                cord[0].x++
            case "L":
                cord[0].x--
            case "U":
                cord[0].y++
            case "D":
                cord[0].y--
            }
            for j := 1; j < 10; j++ {
                if cord[j-1].x == cord[j].x { // same row
                    distance := cord[j-1].y - cord[j].y
                    if distance <= 1 && distance >= -1 { break }
                    if distance < 0 {
                        cord[j].y--
                    } else {
                        cord[j].y++
                    }
                } else if cord[j-1].y == cord[j].y { // same column
                    distance := cord[j-1].x - cord[j].x
                    if distance <= 1 && distance >= -1 { break }
                    if distance < 0 {
                        cord[j].x--
                    } else {
                        cord[j].x++
                    }
                } else {
                    if (cord[j-1].x == cord[j].x+1 && cord[j-1].y == cord[j].y+1) || (cord[j-1].x == cord[j].x-1 && cord[j-1].y == cord[j].y-1) || (cord[j-1].x == cord[j].x+1 && cord[j-1].y == cord[j].y-1) || (cord[j-1].x == cord[j].x-1 && cord[j-1].y == cord[j].y+1) { break }
                    if (cord[j-1].x == cord[j].x+1 && cord[j-1].y == cord[j].y+2) || (cord[j-1].x == cord[j].x+2 && cord[j-1].y == cord[j].y+1) || (cord[j-1].x == cord[j].x+2 && cord[j-1].y == cord[j].y+2){
                        cord[j].x++
                        cord[j].y++
                    } else if (cord[j-1].x == cord[j].x-1 && cord[j-1].y == cord[j].y+2) || (cord[j-1].x == cord[j].x-2 && cord[j-1].y == cord[j].y+1) || (cord[j-1].x == cord[j].x-2 && cord[j-1].y == cord[j].y+2) {
                        cord[j].x--
                        cord[j].y++
                    } else if (cord[j-1].x == cord[j].x-2 && cord[j-1].y == cord[j].y-1) || (cord[j-1].x == cord[j].x-1 && cord[j-1].y == cord[j].y-2) || (cord[j-1].x == cord[j].x-2 && cord[j-1].y == cord[j].y-2) {
                        cord[j].x--
                        cord[j].y--
                    } else {
                        cord[j].x++
                        cord[j].y--
                    }
                }
            }
            visited[cord[9]] = true
        }
    }
    return visited
}

func main() {
    visited := run(read())
    fmt.Println(len(visited))
}
