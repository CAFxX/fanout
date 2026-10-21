package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 255
    v2 := make(map[string]int)
    var v3 float64 = -2.25
    var v4 uint8 = 127
    v5 := make(map[int]bool)
    v6 := make(map[string]string)
    delete(v2, ("a" + "a"))
    var v7 int8 = 100
    var v8 int16 = 0
    v9 := []int{0}
    delete(v5, 42)
    switch ("hello" + "a") {
    case "hello":
        fmt.Println(v4, v7)
    }
    fmt.Println(v7, v4)
    if (true || (!((!false) && (!true)))) {
        var v10 uint = 851
        _ = len(v9)
        _ = v10
    } else {
        v7 = v7
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
