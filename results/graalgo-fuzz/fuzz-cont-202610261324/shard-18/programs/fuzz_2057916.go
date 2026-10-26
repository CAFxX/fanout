package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 string }
func genFn0(p0 int32) int64 {
    var v6 int = 9223372036854775807
    v7 := []int{(v6 % (((^(v6))) | 1)), v6}
    p0 += (-(p0))
    p0 += -1
    _ = p0
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int16, p1 int32, p2 string) uint32 {
    v8 := []int{993, 2147483647}
    var v9 bool = (true && false)
    for fi10 := 0; fi10 < 17; fi10++ {
        p0 = p0
        p0 = (p0 << ((42) & 63))
        if v9 {
            fmt.Println(fi10)
            v9 = v9
        }
        _ = fi10
    }
    var v11 int16 = p0
    for fi12 := 0; fi12 < 5; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            fmt.Println(v11)
            _ = fi13
        }
        var v14 int8 = -38
        _ = fi12
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 float32 = (-2.25 + (3.14159 / 1.0))
    v2 := make(map[int]uint32)
    v3 := []string{}
    var v4 int32 = -2147483648
    v5 := []uint32{0}
    v15 := GenS1{F0: genFn0(v4), F1: 107, F2: "a"}
    for fi16 := 0; fi16 < 8; fi16++ {
        fmt.Println(fi16)
        _ = fi16
    }
    var v17 float64 = float64(float32(float64((1e10 * v1))))
    _ = len(v5)
    for fi18 := 0; fi18 < 2; fi18++ {
        v4 += v4
        _ = fi18
    }
    v19 := GenS1{F0: -9223372036854775808, F1: 100, F2: ("hello" + ("" + ("hello" + "")))}
    if (0 <= 4294967295) {
        var v20 int16 = -360
        if false {
            if (65535 < 255) {
                v21 := GenS1{F0: 0, F1: 21, F2: "a"}
                switch 0 {
                case 2147483647:
                    v5 = append(v5, 2)
                }
                _ = v21
            } else {
                var v22 int = 56
                _ = v22
            }
            v23 := GenS1{F0: 42, F1: -128, F2: "a"}
            _ = v23
        } else {
            v17 = float64((((v1 + 0.0) - 0.0) + v1))
        }
        _ = v20
    }
    var v24 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v17
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
