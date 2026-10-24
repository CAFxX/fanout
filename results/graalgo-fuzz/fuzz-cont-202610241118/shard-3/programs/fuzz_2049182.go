package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint16 }
func genFn0() int16 {
    if false {
        var v6 int64 = -46
        _ = v6
    } else {
        v7 := []int{}
        var v8 int64 = 0
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        if true {
            var v10 int64 = 42
            v10 += v10
            _ = v10
        } else {
            switch (-(fi9)) {
            case -506:
                v11 := []int{fi9, (fi9 ^ (^(fi9))), fi9}
                _ = v11
            case -9223372036854775808:
                v12 := []int{fi9, -400, (fi9 | 127)}
                _ = v12
            default:
                v13 := []int{(fi9 ^ 4294967295), fi9, fi9, fi9}
                _ = v13
            }
            var v14 uint32 = 2147483648
            _ = v14
        }
        var v15 float32 = float32(0.1)
        _ = fi9
        _ = v15
    }
    var v16 int8 = -128
    if (97 == (v16 * (^(v16)))) {
        var v17 float64 = (1.5 / (0.0 - 3.14159))
        v18 := GenS1{F0: v16, F1: 2147483648}
        _ = v17
        _ = v18
    }
    _ = v16
    return 501
}

func main() {
    var v1 int8 = -93
    var v2 int16 = 1
    var v3 int8 = (v1 ^ 0)
    var v4 uint32 = 524
    v5 := make(map[int]uint32)
    for fi19 := 0; fi19 < 15; fi19++ {
        fmt.Println(v1, v3, fi19)
        v20 := GenS1{F0: 127, F1: 65535}
        _ = fi19
        _ = v20
    }
    var v21 int = -1
    v3 = (v1 | (v1 + (-(v1))))
    v22 := []int{(v21 / ((58) | 1)), -1}
    v3 = (v3 - 24)
    delete(v5, v21)
    var v23 bool = true
    v24 := GenS1{F0: v1, F1: (v4 | 4294967295)}
    var v25 float64 = (0.0 + -2.25)
    v4 -= 2147483648
    v26 := GenS2{F0: 87, F1: 897}
    v21 += -9223372036854775808
    for fi27 := 0; fi27 < 12; fi27++ {
        if v23 {
            v2 += (^(v2))
            fmt.Println(v3, v3)
        } else {
            var v28 uint64 = 0
            _ = v28
        }
        var v29 uint8 = 1
        _ = fi27
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
