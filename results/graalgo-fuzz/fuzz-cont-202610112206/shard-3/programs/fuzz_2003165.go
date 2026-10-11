package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint8 }

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]int)
    var v3 float64 = 0.0
    var v4 uint = 2
    if false {
        v2[2147483648] = -9223372036854775808
    }
    if (560 != 32768) {
        fmt.Println(v3)
    }
    if true {
        v5 := []int{-2147483648, -2147483648, 956}
        var v6 int32 = 897
        v6 = 42
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: (false && false), F1: -327, F2: -1}
    fmt.Println(v1, v4)
    v1 = v1
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS2{F0: 42, F1: 65535, F2: 0}
        var v10 string = "foo bar"
        v2[2] = (fi8 << ((v4) & 63))
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint8 = 184
    fmt.Println(v4, v11, v3)
    v12 := GenS1{F0: true, F1: 0, F2: -2}
    v3 -= (float64(float32(float64(0.1))) * v3)
    fmt.Println(v3)
    var v13 int16 = 697
    if (42 < 55) {
        var v14 int64 = -1
        _ = v14
    }
    var v15 int16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
