package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 22
    var v3 bool = true
    var v4 float32 = ((3.14159 - 1e10) - 1e10)
    v3 = v3
    v3 = (!true)
    fmt.Println(v4, v3)
    v5 := GenS1{F0: 4294967295, F1: 2, F2: 170}
    v2 -= 21
    v6 := []int{}
    if len(v6) > 0 {
        v6[0] = 520
    }
    switch "\n\t" {
    case "\n\t":
        var v7 uint16 = 0
        _ = v7
    case "a":
        var v8 uint32 = 2
        _ = v8
    case "":
        fmt.Println(v3, v3, v4)
    }
    fmt.Println(v4, v4)
    var v9 uint = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
