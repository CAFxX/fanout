package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int; F1 uint64; F2 uint64 }
func genFn0(p0 uint32, p1 int64) uint32 {
    p1 = -9223372036854775808
    v6 := GenS1{F0: 42, F1: ""}
    switch 32767 {
    case -202:
        for fi7 := 0; fi7 < 3; fi7++ {
            if ((p1 * p1) == (p1 | (^(p1)))) {
                p1 = p1
            } else {
                var v8 uint8 = 1
                _ = v8
            }
            _ = fi7
        }
    case 42:
        var v9 uint64 = 2
        _ = v9
    default:
        v10 := []int{127, 128, 9223372036854775807, -2147483648}
        _ = v10
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        fmt.Println(p0, p1)
        if (1000 > 1000) {
            v12 := []int{}
            _ = v12
        }
        _ = fi11
    }
    if (128 >= -396) {
        if (0.0 > 0.1) {
            v13 := []int{-107}
            p1 += p1
            p1 = (p1 + 0)
            _ = v13
        } else {
            var v14 int8 = 127
            _ = v14
        }
        if (780 == 4294967295) {
            p0 = p0
            if (173 >= 0) {
                var v15 uint8 = 128
                p0 = p0
                v16 := []int{-1, int(p1)}
                _ = v15
                _ = v16
            } else {
                v17 := GenS1{F0: -349, F1: "\n\t"}
                _ = v17
            }
        }
        var v18 int8 = 0
        _ = v18
    } else {
        for fi19 := 0; fi19 < 13; fi19++ {
            p0 = p0
            var v20 float64 = (3.14159 + (3.14159 + 0.0))
            _ = fi19
            _ = v20
        }
        p1 += (^(p1))
        v21 := []int{-2147483648}
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v6
    return p0
}

func main() {
    var v1 int32 = 151
    v2 := make(map[uint32]string)
    var v3 bool = false
    var v4 uint64 = 864
    v5 := []uint32{810, 1, 562, 42}
    var v22 int = -9223372036854775808
    v1 += v1
    switch v22 {
    case 127:
        var v23 int = v22
        _ = v23
    case 32767:
        var v24 uint64 = (^(v4))
        _ = v24
    default:
        _ = len(v5)
    }
    var v25 float64 = 0.1
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 int32 = (v1 >> ((18446744073709551615) & 63))
        v25 += (v25 - float64(0.0))
        _ = fi26
        _ = v27
    }
    v28 := GenS1{F0: (-(v1)), F1: "hello"}
    if v3 {
        for fi29 := 0; fi29 < 13; fi29++ {
            v30 := GenS2{F0: v22, F1: v4, F2: v4}
            v31 := GenS2{F0: (v22 * -749), F1: v4, F2: (v4 & v4)}
            var v32 uint16 = 42
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
        v25 = -2.25
    }
    fmt.Println(v25)
    var v33 bool = true
    v4 -= v4
    if v3 {
        for fi34 := 0; fi34 < 8; fi34++ {
            v3 = (!(((1.5 + (-2.25 / 0.1)) * float32(v25)) == -2.25))
            switch "\n\t" {
            case "":
                v33 = false
            case "foo bar":
                v1 -= (v1 << ((2) & 63))
            }
            if (!v3) {
                v35 := v2[42]
                _ = v35
            }
            _ = fi34
        }
        v25 += (v25 - v25)
    } else {
        v5 = append(v5, 2)
        fmt.Println(v22, v3)
        var v36 float32 = (1e10 / 0.1)
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
