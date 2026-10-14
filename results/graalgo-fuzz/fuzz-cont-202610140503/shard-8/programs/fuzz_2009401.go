package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
func genFn0(p0 bool, p1 int32, p2 int64) int32 {
    v5 := []int{128, 42, -357}
    if p0 {
        for fi6 := 0; fi6 < 15; fi6++ {
            p0 = (false || (-1 < 32767))
            _ = len(v5)
            p1 += -813
            _ = fi6
        }
        p2 = p2
    } else {
        var v7 bool = p0
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p1 % ((p1) | 1))
}

func main() {
    var v1 int16 = 910
    var v2 int32 = -2147483648
    v3 := make(map[uint32]uint32)
    var v4 int32 = (-(v2))
    fmt.Println(v4, v4)
    var v8 uint = 2
    if (18446744073709551615 != 999) {
        v8 = (^(v8))
    } else {
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int16 = 42
            _ = fi9
            _ = v10
        }
    }
    var v11 uint32 = 2147483647
    var v12 int = 836
    v11 += uint32(v8)
    v4 = (v4 & v2)
    v4 += v2
    var v13 uint32 = v11
    var v14 float32 = ((((3.14159 + 0.0) + ((1e10 * -2.25) - 1e10)) - 3.14159) - 1e10)
    v15 := GenS1{F0: v1, F1: 135}
    var v16 int8 = 93
    var v17 float32 = v14
    switch (v12 >> ((4294967295) & 63)) {
    case -2147483648:
        v8 -= (v8 / (((v8 + (v8 | (v8 >> (((v8 - (v8 << ((571) & 63)))) & 63))))) | 1))
    case -1:
        v11 -= v13
    case -9223372036854775808:
        var v18 uint = (v8 | v8)
        _ = v18
    }
    v13 += v13
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
