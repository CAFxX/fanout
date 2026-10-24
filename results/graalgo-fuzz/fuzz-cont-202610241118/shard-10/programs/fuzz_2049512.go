package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int }
func genFn0(p0 uint32) int32 {
    var v6 int = 127
    v7 := GenS1{F0: 42, F1: "", F2: 127}
    p0 -= 260
    if true {
        v8 := []int{(v6 | v6), -687, (v6 & (-(v6)))}
        var v9 uint32 = p0
        _ = v8
        _ = v9
    } else {
        p0 = p0
        if (9223372036854775807 < 1) {
            var v10 int32 = 993
            _ = v10
        } else {
            v6 += (v6 & (v6 << ((0) & 63)))
            fmt.Println(v6)
            v11 := []int{32768, (v6 - (v6 << ((uint(p0)) & 63))), v6, v6}
            _ = v11
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint64, p1 bool) int8 {
    var v12 uint = 4294967295
    v12 = 42
    _ = p0
    _ = p1
    _ = v12
    return 105
}

func main() {
    var v1 int32 = 0
    var v2 uint16 = 65535
    var v3 uint = 18446744073709551615
    var v4 uint8 = 255
    var v5 int16 = 1
    var v13 float64 = 1e10
    for fi14 := 0; fi14 < 11; fi14++ {
        v4 += (v4 * (v4 >> (((v3 / (((v3 / ((967) | 1))) | 1))) & 63)))
        var v15 int8 = 42
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 14, F1: ("" + ""), F2: 32767}
    var v17 int = 2147483647
    fmt.Println(v1)
    if (4294967295 > 18446744073709551615) {
        var v18 uint64 = 2
        if (0.1 >= (0.0 + 1.5)) {
            var v19 float32 = ((0.1 * ((float32(v13) * float32(v13)) - -2.25)) + float32((v13 / v13)))
            v1 -= 1
            var v20 int64 = -185
            _ = v19
            _ = v20
        }
        _ = v18
    }
    v5 -= (v5 % ((v5) | 1))
    v4 = (^(v4))
    v21 := GenS1{F0: v2, F1: ("" + "\n\t"), F2: (v17 % (((^(v17))) | 1))}
    var v22 uint32 = 905
    v23 := []int{}
    var v24 float64 = v13
    v1 = -430
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
