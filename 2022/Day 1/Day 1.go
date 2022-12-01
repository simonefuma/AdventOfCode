package main

import "os"
import "log"
import "fmt"
import "bufio"
import "strconv"

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() int {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    var maxCalories [3]int
    var calories int
    for scanner.Scan() {
        var line string = scanner.Text()
        if line != "" {
            c, err := strconv.Atoi(line)
            check(err)
            calories += c
        } else {
            if calories > maxCalories[0] {
                maxCalories[2] = maxCalories[1]
                maxCalories[1] = maxCalories[0]
                maxCalories[0] = calories
            } else if calories > maxCalories[1] {
                maxCalories[2] = maxCalories[1]
                maxCalories[1] = calories
            } else if calories > maxCalories[2] {
                maxCalories[2] = calories
            }
            calories = 0
        }
    }
    return maxCalories[0]+maxCalories[1]+maxCalories[2]
}

func main() {
    fmt.Println(read())
}
