package main

import "os"
import "fmt"
import "log"
import "bufio"
import "unicode"

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (bags []string) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        bags = append(bags, scanner.Text())
    }

    return
}

func countPriority(bags []string) (priority int) {
    for _, bag := range bags {
        nextBag:
        for _, objectFirst := range bag[:len(bag)/2] {
            for _, objectSecond := range bag[len(bag)/2:] {
                if objectFirst == objectSecond {
                    if unicode.IsLower(objectFirst) {
                        priority += int(objectFirst)-96
                    } else {
                        priority += int(objectFirst)-38
                    }
                    break nextBag
                }
            }
        }
    }
    return
}

func countPriorityGroup(bags []string) (priority int) {
    for i := 0; i < len(bags); i += 3 {
        nextBag:
        for _, elf1 := range bags[i] {
            for _, elf2 := range bags[i+1] {
                for _, ellf3 := range bags[i+2] {
                    if elf1 == elf2 && elf2 == ellf3 {
                        if unicode.IsLower(elf1) {
                            priority += (int(elf1)-96)
                        } else {
                            priority += (int(elf1)-38)
                        }
                        break nextBag
                    }
                }
            }
        }
    }
    return
}

func main() {
    var bags []string = read()
    fmt.Println(countPriority(bags))
    fmt.Println(countPriorityGroup(bags))
}
