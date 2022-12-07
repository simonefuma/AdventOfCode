package main

import "os"
import "fmt"
import "log"
import "bufio"

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() string {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    if scanner.Scan() { return scanner.Text() }
    return ""
}

func run(s string) int {
    var control bool = true
    for i := 13; i < len(s); i++ {
        next:
        for j := i-13; j < i; j++ {
            for z := j+1; z <= i; z++ {
                if s[j] == s[z] {
                    control = false
                    break next
                }
            }
        }
        if control { return i+1 }
        control = true
    }
    return -1
}

func main() {
    fmt.Println(run(read()))
}
