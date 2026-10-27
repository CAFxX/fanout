package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int32 }
func genFn0(p0 bool) int16 {
    p0 = p0
    var v5 int64 = 624
    _ = p0
    _ = v5
    return -341
}

func main() {
    v1 := []int{}
    var v2 int32 = 0
    var v3 float64 = (1.5 + 1e10)
    var v4 int16 = 0
    if true {
        v6 := GenS1{F0: false, F1: 348, F2: (-(v2))}
        v7 := GenS1{F0: true, F1: -9223372036854775808, F2: 42}
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
        _ = v6
        _ = v7
    } else {
        var v8 float64 = (float64(float32(v3)) / v3)
        for fi9 := 0; fi9 < 17; fi9++ {
            if len(v1) > 0 {
                v1[0] = (fi9 * (fi9 - 4294967295))
            }
            _ = fi9
        }
        for fi10 := 0; fi10 < 2; fi10++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v11 := GenS1{F0: (1 >= 1), F1: 42, F2: v2}
            v12 := GenS1{F0: (float32((-2.25 - float64((((1e10 * ((0.1 + 1.5) - 3.14159)) - float32(v3)) - 1e10)))) == 3.14159), F1: 823, F2: -2147483648}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v8
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            fmt.Println(v2)
            _ = fi14
        }
        _ = fi13
    }
    v4 = -1
    fmt.Println(v2)
    v15 := GenS1{F0: false, F1: 42, F2: -1}
    v2 += 1
    v4 -= 1
    v3 += 3.14159
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 -= (v4 % (((v4 & (v4 | (v4 % ((v4) | 1))))) | 1))
    fmt.Println(v4, v3)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
