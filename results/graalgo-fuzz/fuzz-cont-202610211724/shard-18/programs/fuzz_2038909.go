package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
func genFn0(p0 int, p1 uint8, p2 uint64) int16 {
    var v6 int64 = -9223372036854775808
    p1 = p1
    v7 := []int{559, (p0 * p0)}
    for fi8 := 0; fi8 < 6; fi8++ {
        if (int64(p1) != (-(v6))) {
            if (!((v6 ^ -735) > v6)) {
                _ = len(v7)
                var v9 bool = false
                _ = v9
            }
        } else {
            if false {
                var v10 float32 = 1e10
                var v11 int64 = v6
                _ = v10
                _ = v11
            }
        }
        if len(v7) > 0 {
            v7[0] = (^(fi8))
        }
        _ = fi8
    }
    var v12 int = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    _ = v7
    return -128
}

func main() {
    var v1 uint16 = 42
    var v2 uint8 = 2
    var v3 uint16 = v1
    var v4 int = 58
    var v5 int32 = -1
    v13 := []int{2147483647, v4}
    v13 = append(v13, int(v3))
    v13 = append(v13, v4)
    var v14 uint = 255
    if len(v13) > 0 {
        v13[0] = v4
    }
    if ((!true) || (false || (419 != 2))) {
        v4 = -9223372036854775808
        var v15 int8 = -31
        var v16 bool = false
        _ = v15
        _ = v16
    } else {
        var v17 int64 = 205
        _ = v17
    }
    if (-2.25 > 1.5) {
        if len(v13) > 0 {
            v13[0] = -228
        }
        if (v2 <= v2) {
            v1 += 255
            var v18 int32 = v5
            _ = v18
        }
    }
    if true {
        v3 = uint16(v14)
    }
    var v19 float64 = (3.14159 - (1e10 - -2.25))
    v20 := GenS1{F0: v2, F1: (v5 >> (((v14 % ((2) | 1))) & 63))}
    _ = len(v13)
    v2 = (v2 ^ (v2 * uint8(v14)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
