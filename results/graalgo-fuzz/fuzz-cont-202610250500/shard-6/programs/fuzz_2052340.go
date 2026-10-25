package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int16; F2 int8 }
func genFn0(p0 uint8, p1 uint64) uint16 {
    v7 := []int{127, -869, 844, 127}
    p1 -= (p1 >> ((42) & 63))
    var v8 uint32 = 278
    _ = len(v7)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 596
}

func main() {
    var v1 uint8 = 1
    var v2 int = 286
    var v3 uint64 = 0
    var v4 uint8 = (v1 % ((255) | 1))
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]int)
    if (((190 != uint32(v4)) && (0.0 != 0.1)) || (((float64(float32((1e10 / -2.25))) * float64((1e10 * 1e10))) < 0.0) && true)) {
        var v9 string = ""
        v10 := GenS1{F0: 2}
        var v11 float32 = float32(0.1)
        _ = v10
        _ = v11
        _ = v9
    } else {
        var v12 float32 = (0.0 - (3.14159 + ((-2.25 + 0.1) - -2.25)))
        _ = v12
    }
    v1 += v4
    v13 := v5[2]
    v2 -= (v2 >> ((uint(v1)) & 63))
    switch "\n\t" {
    case "\n\t":
        for fi14 := 0; fi14 < 17; fi14++ {
            v15 := v6[2147483648]
            var v16 uint64 = v3
            _ = fi14
            _ = v15
            _ = v16
        }
    case "foo bar":
        var v17 float64 = 3.14159
        _ = v17
    default:
        v18 := GenS2{F0: v1, F1: 42, F2: 100}
        _ = v18
    }
    var v19 int16 = -1
    for fi20 := 0; fi20 < 15; fi20++ {
        for fi21 := 0; fi21 < 4; fi21++ {
            v22 := GenS2{F0: (v4 << ((18446744073709551615) & 63)), F1: v19, F2: 100}
            v23 := []int{(fi21 | -9223372036854775808), fi20, (^(fi20)), fi21}
            _ = fi21
            _ = v22
            _ = v23
        }
        if ((v19 == v19) || (!true)) {
            v24 := GenS2{F0: (^(v1)), F1: (v19 & (^(v19))), F2: 56}
            fmt.Println(v2)
            _ = v24
        }
        v25 := GenS1{F0: 642}
        _ = fi20
        _ = v25
    }
    fmt.Println(v13)
    v26 := []int{(v2 + v2), 4294967295, v2, (-(v2))}
    v19 += (v19 % ((42) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
