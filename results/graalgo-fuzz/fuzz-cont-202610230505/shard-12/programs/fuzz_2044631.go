package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 bool }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = float64((-2.25 - 0.0))
    var v3 uint = 558
    v4 := make(map[uint32]bool)
    var v5 float32 = (((0.1 - float32((1.5 / (-2.25 - v2)))) * 0.0) + 3.14159)
    var v6 int16 = -1
    if ((!false) || false) {
        fmt.Println(v5)
        v7 := GenS1{F0: -1, F1: 1, F2: true}
        _ = v7
    } else {
        fmt.Println(v5)
        var v8 string = "a"
        _ = v8
    }
    v9 := GenS2{F0: 988}
    v10 := []int{545, -461}
    var v11 uint64 = 831
    v2 -= float64((v5 + 1e10))
    if len(v10) > 0 {
        v10[0] = 168
    }
    var v12 uint32 = 17
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v13 uint8 = 39
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
