package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int; F2 uint8 }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = (0.0 - (0.1 / 0.1))
    var v3 uint = 144
    v4 := make(map[uint32]bool)
    v5 := make(map[int]uint32)
    var v6 uint32 = 2147483647
    v7 := GenS1{F0: 9223372036854775807, F1: v6}
    var v8 uint32 = 743
    var v9 int64 = 1
    var v10 string = "a"
    v11 := GenS1{F0: 42, F1: v6}
    fmt.Println(v6, v9)
    v3 += (v3 / ((v3) | 1))
    var v12 float32 = 0.0
    v3 = (v3 + 1)
    if (950 <= 9) {
        v13 := []int{0, 32767, 564, 931}
        var v14 bool = (false && false)
        switch v10 {
        case "a":
            v15 := GenS2{F0: 844, F1: 127, F2: uint8(v6)}
            _ = v15
        case "\n\t":
            if (!v14) {
                fmt.Println(v6, v3, v12)
                v9 = 1
            }
        default:
            v2 += ((v2 - 1e10) + (v2 - -2.25))
        }
        _ = v13
        _ = v14
    } else {
        switch v10 {
        case "a":
            fmt.Println(v6, v6)
        case "hello":
            v2 = v2
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
