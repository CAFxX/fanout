package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 string }
func genFn0(p0 int32) int16 {
    p0 -= p0
    var v6 uint8 = 178
    fmt.Println(v6, p0)
    for fi7 := 0; fi7 < 12; fi7++ {
        v6 -= (v6 >> ((0) & 63))
        _ = fi7
    }
    _ = p0
    _ = v6
    return 1
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]uint32)
    v3 := []uint32{4, 2}
    var v4 int64 = 42
    v5 := make(map[uint32]uint32)
    if ((1 != 878) && false) {
        if (false || (2 <= 18446744073709551615)) {
            _ = len(v3)
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v1, v4)
    } else {
        v8 := GenS1{F0: 0}
        _ = len(v3)
        _ = v8
    }
    v1 = (v1 + v1)
    switch v1 {
    case "\n\t":
        v4 += v4
    }
    switch -2147483648 {
    case 9223372036854775807:
        v4 -= (v4 & (v4 + v4))
    case -669:
        var v9 uint32 = 4294967295
        _ = v9
    case 127:
        var v10 uint8 = 0
        _ = v10
    }
    var v11 int32 = 1
    v12 := GenS1{F0: 128}
    v11 = v11
    var v13 float32 = (((3.14159 * (((0.0 + 1.5) / -2.25) * ((1e10 / -2.25) - float32((1.5 + 3.14159))))) - 1.5) * (((-2.25 * (0.1 * 1.5)) - (3.14159 * 3.14159)) + 3.14159))
    v14 := GenS2{F0: 635, F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
