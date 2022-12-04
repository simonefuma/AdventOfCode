package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strings"
import "strconv"

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (countContain int, countOverlays int) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        var elps []string = strings.Split(scanner.Text(), ",")
        var elp1S []string = strings.Split(elps[0], "-")
        var elp2S []string = strings.Split(elps[1], "-")
        var elp1i []int = getIntElp(elp1S)
        var elp2i []int = getIntElp(elp2S)
        if contain(elp1i, elp2i) || contain(elp2i, elp1i) { countContain++ }
        if overlays(elp1i, elp2i) || overlays(elp2i, elp1i) { countOverlays++ }
    }
    return
}

func getIntElp(elpS []string) (elp []int) {
    for _, s := range elpS {
        i, err := strconv.Atoi(s)
        check(err)
        elp = append(elp, i)
    }
    return
}

func overlays(elp1 []int, elp2 []int) bool {
    if elp1[0] <= elp2[1] && (elp1[1] >= elp2[1] || elp1[1] >= elp2[0]) { return true }
    return false
}

func contain(elp1 []int, elp2 []int) bool {
    if elp1[0] <= elp2[0] && elp1[1] >= elp2[1] { return true }
    return false
}

func main() {
    fmt.Println(read())
}
