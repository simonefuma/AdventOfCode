package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strconv"
import "strings"

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (fish [9]int) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    scanner.Scan()
    for _, stringDay := range strings.Split(scanner.Text(), ",") {
        intDay, err := strconv.Atoi(stringDay)
        check(err)
        fish[intDay]++
    }
    return
}

func afterTime(fish [9]int, time int) [9]int {
    for i := 0; i < time; i++ {
        fish[0], fish[1], fish[2], fish[3], fish[4], fish[5], fish[6], fish[7], fish[8] = fish[1], fish[2], fish[3], fish[4], fish[5], fish[6], fish[7]+fish[0], fish[8], fish[0]
    }
    return fish
}

func countFish(fish [9]int) (total int) {
    for i := 0; i < len(fish); i++ {
        total += fish[i]
    }
    return
}

func main() {
    var days int = 256
    fmt.Println("After", days, "days there are", countFish(afterTime(read(), days)), "fish")
}
