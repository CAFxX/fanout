package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 string }
type GenS2 struct { F0 uint8; F1 int8; F2 uint8 }

func main() {
    var v1 int16 = -599
    var v2 float32 = (1e10 + 1.5)
    v3 := []uint32{}
    var v4 bool = (!false)
    v5 := make(map[int]int)
    var v6 float32 = (v2 / -2.25)
    v7 := GenS1{F0: 9223372036854775807, F1: 42, F2: "a"}
    v3 = append(v3, 619)
    v5[1] = -9223372036854775808
    v8 := GenS2{F0: 212, F1: 100, F2: 0}
    v9 := GenS2{F0: 72, F1: -128, F2: 0}
    v4 = (119 >= 0)
    v10 := GenS2{F0: 210, F1: -1, F2: 183}
    var v11 bool = false
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 uint8 = 13
        var v14 string = ""
        v3 = append(v3, 2147483648)
        _ = fi12
        _ = v13
        _ = v14
    }
    v6 = (v2 + 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
