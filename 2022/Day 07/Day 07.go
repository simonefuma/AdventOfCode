package main

import "os"
import "fmt"
import "log"
import "bufio"
import "strconv"
import "strings"

type Directory struct {
    id int
    name string
}

type FileSystem struct {
    father map[Directory]Directory
    childrenDir map[Directory][]Directory
    sizeDir map[Directory]int
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func read() *FileSystem {
    reader, err := os.Open("input.txt")
    defer reader.Close()
    check(err)
    scanner := bufio.NewScanner(reader)

    if scanner.Scan() { scanner.Text() }
    var fileSystem *FileSystem = &FileSystem{make(map[Directory]Directory), make(map[Directory][]Directory), make(map[Directory]int)}
    var directoryPosition Directory = Directory{1, "/"}
    var id int = 2
    for scanner.Scan() {
        var line string = scanner.Text()
        if line[0] == '$' {
            var command []string = strings.Split(line[2:], " ")
            if command[0] == "cd" {
                if command[1] != ".." {
                    for _, direcotry := range (*fileSystem).childrenDir[directoryPosition] {
                        if direcotry.name == command[1] {
                            directoryPosition = direcotry
                            break
                        }
                    }
                } else {
                    direcotryFather := (*fileSystem).father[directoryPosition]
                    (*fileSystem).sizeDir[direcotryFather] += (*fileSystem).sizeDir[directoryPosition]
                    directoryPosition = direcotryFather
                }
            }
        } else {
            elementsLine := strings.Split(line, " ")
            if elementsLine[0] == "dir" {
                var newDirectory Directory = Directory{id, elementsLine[1]}
                (*fileSystem).father[newDirectory] = directoryPosition
                (*fileSystem).childrenDir[directoryPosition] = append((*fileSystem).childrenDir[directoryPosition], newDirectory)
                id++
            } else {
                n, err := strconv.Atoi(elementsLine[0])
                check(err)
                (*fileSystem).sizeDir[directoryPosition] += n
            }
        }
    }
    for p := directoryPosition; (*fileSystem).father[p].name != ""; p = (*fileSystem).father[p] {
        (*fileSystem).sizeDir[(*fileSystem).father[p]] += (*fileSystem).sizeDir[p]
    }
    return fileSystem
}

func sumValidDirectory(fileSystem *FileSystem) (sum int) {
    for _, size := range (*fileSystem).sizeDir {
        if size <= 100000 { sum += size }
    }
    return
}

func chooseDirectory(fileSystem *FileSystem) int {
    var totalSpace int = 70000000
    var remainingSpace int = totalSpace - (*fileSystem).sizeDir[Directory{1, "/"}]
    var necessarySpace int = 30000000 - remainingSpace
    var min int = totalSpace

    for _, space := range (*fileSystem).sizeDir {
        if space < min && space > necessarySpace { min = space }
    }
    return min
}

func main() {
    var fileSystem *FileSystem = read()

    fmt.Println("sum memory used in valid directory:", sumValidDirectory(fileSystem))
    fmt.Println("size of the directory to be deleted", chooseDirectory(fileSystem))
}
