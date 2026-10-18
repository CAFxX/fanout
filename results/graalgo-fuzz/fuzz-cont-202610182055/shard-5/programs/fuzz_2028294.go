package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint; F2 uint }

func main() {
    var v1 uint8 = 42
    var v2 float64 = (0.0 / 1.0)
    var v3 float32 = ((3.14159 * 3.14159) + (float32(((v2 / v2) * ((v2 / v2) / v2))) - 1e10))
    var v4 uint32 = 4294967295
    var v5 uint = 444
    var v6 uint16 = 32767
    v7 := []int{-304, 32768}
    fmt.Println(v1, v5, v2)
    v8 := GenS1{F0: 316}
    switch "foo bar" {
    case "\n\t":
        var v9 int16 = 383
        _ = v9
    case "":
        v10 := GenS2{F0: false, F1: 255, F2: (v5 & (^(v5)))}
        _ = v10
    default:
        if false {
            var v11 uint64 = 4294967295
            fmt.Println(v4)
            _ = v11
        }
    }
    if (-666 == 42) {
        var v12 string = "hello"
        _ = v12
    }
    fmt.Println(v5)
    v6 = (v6 ^ 54)
    v13 := GenS1{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
