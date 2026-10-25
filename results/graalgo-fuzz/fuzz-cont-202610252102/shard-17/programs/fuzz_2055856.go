package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint32 }
type GenS2 struct { F0 int8; F1 int16 }
func genFn0(p0 uint32, p1 uint64, p2 bool) uint {
    var v6 int32 = -2147483648
    fmt.Println(v6, p1, v6)
    v7 := []int{0, 599}
    v6 = 42
    if len(v7) > 0 {
        v7[0] = 1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return uint(p1)
}

func main() {
    var v1 int = 847
    var v2 uint32 = 2
    v3 := []uint32{}
    v4 := make(map[uint32]uint32)
    var v5 uint32 = (v2 ^ v2)
    delete(v4, (v5 & (v5 << ((705) & 63))))
    var v8 int8 = -128
    v2 += v5
    v8 -= (v8 + v8)
    var v9 int8 = (v8 % ((v8) | 1))
    var v10 float32 = 0.0
    switch ("hello" + "hello") {
    case "":
        v5 += (v5 >> ((985) & 63))
    case "foo bar":
        v5 = 375
    case "hello":
        fmt.Println(v5, v10)
    default:
        v11 := GenS1{F0: -9223372036854775808, F1: "\n\t", F2: v5}
        _ = v11
    }
    if true {
        var v12 uint16 = 0
        var v13 int8 = (v9 - (-(v8)))
        v14 := GenS2{F0: (^(v13)), F1: 669}
        _ = v12
        _ = v13
        _ = v14
    } else {
        v15 := v4[(v2 * v5)]
        switch (v1 >> ((2) & 63)) {
        case 128:
            _ = len(v3)
        case 127:
            if (432 > 453) {
                var v16 string = ("\n\t" + "a")
                _ = v16
            }
        case -190:
            var v17 int8 = (v9 - (v8 / (((v8 % (((v8 / ((51) | 1))) | 1))) | 1)))
            _ = v17
        }
        v10 = float32(0.1)
        _ = v15
    }
    v18 := GenS1{F0: 9223372036854775807, F1: "foo bar", F2: 4294967295}
    var v19 uint = 820
    v4[v5] = v5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
