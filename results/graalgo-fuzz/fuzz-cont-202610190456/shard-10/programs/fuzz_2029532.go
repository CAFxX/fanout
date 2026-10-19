package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint16 }
func genFn0() int32 {
    switch -9223372036854775808 {
    case 128:
        v4 := GenS1{F0: 42, F1: 845}
        _ = v4
    case 127:
        switch 741 {
        case 32768:
            var v5 bool = ((42 != -128) || true)
            _ = v5
        case 2147483647:
            var v6 float64 = (1e10 + 0.0)
            _ = v6
        default:
            var v7 uint = 888
            _ = v7
        }
    default:
        var v8 uint8 = 2
        _ = v8
    }
    switch "\n\t" {
    case "\n\t":
        switch "\n\t" {
        case "foo bar":
            var v9 int16 = 474
            _ = v9
        default:
            v10 := GenS2{F0: 2}
            _ = v10
        }
    case "a":
        switch "" {
        case "a":
            var v11 int64 = 862
            _ = v11
        }
    }
    return 0
}
func genFn1(p0 int32) int32 {
    var v12 int8 = 67
    v13 := []int{930, -9223372036854775808, -293, 700}
    switch 128 {
    case -9223372036854775808:
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS1{F0: -152, F1: (fi14 % ((0) | 1))}
            v13 = append(v13, 4294967295)
            var v16 uint32 = 781
            _ = fi14
            _ = v15
            _ = v16
        }
    case 127:
        switch "\n\t" {
        case "\n\t":
            var v17 int8 = (v12 % (((^(v12))) | 1))
            _ = v17
        case "hello":
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        default:
            v18 := GenS2{F0: 319}
            _ = v18
        }
    case 128:
        fmt.Println(v12, v12)
    default:
        var v19 int32 = genFn0()
        _ = v19
    }
    v20 := GenS2{F0: 0}
    _ = p0
    _ = v12
    _ = v13
    _ = v20
    return (^(p0))
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    v3 := []uint32{}
    var v21 bool = (0 != -32768)
    var v22 int32 = genFn1(-1)
    var v23 int16 = 42
    var v24 int16 = (v23 % ((1000) | 1))
    for fi25 := 0; fi25 < 12; fi25++ {
        v23 -= (v24 * v23)
        switch "a" {
        case "":
            fmt.Println(v22)
        default:
            for fi26 := 0; fi26 < 6; fi26++ {
                fmt.Println(v24)
                _ = fi26
            }
        }
        if true {
            var v27 uint8 = 127
            v28 := GenS1{F0: -9223372036854775808, F1: (-(fi25))}
            fmt.Println(v21, v21, v22)
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    v24 -= 527
    var v29 uint8 = 128
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
}
