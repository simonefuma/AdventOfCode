package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strings"
import "strconv"

type Node struct {
    value rune
    next *Node
}

type Stack struct {
    top *Node
}

func push(stack *Stack, value rune) {
    (*stack).top = &Node{value, (*stack).top}
}

func pushReverse(stack *Stack, value rune) {
    var follow *Node = nil
    for pointer := (*stack).top; pointer != nil; pointer = (*pointer).next { follow = pointer }
    if follow == nil {
        push(stack, value)
    } else {
        (*follow).next = &Node{value, nil}
    }
}

func pop(stack *Stack) rune {
    var value rune = (*(*stack).top).value
    (*stack).top = (*(*stack).top).next
    return value
}

func isEmpty(stack *Stack) bool {
    return (*stack).top == nil
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (stacks []*Stack, commands [][3]int) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        var text string = scanner.Text()
        if len(text) != 0 && text[1] != '1' {
            for i := 1; i < len(text); i += 4 {
                if text[i] == ' ' { continue }
                if i/4 >= len(stacks) {
                    var j int = len(stacks)
                    for ; j <= i/4; j++ {
                        stacks = append(stacks, &Stack{nil})
                    }
                    pushReverse(stacks[j-1], rune(text[i]))
                } else {
                    pushReverse(stacks[i/4], rune(text[i]))
                }
            }
        } else { break }
    }

    for scanner.Scan() {
        var text []string = strings.Split(scanner.Text(), " ")
        var command [3]int
        for i := 1; i < len(text); i += 2 {
            n, err := strconv.Atoi(string(text[i]))
            check(err)
            if i != 1 { n-- }
            command[i/2] = n
        }
        commands = append(commands, command)
    }
    return
}
func run(stacks []*Stack, commands [][3]int) {
    for _, command := range commands {
        var crates string
        for i := 0; i < command[0]; i++ {
            crates += string(pop(stacks[command[1]]))
        }
        for i := len(crates)-1; i >= 0; i-- {
            push(stacks[command[2]], rune(crates[i]))
        }
    }

    for _, stack := range stacks {
        fmt.Print(string(pop(stack)))
    }
}

func main() {
    run(read())
    fmt.Println()
}
