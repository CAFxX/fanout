package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int64 }
func genFn0(p0 int32, p1 int64) uint16 {
    p0 = (p0 + (p0 << ((42) & 63)))
    if false {
        p1 += p1
        var v5 uint32 = 833
        _ = v5
    } else {
        p1 -= (p1 % ((p1) | 1))
    }
    p0 = p0
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int16 = 1
        _ = fi6
        _ = v7
    }
    var v8 uint = 0
    _ = p0
    _ = p1
    _ = v8
    return 255
}
func genFn1(p0 int, p1 bool, p2 int) bool {
    if (1 < -1) {
        for fi9 := 0; fi9 < 9; fi9++ {
            fmt.Println(p2, p2)
            _ = fi9
        }
        if p1 {
            v10 := []int{(p0 & 1), (^(p0)), (p2 % (((p2 << ((343) & 63))) | 1))}
            _ = v10
        } else {
            p0 -= p2
        }
    } else {
        p2 += -223
    }
    p0 = 127
    var v11 int32 = 512
    var v12 uint16 = 155
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return p1
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]uint32)
    var v3 uint = 1
    v4 := []uint32{0, 0, 2, 899}
    v13 := GenS1{F0: 42, F1: true}
    v14 := GenS2{F0: -1}
    var v15 uint32 = 802
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 uint8 = 171
        v18 := v2[v15]
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 float32 = 1.5
    for fi20 := 0; fi20 < 15; fi20++ {
        v15 = 42
        fmt.Println(v19)
        _ = fi20
    }
    var v21 uint32 = (v15 & (v15 >> ((1) & 63)))
    for fi22 := 0; fi22 < 8; fi22++ {
        if ("" <= (("a" + "") + "")) {
            v4 = append(v4, v15)
            v4 = append(v4, 979)
        } else {
            if false {
                if (318 == v21) {
                    var v23 float64 = (1e10 * 0.1)
                    _ = v23
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                v24 := GenS1{F0: 42, F1: (42 >= -127)}
                var v25 int16 = -1
                var v26 float32 = (v19 - (v19 / ((v19 + v19) / ((0.0 + v19) / v19))))
                _ = v24
                _ = v25
                _ = v26
            }
            fmt.Println(v21)
        }
        if true {
            v1 = (2 < 206)
        }
        _ = fi22
    }
    var v27 int = 4294967295
    fmt.Println(v3, v27, v1)
    for fi28 := 0; fi28 < 19; fi28++ {
        var v29 int32 = 2147483647
        fmt.Println(v29, v15)
        _ = fi28
        _ = v29
    }
    v15 += 178
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v4
}
