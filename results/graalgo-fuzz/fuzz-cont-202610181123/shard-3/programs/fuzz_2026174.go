package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int64 }

func main() {
    var v1 int16 = -32768
    var v2 float32 = -2.25
    v3 := make(map[int]string)
    fmt.Println(v2, v1)
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 bool = (((fi4 * (-(fi4))) != (-(fi4))) || false)
        v5 = (65535 == 54)
        v6 := GenS1{F0: 9223372036854775807, F1: 801, F2: 42}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int = 0
    v2 -= v2
    fmt.Println(v2, v1)
    switch ("hello" + "foo bar") {
    case "":
        for fi8 := 0; fi8 < 12; fi8++ {
            for fi9 := 0; fi9 < 15; fi9++ {
                if (2 >= 0) {
                    v2 -= v2
                    v2 = v2
                    _ = len(v3)
                }
                _ = fi9
            }
            var v10 int16 = v1
            _ = fi8
            _ = v10
        }
    }
    v11 := GenS1{F0: -9223372036854775808, F1: 0, F2: -1}
    fmt.Println(v7)
    switch "hello" {
    case "\n\t":
        var v12 int = 348
        _ = v12
    case "hello":
        _ = len(v3)
    case "":
        v13 := []int{v7, (v7 >> ((1) & 63)), (v7 | (v7 % ((v7) | 1)))}
        _ = v13
    default:
        var v14 uint8 = 0
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS1{F0: -1, F1: 745, F2: 774}
        v2 += ((v2 + 0.0) + 3.14159)
        _ = fi15
        _ = v16
    }
    v2 -= 0.0
    v7 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v7
}
