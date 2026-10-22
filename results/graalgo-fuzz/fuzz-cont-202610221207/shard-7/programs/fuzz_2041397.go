package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int8 }
func genFn0(p0 uint, p1 int32) int32 {
    v5 := GenS1{F0: 95, F1: 184, F2: -128}
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := []int{fi6, 32767, 9223372036854775807}
        p0 = p0
        p1 -= 2147483647
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v5
    return p1
}

func main() {
    var v1 float64 = ((float64(((float32(-2.25) + 1e10) / -2.25)) / (-2.25 + 0.0)) * 1.5)
    var v2 int8 = -84
    var v3 int64 = 642
    var v4 int32 = 102
    v3 = v3
    v3 += (v3 | 234)
    if (128 <= 1) {
        v8 := []int{}
        v9 := GenS1{F0: 576, F1: 42, F2: (v2 % ((42) | 1))}
        v1 -= (1e10 / v1)
        _ = v8
        _ = v9
    }
    var v10 int8 = (v2 & v2)
    var v11 int16 = 1000
    if (318 == 0) {
        v1 = float64(1.5)
        for fi12 := 0; fi12 < 1; fi12++ {
            v3 = (v3 + (v3 >> ((993) & 63)))
            switch "\n\t" {
            case "":
                var v13 int16 = v11
                _ = v13
            case "\n\t":
                v14 := []int{}
                _ = v14
            }
            v15 := []int{}
            _ = fi12
            _ = v15
        }
        var v16 uint32 = 42
        _ = v16
    } else {
        if (2 <= 137) {
            v17 := []int{32767, 42, 2147483647}
            var v18 uint32 = 0
            var v19 int64 = v3
            _ = v17
            _ = v18
            _ = v19
        } else {
            v20 := []int{941}
            _ = v20
        }
        v3 += v3
        var v21 int8 = v2
        _ = v21
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        v4 += genFn0(255, (v4 ^ 42))
        v23 := []int{fi22, (fi22 + int(v11)), (fi22 >> ((2) & 63)), (fi22 & fi22)}
        _ = fi22
        _ = v23
    }
    v1 = v1
    if true {
        v24 := GenS1{F0: (v11 >> ((552) & 63)), F1: 1, F2: v2}
        for fi25 := 0; fi25 < 9; fi25++ {
            v26 := []int{fi25, fi25, fi25}
            if (!false) {
                fmt.Println(v2, v3)
            }
            _ = fi25
            _ = v26
        }
        v1 += v1
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
