package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int64 }
type GenS2 struct { F0 int32; F1 string; F2 uint32 }
func genFn0(p0 uint, p1 uint16, p2 uint32) string {
    if (false && ((1 != -1) && (false && (("\n\t" + "hello") != (("hello" + "a") + ("hello" + "hello")))))) {
        p2 = (p2 / ((p2) | 1))
        v5 := GenS1{F0: 69, F1: (1 > -81), F2: int64(p2)}
        _ = v5
    } else {
        if false {
            v6 := []int{-996, -9223372036854775808}
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            p0 = 4294967295
            _ = v6
        }
        fmt.Println(p1, p2, p0)
    }
    fmt.Println(p1)
    var v7 uint64 = 4294967295
    v8 := []int{0, -8, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    var v1 int64 = 3
    var v2 int32 = 784
    var v3 int64 = (v1 | (v1 & v1))
    v4 := make(map[uint32]uint32)
    var v9 uint16 = 642
    switch 0 {
    case 128:
        var v10 bool = (32767 < 42)
        _ = v10
    }
    v9 = (^(v9))
    v2 += (-(v2))
    v11 := GenS1{F0: 127, F1: (4294967295 >= 724), F2: 0}
    switch 0 {
    case -205:
        v3 += (v1 % ((v1) | 1))
    case -9223372036854775808:
        if false {
            var v12 int16 = -32768
            _ = v12
        }
    case -1:
        var v13 int16 = -1
        _ = v13
    default:
        v14 := []int{-9223372036854775808, -95, -1, 4294967295}
        _ = v14
    }
    var v15 uint16 = 65535
    v2 += (v2 + int32(v1))
    switch "" {
    case "\n\t":
        switch 32767 {
        case 2147483647:
            var v16 float64 = 0.0
            _ = v16
        case -567:
            var v17 uint8 = 2
            _ = v17
        case 448:
            switch -1 {
            case 127:
                v18 := GenS1{F0: 128, F1: (127 < -1), F2: 417}
                _ = v18
            }
        }
    case "a":
        fmt.Println(v9, v9, v3)
    default:
        v3 = v1
    }
    v19 := GenS2{F0: 893, F1: "\n\t", F2: 65535}
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 uint32 = 2147483648
        _ = fi20
        _ = v21
    }
    v15 -= v9
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
