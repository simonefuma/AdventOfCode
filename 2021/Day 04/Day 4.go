package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strconv"
import "strings"

type Board struct {
    board [][]int
    sumUnmarkedNumbers int
    win bool
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() (extracts []int, boards []Board) {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        if scanner.Text() != "" {
            for _, extractString := range strings.Split(scanner.Text(), ",") {
                extractInt, err := strconv.Atoi(extractString)
                check(err)
                extracts = append(extracts, extractInt)
            }
        } else { break }
    }

    var board [][]int = make([][]int, 5)
    var sumUnmarkedNumbers int
    var i int
    for scanner.Scan() {
        if scanner.Text() != "" {
            for _, numberString := range strings.Split(scanner.Text(), " ") {
                if numberString != "" {
                    numberInt, err := strconv.Atoi(numberString)
                    check(err)
                    board[i] = append(board[i], numberInt)
                    sumUnmarkedNumbers += numberInt
                }
            }
            i++
        } else {
            i = 0
            boards = append(boards, Board{board, sumUnmarkedNumbers, false})
            board = make([][]int, 5)
            sumUnmarkedNumbers = 0
        }
    }
    boards = append(boards, Board{board, sumUnmarkedNumbers, false})

    return extracts, boards
}

func scoreFirstLast(extracts []int, boards[]Board) (scoreFirst, scoreLast int) {
    var numberWinners int
    for j, extract := range extracts {
        for i := 0; i < len(boards); i++ {
            if boards[i].win { continue }
            var validColumn []bool = []bool{true, true, true, true, true}
        nextTable:
            for y := 0; y < 5; y++ {
                var validRow bool = true
                for x := 0; x < 5; x++ {
                    /*
                     * if the extracted number is equal to the one taken into consideration,
                     * then the sum of the table is decreased by that value
                     * and the position is circled (-1 is put as value in that position)
                     */
                    if number := boards[i].board[y][x]; number == extract {
                        boards[i].sumUnmarkedNumbers -= number
                        boards[i].board[y][x] = -1
                        /*
                         * if the element taken into consideration is neither the one extracted nor one already circled,
                         * then the row and column of that element cannot have been completed
                         */
                        if j > 3 {
                            /*
                             * the only row/column that can be completed is that of the new circled value
                             */
                            //------------------------------------------------------
                            /*
                             * if the leftmost values ​​of the element row have been circled, then validRow will be true,
                             * check if the leftmost values ​​are true too if they are then a row has been completed
                             */
                            if validRow {
                                for newX := x+1; newX < 5; newX++ {
                                    if boards[i].board[y][newX] != -1 {
                                        validRow = false
                                        break
                                     }
                                }
                                if validRow {
                                    if numberWinners == 0 {
                                        scoreFirst = boards[i].sumUnmarkedNumbers * number
                                    } else if numberWinners == len(boards)-1 {
                                        return scoreFirst, boards[i].sumUnmarkedNumbers * number
                                    }
                                    numberWinners++
                                    boards[i].win = true
                                    break nextTable
                                }
                            }
                            //------------------------------------------------------
                            /*
                            * if the values ​​above on the element column have been circled, then validColumn [x] will be true,
                            * check if the ones below are also true if they are then a column has been completed
                            */
                            if validColumn[x] {
                                for newy := y+1; newy < 5; newy++ {
                                    if boards[i].board[newy][x] != -1 {
                                        validColumn[x] = false
                                        break
                                     }
                                }
                                if validColumn[x] {
                                    if numberWinners == 0 {
                                        scoreFirst = boards[i].sumUnmarkedNumbers * number
                                    } else if numberWinners == len(boards)-1 {
                                        return scoreFirst, boards[i].sumUnmarkedNumbers * number
                                    }
                                    numberWinners++
                                    boards[i].win = true
                                    break nextTable
                                }
                            }
                        }
                        //------------------------------------------------------
                        break nextTable
                        //------------------------------------------------------
                    /*
                     * if the element taken into consideration is neither the one extracted nor one already circled,
                     * then the row and column of that element cannot have been completed
                     */
                    } else if number != -1 {
                        validRow = false
                        validColumn[x] = false
                    }
                }
            }
        }
    }
    return -1, -1
}

func main() {
    var firstBoard, lastBoard int = scoreFirstLast(read())
    fmt.Println("Score of the first board:", firstBoard)
    fmt.Println("Score of the last board:", lastBoard)
}
