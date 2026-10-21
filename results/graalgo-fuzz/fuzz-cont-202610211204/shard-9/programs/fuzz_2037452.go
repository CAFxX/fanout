package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint16; F2 bool }
func genFn0() uint {
    var v7 float32 = 3.14159
    var v8 int64 = -9223372036854775808
    v9 := GenS1{F0: 2147483647}
    v7 = ((((v7 * (0.1 / v7)) / v7) * (0.0 * v7)) / (1e10 * v7))
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 int64) bool {
    if (1000 == 32767) {
        var v10 float32 = (((0.1 - 0.1) / 1.5) - ((1e10 - 3.14159) + -2.25))
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := []int{(-(fi11)), (-(fi11))}
            _ = len(v12)
            v10 -= v10
            _ = fi11
            _ = v12
        }
        var v13 float64 = -2.25
        _ = v10
        _ = v13
    }
    switch 424 {
    case 309:
        p0 = p0
    }
    _ = p0
    return false
}

func main() {
    var v1 float64 = 0.0
    var v2 float64 = (v1 + (v1 - 0.1))
    var v3 int32 = -1
    var v4 bool = (0 <= 128)
    v5 := make(map[uint32]string)
    var v6 int16 = -1
    v14 := GenS2{F0: 747, F1: 331, F2: v4}
    v15 := GenS2{F0: 32767, F1: 65535, F2: ((v3 % ((-1) | 1)) > -305)}
    delete(v5, 918)
    fmt.Println(v3, v4, v3)
    if (2 >= 876) {
        v6 += 1
        if (!v4) {
            for fi16 := 0; fi16 < 6; fi16++ {
                v6 -= int16(v3)
                fmt.Println(v1, v1)
                _ = fi16
            }
        }
    }
    v2 += (v2 * 0.0)
    for fi17 := 0; fi17 < 17; fi17++ {
        v1 += (((v2 + (1e10 / v1)) * (1e10 + ((v2 / v2) * v2))) / v2)
        _ = fi17
    }
    v4 = false
    var v18 bool = v4
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
