package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int32 }

func main() {
    var v1 uint8 = 255
    var v2 uint32 = 2147483647
    v3 := make(map[string]uint32)
    v4 := []int{-9223372036854775808}
    v5 := make(map[uint32]int)
    var v6 int16 = 42
    if len(v4) > 0 {
        v4[0] = -167
    }
    switch 1 {
    case -1:
        v6 = (v6 >> ((255) & 63))
    case 827:
        var v7 uint = 816
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        if (!(-42 < 127)) {
            var v9 int8 = -128
            fmt.Println(v1, v9, v1)
            v9 += 0
            _ = v9
        } else {
            v1 += 238
        }
        _ = fi8
    }
    v1 = 1
    v10 := GenS1{F0: true, F1: (v2 ^ (v2 >> ((0) & 63))), F2: 482}
    var v11 float32 = 3.14159
    v6 = -310
    for fi12 := 0; fi12 < 10; fi12++ {
        if true {
            var v13 uint32 = (^(v2))
            v14 := GenS1{F0: ((65535 > 811) && (false && ((^(v1)) == v1))), F1: (^(v13)), F2: 1}
            v1 = (v1 >> ((65535) & 63))
            _ = v13
            _ = v14
        } else {
            v4 = append(v4, (fi12 / (((fi12 * (fi12 >> ((0) & 63)))) | 1)))
            v4 = append(v4, fi12)
        }
        _ = fi12
    }
    if (false || true) {
        var v15 int = 32767
        var v16 uint16 = 1
        v17 := GenS1{F0: false, F1: 42, F2: -983}
        _ = v15
        _ = v16
        _ = v17
    }
    if (!false) {
        var v18 bool = (65535 > 0)
        v6 -= (v6 - (v6 + v6))
        _ = len(v4)
        _ = v18
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if ("" != "a") {
        fmt.Println(v2)
        v1 = 42
        if false {
            var v19 string = "hello"
            _ = v19
        }
    } else {
        v6 = -50
        fmt.Println(v2, v11)
        var v20 uint64 = 2
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
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
}
