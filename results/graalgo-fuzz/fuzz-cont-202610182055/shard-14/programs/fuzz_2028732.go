package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint }
func genFn0() int8 {
    var v4 int8 = -107
    var v5 uint64 = 42
    _ = v4
    _ = v5
    return (v4 << ((614) & 63))
}

func main() {
    var v1 int8 = 127
    v2 := []int{32768, 128, -2147483648, 2147483647}
    var v3 uint8 = 2
    var v6 int16 = -32
    var v7 int32 = -848
    var v8 int16 = v6
    var v9 int32 = (-(v7))
    switch "hello" {
    case "hello":
        var v10 bool = true
        _ = v10
    case "a":
        v3 += v3
    case "foo bar":
        switch "" {
        case "":
            var v11 uint64 = 42
            _ = v11
        case "hello":
            fmt.Println(v1, v9)
        case "a":
            for fi12 := 0; fi12 < 14; fi12++ {
                v7 += (v7 << ((2) & 63))
                _ = fi12
            }
        }
    }
    var v13 uint = 4294967295
    var v14 uint = v13
    var v15 uint16 = 2
    if (!((0.0 / (((1e10 / (1e10 * 3.14159)) / 0.1) + (0.0 - -2.25))) <= -2.25)) {
        v16 := GenS1{F0: true, F1: (v3 * (v3 << ((v13) & 63))), F2: (v14 & 65535)}
        _ = v16
    }
    v15 = (v15 | v15)
    switch 32768 {
    case 42:
        var v17 int16 = (^(v6))
        _ = v17
    case 554:
        var v18 int16 = (v6 & 683)
        _ = v18
    default:
        v9 = 2147483647
    }
    fmt.Println(v7, v7, v14)
    v2 = append(v2, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
