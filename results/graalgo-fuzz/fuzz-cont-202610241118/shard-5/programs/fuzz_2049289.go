package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint64 }
func genFn0(p0 uint64, p1 int16, p2 uint32) uint16 {
    v4 := []int{897, -9223372036854775808, 0, 0}
    var v5 int32 = 1
    for fi6 := 0; fi6 < 12; fi6++ {
        _ = len(v4)
        var v7 int8 = 100
        _ = fi6
        _ = v7
    }
    v5 = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 774
}
func genFn1() string {
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    var v9 int8 = 0
    switch 42 {
    case 2147483647:
        v9 -= v9
    }
    v10 := GenS1{F0: -2147483648, F1: 111, F2: 53}
    v9 -= v9
    _ = v10
    _ = v9
    return ("a" + ("a" + "a"))
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{-721, 2147483647, 1}
    var v3 float64 = ((((1e10 * 0.0) / -2.25) - (((((0.1 * -2.25) - 0.1) * 1.5) + 1.5) * (float64((3.14159 - (0.1 + float32(-2.25)))) + 0.0))) * 3.14159)
    v3 += (1.5 + v3)
    var v11 int8 = 42
    v3 = ((1.5 + float64(((((0.0 + float32(1e10)) + 0.1) - 3.14159) / (1.5 / 1.0)))) - 1.5)
    var v12 int32 = 1
    fmt.Println(v11, v3)
    switch int(v11) {
    case 4294967295:
        fmt.Println(v11, v12, v12)
    case 452:
        fmt.Println(v12, v11, v12)
    case 384:
        if false {
            switch 802 {
            case 127:
                delete(v1, 974)
            case 32767:
                var v13 uint = 255
                _ = v13
            }
        } else {
            for fi14 := 0; fi14 < 3; fi14++ {
                fmt.Println(fi14, fi14, v11)
                v15 := GenS1{F0: (v12 * 2147483647), F1: 311, F2: (^(fi14))}
                v12 = -649
                _ = fi14
                _ = v15
            }
            var v16 uint = 2
            _ = v16
        }
    default:
        fmt.Println(v11, v11, v3)
    }
    v2 = append(v2, 32768)
    switch 338 {
    case -2147483648:
        if (898 > 939) {
            switch 4294967295 {
            case 772:
                fmt.Println(v12)
            case 9223372036854775807:
                fmt.Println(v12, v12)
            case 814:
                for fi17 := 0; fi17 < 18; fi17++ {
                    v18 := GenS2{F0: -2147483648}
                    _ = fi17
                    _ = v18
                }
            }
            v12 += v12
        }
    default:
        switch 105 {
        case 198:
            v3 -= v3
        case -1:
            fmt.Println(v11, v12, v12)
        }
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 bool = true
        _ = fi19
        _ = v20
    }
    var v21 uint64 = 18446744073709551615
    var v22 int32 = (v12 * v12)
    v23 := GenS3{F0: 251, F1: 1, F2: v21}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
