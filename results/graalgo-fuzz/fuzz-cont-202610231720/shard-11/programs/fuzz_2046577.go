package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int }
func genFn0() int32 {
    v7 := []int{2147483647}
    var v8 uint16 = 32767
    if (65535 == 817) {
        v8 = 1
    } else {
        fmt.Println(v8)
        for fi9 := 0; fi9 < 13; fi9++ {
            _ = len(v7)
            _ = fi9
        }
    }
    _ = v7
    _ = v8
    return -560
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []uint32{v1, (^(v1)), 65535, (^(v1))}
    v3 := make(map[int]int)
    var v4 int32 = 2147483647
    var v5 uint = 408
    var v6 uint16 = 624
    v10 := GenS1{F0: 32767, F1: v6, F2: 0}
    v5 += (v5 + (^(v5)))
    v1 += (^(v1))
    v5 = (v5 ^ v5)
    var v11 uint = (v5 / ((v5) | 1))
    var v12 uint32 = v1
    switch ("a" + ("hello" + "")) {
    case "hello":
        fmt.Println(v11, v1)
    case "a":
        v2 = append(v2, v12)
    case "foo bar":
        fmt.Println(v6)
    }
    var v13 uint16 = v6
    v14 := GenS1{F0: 42, F1: 42, F2: -883}
    _ = len(v2)
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 uint64 = 811
        v5 += uint(v6)
        var v17 int8 = 72
        _ = fi15
        _ = v16
        _ = v17
    }
    v6 = v13
    switch "" {
    case "foo bar":
        var v18 float64 = float64((0.1 * 1.5))
        _ = v18
    case "\n\t":
        fmt.Println(v5, v13)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
