package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint64 }
func genFn0() int32 {
    v5 := GenS1{F0: 0, F1: -866}
    var v6 uint32 = 42
    switch "a" {
    case "hello":
        v7 := []int{2147483647}
        _ = v7
    }
    v6 += (v6 + v6)
    for fi8 := 0; fi8 < 16; fi8++ {
        switch "\n\t" {
        case "\n\t":
            var v9 int32 = -3
            _ = v9
        case "foo bar":
            v6 -= (v6 >> ((710) & 63))
        default:
            v6 -= v6
        }
        _ = fi8
    }
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint8 = 255
    var v3 float32 = 0.1
    var v4 int8 = 72
    fmt.Println(v4)
    v2 = v2
    var v10 int16 = 1
    v4 -= (-(v4))
    v11 := []int{128, 42, 1, 0}
    v12 := GenS1{F0: 485, F1: int16(v1)}
    v4 += 0
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    v1 = 2
    v13 := GenS2{F0: (v2 % ((145) | 1)), F1: 783}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
