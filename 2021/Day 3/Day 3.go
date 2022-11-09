package main

import "os"
import "log"
import "fmt"
import "math"
import "bufio"

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (input []string) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        input = append(input, scanner.Text())
    }
    return
}

func decimalToBinary(binaryNumber string) (decimal int) {
    for i, bit := range binaryNumber {
        if bit == '1' {
            decimal += int(math.Pow(float64(2), float64(len(binaryNumber)-1-i)))
        }
    }
    return
}

func bitZeroTrendPosition(binaryNumbers []string, i int) (zeroBit int) {
    for _, binaryNumber := range binaryNumbers {
        if binaryNumber[i] == '0' {
            zeroBit++
        }
    }
    return
}

func calculatePowerConsumption(binaryNumbers []string) (powerConsumption int) {
    var zeroBit []int = make([]int, len(binaryNumbers[0]))
    for i := 0; i < len(zeroBit); i++ {
        zeroBit[i] = bitZeroTrendPosition(binaryNumbers, i)
    }

    var gamma, epsilon string
    for i := 0; i < len(zeroBit); i++ {
        if zeroBit[i] > len(binaryNumbers)/2 {
            gamma += "0"
            epsilon += "1"
        } else {
            gamma += "1"
            epsilon += "0"
        }
    }
    powerConsumption = decimalToBinary(gamma) * decimalToBinary(epsilon)
    return
}

func filter(binaryNumbers []string, bit byte, i int) (newBinaryNumbers []string) {
    for _, binaryNumber := range binaryNumbers {
        if binaryNumber[i] == bit {
            newBinaryNumbers = append(newBinaryNumbers, binaryNumber)
        }
    }
    return newBinaryNumbers
}

func calculateLifeSupportRating(binaryNumbers []string) (lifeSupportRating int) {
    return calculateOxygenGeneratorRating(binaryNumbers, 0) * calculateCO2ScrubberRating(binaryNumbers, 0)
}

func calculateOxygenGeneratorRating(binaryNumbers []string, i int) (oxygenGeneratorRating int) {
    if len(binaryNumbers) == 1 {
        return decimalToBinary(binaryNumbers[0])
    }
    var zeroI = bitZeroTrendPosition(binaryNumbers, i)
    if zeroI > len(binaryNumbers)/2 {
        return calculateOxygenGeneratorRating(filter(binaryNumbers, '0', i), i+1)
    } else {
        return calculateOxygenGeneratorRating(filter(binaryNumbers, '1', i), i+1)
    }
}

func calculateCO2ScrubberRating(binaryNumbers []string, i int) (CO2ScrubberRating int) {
    if len(binaryNumbers) == 1 {
        return decimalToBinary(binaryNumbers[0])
    }
    var zeroI = bitZeroTrendPosition(binaryNumbers, i)
    if zeroI <= len(binaryNumbers)/2 {
        return calculateCO2ScrubberRating(filter(binaryNumbers, '0', i), i+1)
    } else {
        return calculateCO2ScrubberRating(filter(binaryNumbers, '1', i), i+1)
    }
}

func main() {
    var binaryNumbers []string = read()
    fmt.Println("Power consumption is:", calculatePowerConsumption(binaryNumbers))
    fmt.Println("The life support rating of the submarine is:", calculateLifeSupportRating(binaryNumbers))
}
