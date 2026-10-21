package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int16; F2 int }

func main() {
    var v1 int8 = -1
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]uint32)
    var v4 int8 = v1
    if (1e10 == 0.0) {
        var v5 float64 = 1e10
        _ = v5
    }
    fmt.Println(v1)
    var v6 uint16 = 213
    if false {
        v1 = (v1 / ((v1) | 1))
        delete(v3, 42)
        var v7 string = "\n\t"
        _ = v7
    } else {
        switch ("a" + "foo bar") {
        case "hello":
            fmt.Println(v2, v4, v1)
        case "":
            v8 := []int{-2147483648, 127, 42, 4294967295}
            _ = v8
        default:
            v9 := GenS1{F0: 1000}
            _ = v9
        }
        var v10 int16 = -82
        _ = v10
    }
    if ((-2.25 - 1.5) <= (0.1 * 0.0)) {
        var v11 uint8 = 68
        _ = v11
    }
    if true {
        if true {
            var v12 string = "\n\t"
            _ = v12
        }
        fmt.Println(v1, v1, v2)
        v1 = v4
    }
    var v13 uint64 = 320
    var v14 int8 = (v4 << ((2) & 63))
    v6 -= (v6 + v6)
    v15 := GenS2{F0: v6, F1: 42, F2: -1}
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
