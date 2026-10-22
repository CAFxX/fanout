package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
func genFn0(p0 int64) uint {
    p0 += 1
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int32 = 0
        _ = fi7
        _ = v8
    }
    v9 := []int{}
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(fi10, p0)
        var v11 int8 = 1
        var v12 int32 = 2147483647
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v9
    return 902
}

func main() {
    var v1 float64 = 0.1
    var v2 uint64 = 2
    v3 := make(map[uint32]string)
    var v4 int64 = 1
    v5 := make(map[string]bool)
    v6 := []int{127, -9223372036854775808, 42, -881}
    var v13 int = 1
    v13 = 2147483647
    v2 -= (v2 / ((18446744073709551615) | 1))
    fmt.Println(v2)
    for fi14 := 0; fi14 < 9; fi14++ {
        for fi15 := 0; fi15 < 6; fi15++ {
            if ((0.0 > 1.5) && (false || (true || true))) {
                if ((!(((v1 + -2.25) <= v1) || false)) || (!(495 >= -866))) {
                    if (-439 != 42) {
                        v16 := GenS1{F0: 127, F1: genFn0(-9223372036854775808)}
                        var v17 uint8 = 33
                        _ = v16
                        _ = v17
                    }
                    var v18 float32 = float32(v1)
                    _ = v18
                }
                var v19 uint16 = 2
                _ = v19
            } else {
                v20 := GenS1{F0: -40, F1: 1}
                var v21 float64 = (float64(3.14159) - (v1 * 0.1))
                _ = v20
                _ = v21
            }
            _ = fi15
        }
        _ = fi14
    }
    v13 = (v13 << ((186) & 63))
    if true {
        v13 += (v13 >> ((45) & 63))
    } else {
        v1 -= (-2.25 / v1)
        v22 := v5["hello"]
        var v23 uint32 = 677
        _ = v22
        _ = v23
    }
    v1 = (1.5 + -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
