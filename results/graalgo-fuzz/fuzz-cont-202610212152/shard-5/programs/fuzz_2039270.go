package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int8 }
func genFn0() string {
    var v5 uint16 = 255
    var v6 uint = 42
    switch 706 {
    case 9223372036854775807:
        if (v5 > v5) {
            switch 1 {
            case 128:
                v6 -= v6
            case -814:
                var v7 int = 129
                _ = v7
            case 0:
                var v8 uint8 = 1
                _ = v8
            }
            var v9 uint16 = v5
            _ = v9
        }
    default:
        var v10 int64 = -9223372036854775808
        _ = v10
    }
    _ = v5
    _ = v6
    return "\n\t"
}
func genFn1(p0 uint64, p1 uint32) int {
    var v11 uint64 = p0
    if (!true) {
        var v12 uint64 = 2
        v13 := []int{127, -2147483648, -2147483648}
        _ = v12
        _ = v13
    } else {
        var v14 uint16 = 377
        v15 := GenS1{F0: -506, F1: 32767, F2: 42}
        _ = v14
        _ = v15
    }
    fmt.Println(p1, v11, p0)
    _ = p0
    _ = p1
    _ = v11
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = (0.0 + (((1e10 - 1e10) + (float64(0.1) * (1.5 * 3.14159))) - 0.1))
    v3 := make(map[int]int)
    v4 := make(map[uint32]uint32)
    fmt.Println(v2)
    var v16 int8 = -25
    v2 += 1.5
    v2 += v2
    v2 += (v2 + ((v2 + v2) / v2))
    v17 := GenS1{F0: -1, F1: 42, F2: -128}
    var v18 int32 = 0
    if true {
        for fi19 := 0; fi19 < 10; fi19++ {
            for fi20 := 0; fi20 < 11; fi20++ {
                fmt.Println(v16, v16)
                v2 = 1.5
                _ = fi20
            }
            switch "hello" {
            case "\n\t":
                var v21 uint32 = 148
                _ = v21
            case "foo bar":
                var v22 int64 = -803
                _ = v22
            case "":
                var v23 uint64 = 2
                _ = v23
            default:
                var v24 uint16 = 322
                _ = v24
            }
            var v25 bool = false
            _ = fi19
            _ = v25
        }
        var v26 int32 = 0
        var v27 int64 = 42
        _ = v26
        _ = v27
    } else {
        var v28 uint32 = 42
        _ = v28
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        v30 := GenS1{F0: 1, F1: -1, F2: -128}
        _ = fi29
        _ = v30
    }
    var v31 uint64 = 476
    v31 = 0
    if (9223372036854775807 <= 0) {
        v32 := []int{128}
        _ = v32
    }
    var v33 int = 634
    fmt.Println(v31, v2, v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v4
}
