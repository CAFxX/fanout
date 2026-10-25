package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int32 }
func genFn0(p0 int16, p1 uint) uint {
    switch 0 {
    case 0:
        var v5 int64 = -778
        _ = v5
    default:
        p0 -= (p0 % (((-(p0))) | 1))
    }
    v6 := []int{893, 128, 127, 32768}
    for fi7 := 0; fi7 < 2; fi7++ {
        p1 -= p1
        _ = fi7
    }
    var v8 float64 = ((0.1 + 0.1) * float64(3.14159))
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 65535
}
func genFn1(p0 uint, p1 uint32) int64 {
    var v9 int8 = -1
    v10 := []int{-9223372036854775808, 0, 4294967295}
    var v11 uint16 = 311
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return -1
}

func main() {
    var v1 uint16 = 2
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]uint32)
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS1{F0: 42, F1: -89, F2: -2147483648}
        v14 := GenS1{F0: -495, F1: 42, F2: 0}
        v1 = (v1 | (^(v1)))
        _ = fi12
        _ = v13
        _ = v14
    }
    v1 = v1
    switch 2147483647 {
    case 0:
        fmt.Println(v1)
    }
    var v15 int16 = -32768
    if (genFn0(v15, 65535) <= 18446744073709551615) {
        var v16 bool = (1e10 > 1e10)
        fmt.Println(v15)
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            fmt.Println(fi18, v1, fi17)
            var v19 uint = 839
            var v20 string = "a"
            v15 -= (v15 | -1)
            v1 += 228
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    v21 := GenS1{F0: -1, F1: 42, F2: -704}
    switch "a" {
    case "\n\t":
        if false {
            switch 32768 {
            case 32768:
                v22 := v3[2]
                _ = v22
            }
        } else {
            if len(v2) > 0 {
                v2[0] = 42
            }
        }
    case "":
        if true {
            v23 := GenS1{F0: 1, F1: 42, F2: 0}
            _ = v23
        }
    case "foo bar":
        var v24 int16 = 1000
        _ = v24
    default:
        var v25 int8 = 25
        _ = v25
    }
    var v26 bool = ((71 > 0) && true)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v4
}
