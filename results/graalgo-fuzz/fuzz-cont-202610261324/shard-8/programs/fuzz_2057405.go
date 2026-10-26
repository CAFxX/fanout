package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 string; F1 int; F2 uint8 }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint }
func genFn0() uint8 {
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := []int{fi7, fi7, fi7, (fi7 & fi7)}
        var v9 uint64 = 399
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 2147483647
    if false {
        if false {
            for fi11 := 0; fi11 < 10; fi11++ {
                fmt.Println(v10)
                _ = fi11
            }
            v12 := GenS1{F0: (266 != 611), F1: (^(v10)), F2: 296}
            _ = v12
        }
        switch 740 {
        case -472:
            v13 := []int{-9223372036854775808, 32767, 2147483647, 630}
            _ = v13
        default:
            fmt.Println(v10)
        }
    } else {
        switch -9223372036854775808 {
        case 32767:
            v14 := GenS1{F0: ((float32((((-2.25 - (3.14159 + (1.5 / float64(1.5)))) - -2.25) - 3.14159)) * float32(float64(((1.5 * 0.0) + (float32(3.14159) + 3.14159))))) == (0.0 + (3.14159 + (0.0 - (((0.0 / ((3.14159 + (0.0 - (((-2.25 / 1.5) / 1e10) - (1e10 * 0.1)))) * (1.5 / (1e10 / float32((0.1 * -2.25)))))) - 3.14159) + (1.5 / (3.14159 * float32(3.14159)))))))), F1: 2, F2: 128}
            _ = v14
        case -1:
            if true {
                v15 := GenS1{F0: true, F1: (v10 + v10), F2: -2147483648}
                _ = v15
            }
        case 0:
            for fi16 := 0; fi16 < 4; fi16++ {
                v10 = v10
                _ = fi16
            }
        }
        var v17 bool = (-32768 <= 32767)
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        v10 += v10
        _ = fi18
    }
    v10 = v10
    _ = v10
    return 128
}
func genFn1(p0 int64, p1 int, p2 int32) uint {
    p0 += (p0 * 42)
    var v19 int64 = p0
    for fi20 := 0; fi20 < 18; fi20++ {
        fmt.Println(v19, p0)
        _ = fi20
    }
    p1 = (p1 / ((127) | 1))
    var v21 int = (p1 >> ((4294967295) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v21
    return 65535
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    v4 := []uint32{65535, 2147483647, 467}
    v5 := []string{}
    var v6 int64 = 1
    v6 += v6
    v6 = v6
    var v22 uint32 = 2
    fmt.Println(v6, v22)
    v2[v22] = v22
    v22 = v22
    fmt.Println(v6, v22)
    switch "" {
    case "hello":
        if (((^(v22)) < v22) || (0.1 != (1e10 * (0.1 * 3.14159)))) {
            switch 32768 {
            case 32768:
                var v23 uint = genFn1(v6, 315, -2147483648)
                _ = v23
            case -9223372036854775808:
                v24 := GenS2{F0: 754, F1: 42}
                _ = v24
            default:
                var v25 bool = true
                _ = v25
            }
        } else {
            var v26 uint16 = 0
            if true {
                var v27 int16 = 1000
                _ = v27
            }
            var v28 int8 = 81
            _ = v26
            _ = v28
        }
    case "a":
        v29 := GenS3{F0: "", F1: -1, F2: 0}
        _ = v29
    default:
        var v30 float32 = float32(-2.25)
        _ = v30
    }
    fmt.Println(v6, v6)
    var v31 uint = 1
    v32 := GenS4{F0: 128, F1: v22, F2: v31}
    var v33 int16 = -32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
