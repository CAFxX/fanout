package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint16 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int64, p1 uint8, p2 int) int16 {
    v4 := GenS1{F0: (p0 + (p0 % ((9223372036854775807) | 1))), F1: (p2 ^ (p2 << ((42) & 63))), F2: 255}
    var v5 uint32 = 655
    v6 := GenS2{F0: 1, F1: 217, F2: 210}
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 float64 = (0.0 / (3.14159 / 1.5))
        var v9 int = 415
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint, p1 int16, p2 int) uint64 {
    p0 -= (^(p0))
    v10 := []int{p2, (p2 % ((p2) | 1)), p2, (p2 | p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}

func main() {
    var v1 uint8 = 0
    var v2 uint = 42
    v3 := make(map[string]int)
    v3["hello"] = 42
    v11 := v3[""]
    if ((!true) && true) {
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS1{F0: -857, F1: 128, F2: 2}
            if true {
                var v14 string = "\n\t"
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        var v15 uint64 = 995
        _ = v15
    } else {
        v16 := GenS2{F0: 65535, F1: (^(v1)), F2: 213}
        _ = v16
    }
    switch ((((("foo bar" + "hello") + "foo bar") + ("foo bar" + "hello")) + "") + "foo bar") {
    case "":
        var v17 uint32 = 65535
        _ = v17
    case "a":
        if (-960 >= genFn0(993, (v1 | (v1 << ((42) & 63))), 2147483647)) {
            _ = len(v3)
        }
    case "hello":
        if false {
            if (false && false) {
                v18 := GenS1{F0: 42, F1: v11, F2: 255}
                for fi19 := 0; fi19 < 4; fi19++ {
                    v1 = v1
                    v1 -= uint8(v2)
                    _ = fi19
                }
                fmt.Println(v11, v1)
                _ = v18
            } else {
                var v20 uint16 = 255
                _ = v20
            }
        }
    default:
        v21 := GenS3{F0: 1, F1: 375}
        _ = v21
    }
    v22 := GenS1{F0: 42, F1: v11, F2: 2}
    v23 := GenS4{F0: v2}
    v24 := []int{(v11 + (^(v11))), (v11 & v11), (v11 + v11), 1}
    v25 := GenS5{F0: -9223372036854775808, F1: 821}
    for fi26 := 0; fi26 < 14; fi26++ {
        for fi27 := 0; fi27 < 8; fi27++ {
            v28 := GenS4{F0: (v2 << ((v2) & 63))}
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    if (!(v1 < uint8(v2))) {
        var v29 int = 9223372036854775807
        v30 := GenS2{F0: v2, F1: 1, F2: 0}
        v31 := GenS4{F0: (v2 - (^(v2)))}
        _ = v29
        _ = v30
        _ = v31
    } else {
        fmt.Println(v1, v11)
        v32 := GenS4{F0: v2}
        _ = len(v24)
        _ = v32
    }
    if (!((v2 / ((v2) | 1)) <= v2)) {
        var v33 int8 = -1
        v34 := GenS6{F0: 2, F1: v1}
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
