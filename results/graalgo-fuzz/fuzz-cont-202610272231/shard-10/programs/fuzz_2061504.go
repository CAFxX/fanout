package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int8; F1 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 1
    v3 := make(map[string]bool)
    var v4 string = "foo bar"
    v5 := make(map[int]int)
    switch 2147483647 {
    case 42:
        v6 := GenS1{F0: v2, F1: 2147483647, F2: 845}
        _ = v6
    default:
        var v7 int64 = (v2 + (v2 - v2))
        _ = v7
    }
    v4 = ("" + v4)
    _ = len(v3)
    v2 -= (v2 + (v2 | (v2 | 232)))
    v2 -= v2
    switch v4 {
    case "hello":
        v2 = v2
    }
    delete(v3, v4)
    v8 := GenS2{F0: 100, F1: 1}
    if (-2147483648 == 0) {
        fmt.Println(v4)
        v2 -= 886
        v9 := []int{}
        _ = v9
    }
    v10 := GenS1{F0: v2, F1: 0, F2: 32767}
    v11 := GenS2{F0: -128, F1: 1}
    var v12 uint = 2
    var v13 uint8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
