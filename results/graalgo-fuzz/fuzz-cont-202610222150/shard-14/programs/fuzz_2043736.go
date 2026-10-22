package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }

func main() {
    var v1 float32 = 1e10
    var v2 int32 = -1
    var v3 int8 = -47
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]int)
    v6 := []uint32{1, 4294967295, 2147483647}
    v1 += 1e10
    v7 := v5[4294967295]
    v3 += -75
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: false, F1: (fi8 / (((v7 ^ (v7 - fi8))) | 1))}
        _ = fi8
        _ = v9
    }
    _ = len(v5)
    for fi10 := 0; fi10 < 14; fi10++ {
        v6 = append(v6, 65535)
        var v11 int = (v7 | (^(v7)))
        if (v3 <= (v3 / ((v3) | 1))) {
            v11 += (^(v7))
        }
        _ = fi10
        _ = v11
    }
    var v12 int = v7
    _ = len(v5)
    if (9223372036854775807 < -427) {
        v3 = v3
    }
    fmt.Println(v3, v1)
    v3 = v3
    if false {
        for fi13 := 0; fi13 < 7; fi13++ {
            v3 += (-(v3))
            v14 := v5[65535]
            var v15 int64 = 42
            _ = fi13
            _ = v14
            _ = v15
        }
    } else {
        v6 = append(v6, 4294967295)
        for fi16 := 0; fi16 < 19; fi16++ {
            v17 := GenS1{F0: (!false), F1: (v7 * (v7 << ((1) & 63)))}
            _ = fi16
            _ = v17
        }
        fmt.Println(v2, v1)
    }
    v1 -= ((((float32(-2.25) * ((float32(float64(-2.25)) + v1) / (3.14159 - v1))) / 0.1) - (v1 * 0.1)) + (1.5 * float32((float64(((float32((float64((((float32(float64(v1)) * 3.14159) + float32((1e10 - float64((v1 - float32((0.1 + float64(v1)))))))) + (v1 * (-2.25 + (0.1 * (v1 * (float32(-2.25) - 3.14159))))))) / 1.5)) / v1) / (((float32((3.14159 * float64(0.1))) * float32(float64((v1 / 1.0)))) / (float32((3.14159 + 0.1)) - v1)) * v1))) * (0.0 + (3.14159 + 0.1))))))
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 bool = false
        fmt.Println(v2)
        v4[65535] = 4294967295
        _ = fi18
        _ = v19
    }
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
