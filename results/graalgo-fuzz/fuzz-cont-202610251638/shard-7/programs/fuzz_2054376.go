package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint32; F1 string; F2 int }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 int16 = -247
    var v4 int32 = -2147483648
    var v5 uint8 = 255
    var v6 uint = 0
    v6 += v6
    fmt.Println(v4)
    switch "" {
    case "":
        for fi7 := 0; fi7 < 6; fi7++ {
            v4 = v4
            _ = fi7
        }
    case "hello":
        fmt.Println(v4, v3)
    case "\n\t":
        fmt.Println(v6)
    default:
        v5 = 0
    }
    var v8 uint64 = 42
    v9 := GenS1{F0: 42, F1: 2}
    v4 -= (v4 % (((-(v4))) | 1))
    var v10 uint16 = 851
    if true {
        fmt.Println(v10, v3)
        var v11 bool = false
        _ = v11
    } else {
        fmt.Println(v6)
    }
    if ((^(v10)) >= (v10 - (^(v10)))) {
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS2{F0: 2147483648, F1: "\n\t", F2: fi12}
            _ = fi12
            _ = v13
        }
    } else {
        v14 := GenS1{F0: v10, F1: 18446744073709551615}
        _ = v14
    }
    _ = len(v1)
    v4 -= 2147483647
    var v15 float64 = 0.0
    v8 = (^(v8))
    v16 := []int{1, 128, -2147483648, -1}
    v4 -= v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
