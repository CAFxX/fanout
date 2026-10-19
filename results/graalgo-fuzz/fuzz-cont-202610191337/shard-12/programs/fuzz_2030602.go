package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 bool }
func genFn0(p0 bool) int64 {
    v5 := GenS1{F0: 2, F1: 995, F2: true}
    for fi6 := 0; fi6 < 16; fi6++ {
        fmt.Println(fi6, fi6)
        var v7 int = (fi6 & (fi6 << ((973) & 63)))
        v8 := []int{1, fi6}
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v5
    return -1
}
func genFn1(p0 uint32, p1 uint, p2 int8) uint16 {
    var v9 float32 = 1.5
    v10 := GenS1{F0: 42, F1: -1, F2: true}
    for fi11 := 0; fi11 < 15; fi11++ {
        p0 = (^(p0))
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 193
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    var v12 uint8 = 255
    v13 := GenS1{F0: 498, F1: -1, F2: (4294967295 > 0)}
    var v14 uint16 = 766
    switch -655 {
    case 42:
        fmt.Println(v14, v12)
    case 32768:
        fmt.Println(v14)
    case -2147483648:
        switch 361 {
        case -9223372036854775808:
            for fi15 := 0; fi15 < 3; fi15++ {
                v16 := []int{799, (fi15 / ((int(v14)) | 1)), fi15}
                for fi17 := 0; fi17 < 10; fi17++ {
                    v18 := GenS1{F0: 18446744073709551615, F1: -2147483648, F2: false}
                    var v19 int32 = int32(fi15)
                    fmt.Println(v14, v12, fi15)
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                fmt.Println(fi15, v12)
                _ = fi15
                _ = v16
            }
        default:
            var v20 int64 = -9223372036854775808
            _ = v20
        }
    default:
        v14 -= (v14 + (v14 / ((v14) | 1)))
    }
    v14 = 384
    var v21 float64 = 3.14159
    v4[128] = -1
    switch -1 {
    case -684:
        var v22 float32 = (0.0 - 1.5)
        _ = v22
    }
    var v23 uint16 = v14
    v14 -= genFn1(165, 0, 100)
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := v1["a"]
        var v26 float64 = -2.25
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
}
