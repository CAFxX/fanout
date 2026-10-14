package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 bool; F2 uint32 }
func genFn0(p0 int16) int8 {
    p0 = p0
    fmt.Println(p0)
    v4 := []int{9223372036854775807, 128, -1}
    _ = p0
    _ = v4
    return -111
}
func genFn1(p0 int64) int32 {
    p0 -= 81
    v5 := []int{1}
    _ = p0
    _ = v5
    return -1
}

func main() {
    var v1 float64 = 0.0
    var v2 uint32 = 2
    var v3 int16 = 44
    v1 = v1
    v2 += (^(v2))
    var v6 int8 = 64
    fmt.Println(v1, v6, v1)
    v3 = v3
    if (!false) {
        var v7 float64 = v1
        switch 127 {
        case -386:
            v2 = v2
        default:
            v8 := []int{int(v2), 9223372036854775807, -778}
            _ = v8
        }
        fmt.Println(v7, v1, v2)
        _ = v7
    } else {
        var v9 string = ""
        if (-280 <= -9223372036854775808) {
            v6 = genFn0(v3)
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 string = v9
                var v12 uint8 = 2
                v13 := GenS1{F0: v6}
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
        } else {
            switch 2147483647 {
            case 2147483647:
                fmt.Println(v3, v9)
            }
            v6 = (v6 % (((v6 >> ((42) & 63))) | 1))
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 uint64 = 42
                _ = fi14
                _ = v15
            }
        }
        v16 := []int{4294967295, -401}
        _ = v16
        _ = v9
    }
    switch 47 {
    case -2147483648:
        var v17 int = 2147483647
        _ = v17
    default:
        v3 = (v3 & (v3 & 1000))
    }
    fmt.Println(v3, v3)
    var v18 int8 = v6
    var v19 float32 = -2.25
    for fi20 := 0; fi20 < 18; fi20++ {
        fmt.Println(v19)
        if true {
            if false {
                fmt.Println(v18, v6, v18)
                v3 += 32767
            }
        }
        _ = fi20
    }
    v21 := GenS1{F0: v6}
    if ((v1 - 0.1) > 1e10) {
        v22 := GenS2{F0: 1000, F1: (255 != 255), F2: v2}
        _ = v22
    }
    v23 := []int{}
    var v24 int32 = 0
    v19 -= v19
    v23 = append(v23, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v6
}
