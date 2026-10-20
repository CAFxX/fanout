package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 42
    v3 := []int{-1, -1}
    var v4 uint16 = 65535
    var v5 uint8 = 159
    var v6 int64 = -1
    for fi7 := 0; fi7 < 12; fi7++ {
        v3 = append(v3, int(v6))
        v8 := GenS1{F0: (fi7 + (fi7 ^ fi7))}
        v1[2147483648] = (!true)
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 127}
    v2 = (v2 / ((v2) | 1))
    v2 -= v2
    if (true && (true && ((v5 / ((v5) | 1)) != (^(v5))))) {
        fmt.Println(v4)
    } else {
        v5 += 1
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v10 := GenS2{F0: 65535, F1: -2147483648}
    fmt.Println(v5, v2, v4)
    var v11 uint32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
