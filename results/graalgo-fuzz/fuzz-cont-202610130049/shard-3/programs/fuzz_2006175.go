package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 163
    v3 := []int{}
    var v4 bool = (1 >= 432)
    var v5 int32 = 42
    v6 := GenS1{F0: (v2 * (v2 - (v2 >> ((65535) & 63)))), F1: (!v4), F2: 163}
    fmt.Println(v4, v2, v2)
    v1[42] = 291
    fmt.Println(v5, v4)
    v5 -= -58
    v2 -= 660
    var v7 float64 = 3.14159
    var v8 int64 = 0
    if (!(0.1 < -2.25)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v9 := GenS2{F0: -1}
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        _ = v9
    } else {
        _ = len(v3)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch 9223372036854775807 {
    case 128:
        fmt.Println(v2, v7, v8)
    case 962:
        fmt.Println(v7)
    }
    fmt.Println(v8, v2, v8)
    var v10 uint8 = 100
    var v11 int32 = -374
    var v12 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
}
