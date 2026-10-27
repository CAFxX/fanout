package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int8 }
type GenS2 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint64, p1 uint16, p2 int) uint64 {
    var v4 uint32 = 533
    v5 := []int{1, (p2 - -2147483648), -2147483648}
    var v6 float32 = 1.5
    v7 := GenS1{F0: (^(v4)), F1: 2, F2: 72}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (p0 * p0)
}
func genFn1(p0 uint32, p1 uint) uint {
    if false {
        fmt.Println(p1, p0)
        var v8 int32 = 0
        _ = v8
    }
    v9 := []int{32768, 24, 4294967295}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = v9
    return p1
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[uint32]string)
    var v3 bool = (1 <= 1)
    var v10 int64 = -1
    v10 -= (-(v10))
    var v11 float64 = 1e10
    v12 := GenS1{F0: 2, F1: 65535, F2: -12}
    var v13 float32 = (0.0 + 0.1)
    if v3 {
        var v14 int32 = 2147483647
        switch "\n\t" {
        case "hello":
            switch "foo bar" {
            case "hello":
                v15 := GenS1{F0: 854, F1: 988, F2: int8(v1)}
                _ = v15
            default:
                var v16 bool = (float32(v11) >= (v13 / (float32(0.0) + v13)))
                _ = v16
            }
        case "foo bar":
            v1 -= (v1 - (^(v1)))
        case "":
            v3 = false
        default:
            if (v3 || (!(42 < 83))) {
                _ = len(v2)
            }
        }
        if (1 != 2) {
            var v17 float32 = 1.5
            for fi18 := 0; fi18 < 18; fi18++ {
                v3 = ((((v11 * (v11 + 0.0)) / -2.25) + ((float64(v13) / (-2.25 / v11)) - ((v11 * ((v11 - 0.1) + v11)) / (float64(0.1) + (v11 - v11))))) != -2.25)
                _ = fi18
            }
            v1 = v1
            _ = v17
        } else {
            fmt.Println(v10, v3, v11)
        }
        _ = v14
    } else {
        v19 := GenS1{F0: 1, F1: 2, F2: 0}
        _ = v19
    }
    v20 := []int{-605}
    for fi21 := 0; fi21 < 20; fi21++ {
        v22 := GenS2{F0: 4294967295, F1: 476}
        _ = fi21
        _ = v22
    }
    if (-128 <= 127) {
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 uint8 = 127
            var v25 int = (fi23 % ((1) | 1))
            _ = fi23
            _ = v24
            _ = v25
        }
        v26 := GenS1{F0: 42, F1: 417, F2: 84}
        _ = v26
    }
    fmt.Println(v1, v10, v11)
    switch -9223372036854775808 {
    case 841:
        var v27 uint8 = uint8(v1)
        _ = v27
    }
    fmt.Println(v13)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
}
