package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 bool }
type GenS2 struct { F0 int; F1 int32; F2 uint32 }
type GenS3 struct { F0 int64 }
func genFn0() int32 {
    v5 := []int{1, 9223372036854775807, 922}
    v6 := GenS1{F0: -1, F1: -1, F2: false}
    if (-278 <= -2147483648) {
        var v7 uint64 = 2
        for fi8 := 0; fi8 < 6; fi8++ {
            v5 = append(v5, (fi8 | 907))
            _ = fi8
        }
        v7 = (v7 % ((v7) | 1))
        _ = v7
    } else {
        var v9 uint16 = 813
        v10 := GenS2{F0: 2147483647, F1: -92, F2: 1}
        v9 = 65535
        _ = v10
        _ = v9
    }
    var v11 int16 = 1
    _ = v11
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1(p0 int16) uint8 {
    if false {
        switch 42 {
        case -526:
            if (false || ((float64(float32((1.5 - 0.1))) < 1e10) && true)) {
                p0 -= p0
                if true {
                    var v12 float64 = ((float64((float32((-2.25 - 0.1)) - float32(float64(3.14159)))) * 0.1) + (1e10 - (0.0 * ((-2.25 - 0.1) + 1e10))))
                    fmt.Println(v12, v12)
                    v13 := GenS3{F0: -800}
                    _ = v12
                    _ = v13
                } else {
                    if true {
                        var v14 int64 = -9223372036854775808
                        _ = v14
                    } else {
                        var v15 int32 = 2147483647
                        _ = v15
                    }
                }
            } else {
                var v16 int = 32768
                _ = v16
            }
        case -752:
            p0 = p0
        }
    }
    v17 := []int{936, 127}
    v18 := GenS2{F0: 42, F1: -2147483648, F2: 207}
    v19 := GenS1{F0: p0, F1: -1, F2: true}
    _ = p0
    _ = v17
    _ = v18
    _ = v19
    return 117
}

func main() {
    var v1 uint32 = 2147483647
    v2 := []int{127, 157, -2147483648, 42}
    v3 := make(map[int]bool)
    var v4 int = 1
    var v20 int64 = 9223372036854775807
    for fi21 := 0; fi21 < 2; fi21++ {
        for fi22 := 0; fi22 < 7; fi22++ {
            var v23 uint8 = genFn1(-1)
            v1 = 2
            fmt.Println(v20)
            _ = fi22
            _ = v23
        }
        v1 -= v1
        v20 -= v20
        _ = fi21
    }
    switch "foo bar" {
    case "":
        v4 -= (^(v4))
    }
    var v24 int16 = 1
    fmt.Println(v24, v20, v24)
    if true {
        v24 = 1000
        var v25 int32 = int32(v24)
        v26 := GenS1{F0: v24, F1: -69, F2: (!false)}
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        for fi28 := 0; fi28 < 14; fi28++ {
            var v29 uint64 = 292
            v4 -= v4
            v3[v4] = false
            _ = fi28
            _ = v29
        }
        var v30 float32 = float32((1.5 * 3.14159))
        _ = fi27
        _ = v30
    }
    switch ("a" + "foo bar") {
    case "a":
        _ = len(v2)
    default:
        if (v4 <= (-(v4))) {
            v1 = (^(v1))
            v31 := v3[v4]
            v32 := GenS2{F0: 42, F1: 1, F2: v1}
            _ = v31
            _ = v32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
}
