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

func read() (template map[string]int, rules map[string]string, letter map[rune]int) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    template = make(map[string]int)
    rules = make(map[string]string)
    letter = make(map[rune]int)
    if scanner.Scan() {
        var text string = scanner.Text()
        for i := 0; i < len(text)-1; i++ {
            template[string(text[i:i+2])]++
            letter[rune(text[i])]++
        }
        letter[rune(text[len(text)-1])]++
    }
    if scanner.Scan() { scanner.Text() }

    for scanner.Scan() {
        var text string = scanner.Text()
        rules[string(text[0:2])] = string(text[6])
    }
    return
}

func run(template map[string]int, rules map[string]string, letter map[rune]int, step int) int {
    for i := 0; i < step; i++ {
        temp := make(map[string]int)
        for key, value := range template {
            temp[string(key[0])+rules[key]] += value
            temp[rules[key]+string(key[1])] += value
            temp[key] -= value
            letter[rune(rules[key][0])] += value
        }
        for key, value := range temp {
            template[key] += value
        }
    }
    var max, min int
    for _, value := range letter {
        min = value
        break
    }
    for _, value := range letter {
        if value > max { max = value }
        if value < min { min = value }
    }
    return max-min
}

func main() {
    template, rules, letter := read()
    fmt.Println(run(template, rules, letter, 40))
}
