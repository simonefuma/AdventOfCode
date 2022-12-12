package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strings"
import "strconv"

type Number struct {
    divisible map[int]int
}

type Operation struct {
    op string
    value int
}

type Monkey struct {
    worry []Number
    operation Operation
    divider int
    selection [2]int
    inspects int
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (text []string) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        text = append(text, scanner.Text())
    }
    return
}

func createMonkeys(text []string) (monkeys []Monkey) {
    var dividers []int
    for i := 3; i < len(text); i += 7 {
        n, err := strconv.Atoi(text[i][21:])
        check(err)
        dividers = append(dividers, n)
    }

    for i := 1; i < len(text); i += 7 {
        var monkey Monkey = Monkey{}

        for _, element := range strings.Split(text[i][18:], ", ") {
            n, err := strconv.Atoi(element)
            check(err)
            divisible := make(map[int]int)
            for _, divider := range dividers {
                divisible[divider] = n%divider
            }
            monkey.worry = append(monkey.worry, Number{divisible})
        }

        n, err := strconv.Atoi(text[i+1][25:])
        monkey.operation = Operation{string(text[i+1][23]), n}

        monkey.divider = dividers[i/7]

        n, err = strconv.Atoi(text[i+3][29:])
        check(err)
        monkey.selection[0] = n

        n, err = strconv.Atoi(text[i+4][30:])
        check(err)
        monkey.selection[1] = n

        monkeys = append(monkeys, monkey)
    }
    return
}

func run(monkeys []Monkey, rounds int) {
    for round := 0; round < rounds; round++ {
        for monkey := 0; monkey < len(monkeys); monkey++ {
            for item := 0; item < len(monkeys[monkey].worry); item++ {
                switch monkeys[monkey].operation.op {
                case "+":
                    if monkeys[monkey].operation.value != 0 {
                        for key, _ := range monkeys[monkey].worry[item].divisible {
                            monkeys[monkey].worry[item].divisible[key] = (monkeys[monkey].worry[item].divisible[key] + monkeys[monkey].operation.value)%key
                        }
                    } else {
                        for key, _ := range monkeys[monkey].worry[item].divisible {
                            monkeys[monkey].worry[item].divisible[key] = (monkeys[monkey].worry[item].divisible[key] + monkeys[monkey].worry[item].divisible[key])%key
                        }
                    }
                case "*":
                    if monkeys[monkey].operation.value != 0 {
                        for key, _ := range monkeys[monkey].worry[item].divisible {
                            monkeys[monkey].worry[item].divisible[key] = (monkeys[monkey].worry[item].divisible[key] * monkeys[monkey].operation.value)%key
                        }
                    } else {
                        for key, _ := range monkeys[monkey].worry[item].divisible {
                            monkeys[monkey].worry[item].divisible[key] = (monkeys[monkey].worry[item].divisible[key] * monkeys[monkey].worry[item].divisible[key])%key
                        }
                    }
                }

                worry := monkeys[monkey].worry[0]
                if len(monkeys[monkey].worry) > 1 {
                    monkeys[monkey].worry = monkeys[monkey].worry[1:]
                    item--
                } else {
                    monkeys[monkey].worry = []Number{}
                }

                if worry.divisible[monkeys[monkey].divider] == 0 {
                    monkeys[monkeys[monkey].selection[0]].worry = append(monkeys[monkeys[monkey].selection[0]].worry, worry)
                } else {
                    monkeys[monkeys[monkey].selection[1]].worry = append(monkeys[monkeys[monkey].selection[1]].worry, worry)
                }

                monkeys[monkey].inspects++
            }
        }
    }
}

func main() {
    monkeys := createMonkeys(read())
    run(monkeys, 10000)

    max1, max2 := monkeys[0].inspects, monkeys[1].inspects
    if max1 < max2 { max1, max2 = max2, max1 }
    for i := 2; i < len(monkeys); i++ {
        if monkeys[i].inspects > max1 {
            max1, max2 = monkeys[i].inspects, max2
        } else if monkeys[i].inspects > max2 {
            max2 = monkeys[i].inspects
        }
    }
    fmt.Println(max1*max2)
}
