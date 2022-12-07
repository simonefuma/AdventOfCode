package main

import "os"
import "fmt"
import "log"
import "bufio"
import "math"

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (message string) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)
    
    if scanner.Scan() { message = fromHexToBinary(scanner.Text()) }
    return
}

func fromHexToBinary(hex string) (binary string) {
    var table map[rune]string = map[rune]string{'0': "0000",
                                                '1': "0001",
                                                '2': "0010",
                                                '3': "0011",
                                                '4': "0100",
                                                '5': "0101",
                                                '6': "0110",
                                                '7': "0111",
                                                '8': "1000",
                                                '9': "1001",
                                                'A': "1010",
                                                'B': "1011",
                                                'C': "1100",
                                                'D': "1101",
                                                'E': "1110",
                                                'F': "1111"}
    for _, r := range hex {
        binary += table[r]
    }
    return
}

func fromBinaryToDecimal(binary string) (decimal int) {
    for i := 0; i < len(binary); i++ {
        if binary[len(binary)-1-i] == '1' { decimal += int(math.Pow(2, float64(i))) }
    }
    return
}

func readMessage(binary string) (version int, value int) {
    version, _, value = readPackage(binary)
    return
}

func readPackage(binary string) (version int, bitRead int, valueLiteral int) {
    version += fromBinaryToDecimal(binary[0:3])
    id := fromBinaryToDecimal(binary[3:6])
    var results []int

    if id == 4 {
        // letterale
        var literal string
        var i int = 6
        for ; binary[i] != '0'; i += 5 {
            literal += binary[i+1:i+5]
            bitRead += 5
        }
        literal += binary[i+1:i+5]
        bitRead += 11
        valueLiteral = fromBinaryToDecimal(literal)
    } else {
        // operatore
        if binary[6] == '0' {
            // 15 lunghezza totale dei bit
            var length int = fromBinaryToDecimal(binary[7:22])
            bitRead += 22

            for ; length != 0 ; {
                v, t, l := readPackage(binary[bitRead:])
                version += v
                bitRead += t
                length -= t
                results = append(results, l)
            }
        } else {
            // 11 numero dei pacchetti
            var packet int = fromBinaryToDecimal(binary[7:18])
            bitRead += 18

            for i := 0; i < packet; i++ {
                v, t, l := readPackage(binary[bitRead:])
                version += v
                bitRead += t
                results = append(results, l)
            }
        }

        switch id {
        case 0:
            for _, element := range results {
                valueLiteral += element
            }
        case 1:
            valueLiteral = 1
            for _, element := range results {
                valueLiteral *= element
            }
        case 2:
            valueLiteral = results[0]
            for i := 1; i < len(results); i++ {
                if results[i] < valueLiteral { valueLiteral = results[i] }
            }
        case 3:
            valueLiteral = results[0]
            for i := 1; i < len(results); i++ {
                if results[i] > valueLiteral { valueLiteral = results[i] }
            }
        case 5:
            if results[0] > results[1] {
                valueLiteral = 1
            } else {
                valueLiteral = 0
            }
        case 6:
            if results[0] < results[1] {
                valueLiteral = 1
            } else {
                valueLiteral = 0
            }
        case 7:
            if results[0] == results[1] {
                valueLiteral = 1
            } else {
                valueLiteral = 0
            }
        }

    }
    return
}

func main() {
    var message string = read()
    version, value := readMessage(message)
    fmt.Println("sumVersion:", version)
    fmt.Println("value:", value)
}
