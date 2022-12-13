package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strconv"

type Package struct {
    values []int
    packages []Package
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (packages []string) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        text := scanner.Text()
        if text != "" { packages = append(packages, text) }
    }
    return
}

func createPackage(text []string) (packs []Package) {
    for _, line := range text {
        newPackage, _ := createPackageRic(line)
        packs = append(packs, newPackage)
    }
    return
}

func createPackageRic(s string) (pack Package, i int) {
    for i = 1; i < len(s); i++ {
        if s[i] == '[' {
            newPackage, newPosition := createPackageRic(s[i:])
            pack.packages = append(pack.packages, newPackage)
            pack.values = append(pack.values, -1)
            i += newPosition
        } else if s[i] == ']' {
            return pack, i
        } else if s[i] != ',' {
            if s[i] == '1' && s[i+1] == '0' {
                n, err := strconv.Atoi(s[i:i+2])
                check(err)
                pack.values = append(pack.values, n)
                i++
            } else {
                n, err := strconv.Atoi(string(s[i]))
                check(err)
                pack.values = append(pack.values, n)
            }
        }
    }
    return pack, i
}

func order(packs []Package) (sum int) {
    for i := 0; i < len(packs); i += 2 {
        isRight, _ := orderRic(packs[i], packs[i+1])
        if isRight { sum += int(i/2)+1 }
    }
    return sum
}

func orderRic(packLeft, packRight Package) (isRight, isOk bool) {
    var pointerValues int
    var pointerPackagesLeft, pointerPackagesRight int
    for pointerValues < len(packLeft.values) && pointerValues < len(packRight.values) {
        if packLeft.values[pointerValues] != -1 && packRight.values[pointerValues] != -1 {
            if packLeft.values[pointerValues] < packRight.values[pointerValues] {
                return true, false
            } else if packLeft.values[pointerValues] > packRight.values[pointerValues] {
                return false, false
            }
        } else if packLeft.values[pointerValues] == -1 && packRight.values[pointerValues] == -1 {
            isRight, isOk = orderRic(packLeft.packages[pointerPackagesLeft], packRight.packages[pointerPackagesRight])
            if isRight {
                return true, false
            } else if !isOk { return false, false }
            pointerPackagesLeft++
            pointerPackagesRight++
        } else {
            if packLeft.values[pointerValues] != -1 {
                isRight, isOk := orderRic(Package{[]int{packLeft.values[pointerValues]}, []Package{}}, packRight.packages[pointerPackagesRight])
                if isRight {
                    return true, false
                } else if !isOk { return false, false }
                pointerPackagesRight++
            } else {
                isRight, isOk := orderRic(packLeft.packages[pointerPackagesLeft], Package{[]int{packRight.values[pointerValues]}, []Package{}})
                if isRight {
                    return true, false
                } else if !isOk { return false, false }
                pointerPackagesLeft++
            }
        }
        pointerValues++
    }
    if len(packLeft.values) == len(packRight.values) { return false, true }
    if pointerValues >= len(packLeft.values) { return true, false }
    return false, false
}

func sort(packages []Package) {
    for i := 0; i < len(packages)-1; i++ {
        var m int = i
        for j := i+1; j < len(packages); j++ {
            isRight, _ := orderRic(packages[j], packages[m])
            if isRight { m = j }
        }
        packages[i], packages[m] = packages[m], packages[i]
    }
}

func main() {
    text := read()
    packs := createPackage(text)
    fmt.Println("sum is:", order(packs))
    var first, second Package = Package{[]int{-1}, []Package{Package{[]int{2}, []Package{}}}}, Package{[]int{-1}, []Package{Package{[]int{6}, []Package{}}}}

    packs = append(append(packs, first), second)
    sort(packs)
    var firstPosition int
    for i := 0; i < len(packs); i++ {
        _, equal := orderRic(first, packs[i])
        if equal {
            firstPosition = i+1
            break
        }
    }

    for i := firstPosition; i < len(packs); i++ {
        _, equal := orderRic(second, packs[i])
        if equal {
            fmt.Println(firstPosition*(i+1))
            return
        }
    }
}
