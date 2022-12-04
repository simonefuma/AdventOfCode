package main

import "os"
import "fmt"
import "log"
import "math"
import "bufio"
import "strings"

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (entrys []string) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        entrys = append(entrys, scanner.Text())
    }
    return
}

func sort(in []string) []string {
    var newIn []string = []string{9:""}
    var five, six int
    for i := 0; i < len(in); i++ {
        switch len(in[i]) {
        case 2:
            newIn[0] = in[i]
        case 3:
            newIn[1] = in[i]
        case 4:
            newIn[2] = in[i]
        case 7:
            newIn[3] = in[i]
        case 5:
            newIn[4+five] = in[i]
            five++
        case 6:
            newIn[7+six] = in[i]
            six++
        }
    }
    return newIn
}

func remove(s, rem string) (new string) {
    for i := 0; i < len(s); i++ {
        var thereIs bool
        for j := 0; j < len(rem); j++ {
            if s[i] == rem[j] {
                thereIs = true
                break
            }
        }
        if !thereIs {
            new += string(s[i])
        }
    }
    return new
}

func isIn(s, r string) bool {
    for _, lecter := range s {
        if string(lecter) == r { return true}
    }
    return false
}

func ContainsAny(str, char string) bool {
    var thereIs bool = false
    for _, s := range str {
        for _, c := range char {
            if s == c {
                thereIs = true
                break
            }
        }
        if !thereIs {
            return false
        }
        thereIs = false
    }
    return true
}

func checkSegments(digits []string, segments [7]string) (orderDigit []string, isValid bool) {
    orderDigit = []string{4:""}
    var isNumbers [5]bool
    for _, digit := range digits {
        if !isIn(digit, segments[1]) && isIn(digit, segments[2]) && isIn(digit, segments[3]) && isIn(digit, segments[4]) && !isIn(digit, segments[5]) {
            orderDigit[0] = digit
            isNumbers[0] = true
            continue
        }
        if !isIn(digit, segments[1]) && isIn(digit, segments[2]) && isIn(digit, segments[3]) && !isIn(digit, segments[4]) && isIn(digit, segments[5]) {
            orderDigit[1] = digit
            isNumbers[1] = true
            continue
        }
        if isIn(digit, segments[1]) && !isIn(digit, segments[2]) && isIn(digit, segments[3]) && !isIn(digit, segments[4]) && isIn(digit, segments[5]) {
            orderDigit[2] = digit
            isNumbers[2] = true
            continue
        }
        if isIn(digit, segments[1]) && !isIn(digit, segments[2]) && isIn(digit, segments[3]) && isIn(digit, segments[4]) && isIn(digit, segments[5]) {
            orderDigit[3] = digit
            isNumbers[3] = true
            continue
        }
        if isIn(digit, segments[1]) && isIn(digit, segments[2]) && isIn(digit, segments[3]) && !isIn(digit, segments[4]) && isIn(digit, segments[5]) {
            orderDigit[4] = digit
            isNumbers[4] = true
        }
    }

    for _, element := range isNumbers {
        if !element {
            return orderDigit, false
        }
    }
    return orderDigit, true
}

func analisys(in []string) []string {
    var newIn []string = sort(in)
    var segments [7]string
    var zero string = remove(newIn[1], newIn[0])
    var twoOrfive string = newIn[0]
    var oneOrThree string = remove(newIn[2], newIn[0])
    var fourOrSix string = remove(remove(newIn[3], newIn[2]), zero)

    segments[0] = zero
    for i := 0; i < len(twoOrfive); i++ {
        segments[2] = string(twoOrfive[i])
        segments[5] = string(twoOrfive[1-i])
        for j := 0; j < len(oneOrThree); j++ {
            segments[1] = string(oneOrThree[j])
            segments[3] = string(oneOrThree[1-j])
            for w := 0; w < len(fourOrSix); w++ {
                segments[4] = string(fourOrSix[w])
                segments[6] = string(fourOrSix[1-w])
                p, check := checkSegments(newIn[4:], segments)
                if check {
                    return p
                }
            }
        }
    }
    return []string{}
}

func sumOutput(entrys []string) (sum int) {
    for _, entry := range entrys {
        var partialSum int
        var splitEntry []string = strings.Split(entry, " | ")
        var in, out []string = strings.Split(splitEntry[0], " "), strings.Split(splitEntry[1], " ")

        var sorted []string = analisys(in)

        for i, number := range out {
            switch len(number) {
            case 2: // 1
                partialSum += int(math.Pow10(3-i))
            case 3: // 7
                partialSum += 7 * int(math.Pow10(3-i))
            case 4: // 4
                partialSum += 4 * int(math.Pow10(3-i))
            case 7: // 8
                partialSum += 8 * int(math.Pow10(3-i))
            case 5: // 2, 3, 5
                switch {
                case ContainsAny(sorted[0], number): //2
                    partialSum += 2 * int(math.Pow10(3-i))
                case ContainsAny(sorted[1], number): // 3
                    partialSum += 3 * int(math.Pow10(3-i))
                case ContainsAny(sorted[2], number): // 5
                    partialSum += 5 * int(math.Pow10(3-i))
                }
            case 6: // 0, 6, 9
                switch {
                case ContainsAny(sorted[3], number): // 6
                    partialSum += 6 * int(math.Pow10(3-i))
                case ContainsAny(sorted[4], number): // 9
                    partialSum += 9 * int(math.Pow10(3-i))
                }
            }
        }
        sum += partialSum
    }
    return sum
}


func main() {
    fmt.Println(sumOutput(read()))
}
