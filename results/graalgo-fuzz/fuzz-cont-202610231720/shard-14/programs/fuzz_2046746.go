package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint64; F1 int64; F2 int64 }
func genFn0(p0 bool, p1 uint16, p2 int16) int16 {
    for fi6 := 0; fi6 < 13; fi6++ {
        if (p0 && (p2 < p2)) {
            p1 = p1
            v7 := GenS1{F0: p2, F1: -330}
            if false {
                if (820 < 18446744073709551615) {
                    switch "hello" {
                    case "hello":
                        var v8 float64 = float64(3.14159)
                        _ = v8
                    }
                }
            } else {
                fmt.Println(p2)
                v9 := []int{fi6, (fi6 >> ((65535) & 63)), (^(fi6)), (-(fi6))}
                _ = v9
            }
            _ = v7
        } else {
            var v10 int8 = 1
            _ = v10
        }
        _ = fi6
    }
    var v11 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p2
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]string)
    var v3 int64 = 42
    v4 := make(map[int]string)
    var v5 bool = true
    v1[988] = 2
    fmt.Println(v5, v3)
    v12 := GenS2{F0: 710, F1: 2147483647}
    if v5 {
        var v13 uint = 754
        v14 := v4[2147483647]
        _ = v13
        _ = v14
    }
    var v15 float64 = ((-2.25 * -2.25) + -2.25)
    var v16 uint32 = 1
    v5 = true
    v15 = 3.14159
    for fi17 := 0; fi17 < 10; fi17++ {
        switch 19 {
        case 127:
            for fi18 := 0; fi18 < 17; fi18++ {
                var v19 int64 = v3
                v15 += (float64(0.0) + 0.0)
                _ = fi18
                _ = v19
            }
        case 4294967295:
            var v20 bool = v5
            _ = v20
        default:
            var v21 uint64 = 1
            _ = v21
        }
        _ = fi17
    }
    if v5 {
        fmt.Println(v5)
        var v22 string = "\n\t"
        fmt.Println(v16, v16)
        _ = v22
    }
    v5 = v5
    for fi23 := 0; fi23 < 16; fi23++ {
        v16 = (^(v16))
        var v24 float64 = (v15 - (v15 + (float64(3.14159) + 3.14159)))
        _ = len(v2)
        _ = fi23
        _ = v24
    }
    v25 := GenS3{F0: 1, F1: v3, F2: (v3 + (v3 * (v3 ^ v3)))}
    switch "" {
    case "\n\t":
        for fi26 := 0; fi26 < 9; fi26++ {
            v5 = ((^(fi26)) >= fi26)
            v3 += v3
            v27 := GenS3{F0: 1, F1: (v3 | v3), F2: (v3 >> ((171) & 63))}
            _ = fi26
            _ = v27
        }
    default:
        var v28 int16 = 586
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
