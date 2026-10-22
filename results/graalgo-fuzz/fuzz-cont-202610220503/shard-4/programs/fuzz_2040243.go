package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint64 }
type GenS2 struct { F0 int32; F1 string; F2 uint8 }
func genFn0() int {
    v6 := []int{746}
    var v7 int16 = 42
    switch 128 {
    case 9223372036854775807:
        for fi8 := 0; fi8 < 17; fi8++ {
            if (1 != 0) {
                var v9 int8 = 0
                v7 -= (v7 - 1000)
                _ = v9
            } else {
                var v10 int16 = (v7 / ((v7) | 1))
                v10 -= (-(v10))
                v10 = (v10 << ((18446744073709551615) & 63))
                _ = v10
            }
            fmt.Println(fi8, v7)
            _ = fi8
        }
    }
    v7 += 428
    _ = v6
    _ = v7
    return -378
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = (3.14159 + -2.25)
    var v3 uint16 = 620
    var v4 float64 = v2
    var v5 uint32 = 2
    var v11 uint32 = (^(v5))
    var v12 uint8 = 226
    v4 -= -2.25
    v2 -= (v2 * v4)
    v3 = v3
    if false {
        v12 = 0
    } else {
        v13 := GenS1{F0: ("" + "a"), F1: "a", F2: 4294967295}
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        switch genFn0() {
        case -266:
            for fi15 := 0; fi15 < 19; fi15++ {
                fmt.Println(v11, v2)
                _ = fi15
            }
        case 397:
            v16 := GenS2{F0: -2147483648, F1: (("foo bar" + "") + "foo bar"), F2: v12}
            _ = v16
        case 0:
            if (!false) {
                v2 += v4
            }
        default:
            v3 = (v3 - v3)
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            fmt.Println(fi14, v3, fi17)
            _ = fi17
        }
        var v18 uint16 = 352
        _ = fi14
        _ = v18
    }
    v19 := v1[v11]
    var v20 float64 = 1.5
    v21 := GenS1{F0: "", F1: "\n\t", F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
