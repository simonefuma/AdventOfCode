package main

import "os"
import "fmt"
import "log"
import "bufio"

type Node struct {
    value rune
    next *Node
}

type Stack struct {
    top *Node
}

func push(stack *Stack, value rune) {
    newNode := &Node{value, (*stack).top}
    (*stack).top = newNode
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

func insertSort(slice []int, value int) []int {
    slice = append(slice, value)
    for i := 0; i < len(slice)-1; i++ {
        if slice[i] > value {
            for j := len(slice)-2; j >= i ;j-- { slice[j+1] = slice[j] }
            slice[i] = value
            break
        }
    }
    return slice
}

func read() (pointCheckers int, pointAutocomplete int) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    var pointsAutocomplete []int
    var isValid bool = true
    for scanner.Scan() {
        var stack *Stack = &Stack{nil}
        for _, element := range scanner.Text() {
            switch element {
            case '(', '[', '{', '<':
                push(stack, element)
            default:
                var close rune = pop(stack)

                switch element {
                case ')':
                    if close != '(' {
                        pointCheckers += 3
                        isValid = false
                        break
                    }
                case ']':
                    if close != '[' {
                        pointCheckers += 57
                        isValid = false
                        break
                    }
                case '}':
                    if close != '{' {
                        pointCheckers += 1197
                        isValid = false
                        break
                    }
                case '>':
                    if close != '<' {
                        pointCheckers += 25137
                        isValid = false
                        break
                    }
                }
            }
        }
        if isValid {
            for ; !isEmpty(stack) ; {
                var close rune = pop(stack)
                pointAutocomplete *= 5
                switch close {
                case '(':
                    pointAutocomplete += 1
                case '[':
                    pointAutocomplete += 2
                case '{':
                    pointAutocomplete += 3
                case '<':
                    pointAutocomplete += 4
                }
            }
            pointsAutocomplete = insertSort(pointsAutocomplete, pointAutocomplete)
            pointAutocomplete = 0
        } else {
            isValid = true
        }
    }
    return pointCheckers, pointsAutocomplete[len(pointsAutocomplete)/2]
}

func main() {
    fmt.Println(read())
}
