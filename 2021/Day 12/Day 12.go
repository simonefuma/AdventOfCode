package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strings"
import "unicode"

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (graph map[string][]string) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    graph = make(map[string][]string)
    for scanner.Scan() {
        var vertixs []string = strings.Split(scanner.Text(), "-")
        graph[vertixs[0]] = append(graph[vertixs[0]], vertixs[1])
        graph[vertixs[1]] = append(graph[vertixs[1]], vertixs[0])
    }
    return graph
}

func isLower(s string) bool {
    for _, r := range s {
        if !unicode.IsLower(r) { return false }
    }
    return true
}

func visit1(graph map[string][]string, vertix string, visited map[string]bool) (count int) {
    if vertix == "end" { return 1 }

    visited[vertix] = true
    for _, v := range graph[vertix] {
        if isLower(v) {
            if !visited[v] {
                count += visit1(graph, v, visited)
            }
        } else { count += visit1(graph, v, visited) }
    }
    visited[vertix] = false
    return
}

func visit2(graph map[string][]string, vertix string, visited map[string]int, minVisited bool) (count int) {
    if vertix == "end" { return 1 }

    visited[vertix]++
    for _, v := range graph[vertix] {
        if isLower(v) {
            if visited[v] == 0 {
                count += visit2(graph, v, visited, minVisited)
            } else if !minVisited && v != "start" {
                count += visit2(graph, v, visited, true)
            }
        } else { count += visit2(graph, v, visited, minVisited) }
    }

    visited[vertix]--
    return
}

func main() {
    fmt.Println(visit1(read(), "start", make(map[string]bool)))
    fmt.Println(visit2(read(), "start", make(map[string]int), false))
}
