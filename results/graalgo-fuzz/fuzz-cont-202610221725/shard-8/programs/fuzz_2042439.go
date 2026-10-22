package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint }
func genFn0(p0 uint32) uint {
    if true {
        if (4294967295 > -794) {
            switch -743 {
            case -867:
                var v7 float32 = -2.25
                _ = v7
            default:
                p0 = (p0 | 42)
            }
            for fi8 := 0; fi8 < 6; fi8++ {
                v9 := []int{(fi8 * (-(fi8))), fi8, (fi8 - fi8), 42}
                _ = fi8
                _ = v9
            }
        } else {
            p0 += p0
            var v10 int8 = 68
            _ = v10
        }
        v11 := GenS1{F0: 42, F1: 1}
        p0 -= p0
        _ = v11
    }
    v12 := GenS2{F0: -2147483648, F1: (909 <= 4294967295)}
    p0 = 494
    var v13 int16 = 441
    _ = p0
    _ = v12
    _ = v13
    return 910
}
func genFn1(p0 int64, p1 int, p2 int16) int16 {
    if (-868 != -1) {
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS1{F0: 1, F1: -1}
            _ = fi14
            _ = v15
        }
    }
    var v16 uint32 = 4294967295
    if false {
        var v17 float32 = (-2.25 * 0.1)
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return p2
}

func main() {
    var v1 float32 = 1.5
    var v2 string = ("" + "")
    var v3 int32 = -186
    var v4 int = 9223372036854775807
    var v5 float32 = (float32((float64(((float32(0.1) * 1.5) * (v1 * (0.1 / (0.1 / v1))))) - 1.5)) * (v1 - v1))
    v6 := []string{""}
    v18 := GenS2{F0: 42, F1: false}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v1 -= float32(1e10)
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := GenS1{F0: 1, F1: (v3 / ((-2147483648) | 1))}
        _ = fi19
        _ = v20
    }
    v6 = append(v6, "a")
    if false {
        v6 = append(v6, "\n\t")
        var v21 float64 = 1.5
        _ = v21
    }
    fmt.Println(v1, v1)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v22 := GenS3{F0: 4294967295}
    if true {
        var v23 float32 = v5
        _ = v23
    }
    _ = len(v6)
    var v24 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
