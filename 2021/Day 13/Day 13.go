package  main

import "os"
import "fmt"
import "log"
import "bufio"
import "strings"
import "strconv"

type pair struct {
    x, y int
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (pairs map[pair]bool, commands []string, maxX, maxY int) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    pairs = make(map[pair]bool)
    for scanner.Scan() {
        if scanner.Text() == "" { break }
        var newPair pair = pair{0,0}
        var poinsS []string = strings.Split(scanner.Text(), ",")

        n, err := strconv.Atoi(poinsS[0])
        check(err)
        newPair.x = n
        if n > maxX { maxX = n }

        n, err = strconv.Atoi(poinsS[1])
        check(err)
        newPair.y = n
        if n > maxY { maxY = n }

        pairs[newPair] = true
    }

    for scanner.Scan() {
        commands = append(commands, scanner.Text())
    }

    return pairs, commands, maxX, maxY
}

func fold(pairs map[pair]bool, commands []string, maxX, maxY int) (int, int) {
    for _, command := range commands {
        n, err := strconv.Atoi(string(command[13:]))
        check(err)
        if command[11] == 'y' {
            foldY(pairs, n, maxX, maxY)
            maxY = n-1
        } else {
            foldX(pairs, n, maxX, maxY)
            maxX = n-1
        }
    }
    return maxX, maxY
}

func foldY(pairs map[pair]bool, command int, maxX, maxY int) {
    for y := 0; y < command; y++ {
        for x := 0; x <= maxX; x++ {
            newPair := pair{x, maxY-y}
            if pairs[pair{x,y}] || pairs[newPair] {
                pairs[pair{x,y}] = true
                delete(pairs, newPair)
            }
        }
    }
}

func foldX(pairs map[pair]bool, command int, maxX, maxY int) {
    for y := 0; y <= maxY; y++ {
        for x := 0; x < command; x++ {
            newPair := pair{maxX-x, y}
            if pairs[pair{x,y}] || pairs[newPair] {
                pairs[pair{x,y}] = true
                delete(pairs, newPair)
            }
        }
    }
}

func print(pairs map[pair]bool, maxX, maxY int) {
    for y := 0; y <= maxY; y++ {
        for x := 0; x <= maxX; x++ {
            if pairs[pair{x,y}] {
                fmt.Print("#")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
}

func main() {
    pairs, commands, maxX, maxY := read()
    maxX, maxY = fold(pairs, commands, maxX, maxY)
    print(pairs, maxX, maxY)
}
