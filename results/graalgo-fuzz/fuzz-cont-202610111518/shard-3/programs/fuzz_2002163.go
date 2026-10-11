package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint; F1 uint32 }

func main() {
    var v1 float64 = -2.25
    var v2 int64 = -1
    v3 := make(map[uint32]int)
    v2 += v2
    var v4 int32 = 396
    switch "a" {
    case "":
        v1 = (v1 + 3.14159)
    case "\n\t":
        v5 := GenS1{F0: 226, F1: -128}
        _ = v5
    case "a":
        switch -9223372036854775808 {
        case 1:
            fmt.Println(v4, v1, v1)
        }
    }
    v6 := GenS2{F0: 2, F1: 849}
    fmt.Println(v4)
    v7 := GenS1{F0: 2, F1: int8(v4)}
    if (1 != 756) {
        var v8 int = -10
        var v9 int = -9223372036854775808
        _ = v8
        _ = v9
    }
    fmt.Println(v4)
    v2 = v2
    if (2147483648 != 2) {
        v2 = v2
        v10 := []int{-612, -115, 1}
        _ = v10
    } else {
        var v11 float64 = (v1 - v1)
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
