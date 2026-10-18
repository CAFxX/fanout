package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }

func main() {
    v1 := make(map[int]int)
    var v2 uint = 18446744073709551615
    var v3 string = ("\n\t" + "hello")
    var v4 uint16 = 255
    v5 := GenS1{F0: 4294967295, F1: (!true)}
    switch 643 {
    case 2147483647:
        var v6 uint16 = (v4 / ((v4) | 1))
        _ = v6
    default:
        v7 := GenS1{F0: 437, F1: (255 <= 255)}
        _ = v7
    }
    _ = len(v1)
    var v8 uint16 = v4
    var v9 int64 = 42
    var v10 uint8 = 2
    fmt.Println(v3, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
