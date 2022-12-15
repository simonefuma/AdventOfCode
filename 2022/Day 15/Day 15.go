package main

import "os"
import "fmt"
import "log"
import "math"
import "bufio"
import "strconv"
import "strings"

type Pair struct {
    x,y int
}

type Sensor struct {
    position Pair
    manhattanDistance int
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (sensors []Sensor){
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        text := strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(scanner.Text(), "Sensor at ", "", 1), " closest beacon is at ", "", 1), " ", "", 2), "=", "", 4), "x", "", 2), "y", "", 2)
        pairs := strings.Split(text, ":")
        sensorString := strings.Split(pairs[0], ",")
        x, _ := strconv.Atoi(sensorString[0])
        y, _ := strconv.Atoi(sensorString[1])
        var sensorPosition [2]int = [2]int{x,y}
        beaconString := strings.Split(pairs[1], ",")
        x, _ = strconv.Atoi(beaconString[0])
        y, _ = strconv.Atoi(beaconString[1])
        var beaconPosition [2]int = [2]int{x,y}
        sensors = append(sensors, Sensor{Pair{sensorPosition[0], sensorPosition[1]}, int(math.Abs(float64(sensorPosition[0]-beaconPosition[0]))+math.Abs(float64(sensorPosition[1]-beaconPosition[1])))})
    }
    return
}

func searchPosition(sensors []Sensor, min, max int) {
    for y := min; y <= max; y++ {
        for x := min; x <= max; x++ {
            var isPosition bool = true
            for _, sensor := range sensors {
                if y >= sensor.position.y - sensor.manhattanDistance && y <= sensor.position.y + sensor.manhattanDistance && x >= sensor.position.x - sensor.manhattanDistance && x <= sensor.position.x + sensor.manhattanDistance {
                    var newManhattanDistanceX int = sensor.manhattanDistance - int(math.Abs(float64(sensor.position.y - y)))
                    var newManhattanDistanceY int = sensor.manhattanDistance - int(math.Abs(float64(sensor.position.x - x)))
                    if x >= sensor.position.x - newManhattanDistanceX && x <= sensor.position.x + newManhattanDistanceX && y >= sensor.position.y - newManhattanDistanceY && y <= sensor.position.y + newManhattanDistanceY {
                        isPosition = false
                        x += sensor.position.x - x + newManhattanDistanceX
                        break
                    }
                }
            }
            if isPosition {
                fmt.Println(x*4000000+y)
                return
            }
            isPosition = true
        }
    }
}

func main() {
    var min, max int = 0, 4000000
    var sensors []Sensor = read()
    searchPosition(sensors, min, max)
}
