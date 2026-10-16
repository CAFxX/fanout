package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint8; F2 int8 }
func genFn0(p0 bool, p1 int16) int8 {
    var v4 uint32 = 42
    var v5 int64 = -209
    fmt.Println(p0, p1)
    if false {
        if true {
            var v6 uint16 = 1
            for fi7 := 0; fi7 < 12; fi7++ {
                v8 := GenS1{F0: 269, F1: v4}
                _ = fi7
                _ = v8
            }
            _ = v6
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -128
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[int]string)
    var v3 uint64 = 0
    switch 128 {
    case 127:
        if (!false) {
            fmt.Println(v1, v3)
            for fi9 := 0; fi9 < 19; fi9++ {
                v1 -= (^(v1))
                if (2 > (v3 ^ v3)) {
                    v10 := v2[(fi9 - (fi9 ^ (fi9 - -585)))]
                    fmt.Println(v3)
                    _ = v10
                } else {
                    v3 = 281
                }
                _ = fi9
            }
            v11 := GenS1{F0: 42, F1: (v1 / (((v1 >> ((18446744073709551615) & 63))) | 1))}
            _ = v11
        }
    case 32768:
        for fi12 := 0; fi12 < 12; fi12++ {
            fmt.Println(v1)
            var v13 float64 = (1e10 * 0.0)
            var v14 uint8 = 128
            _ = fi12
            _ = v13
            _ = v14
        }
    case 128:
        v15 := GenS2{F0: -1, F1: 42, F2: genFn0(false, 1)}
        _ = v15
    default:
        fmt.Println(v3)
    }
    if false {
        fmt.Println(v3, v3)
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 int64 = -440
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        if false {
            var v19 uint64 = (v3 >> ((18446744073709551615) & 63))
            if true {
                v20 := []int{fi18, 127, fi18, int(v1)}
                fmt.Println(v19, v3, v19)
                if (true && (false || (-128 >= 0))) {
                    var v21 bool = (!true)
                    v21 = (!false)
                    _ = v21
                }
                _ = v20
            } else {
                var v22 int16 = 0
                _ = v22
            }
            var v23 uint16 = 749
            _ = v19
            _ = v23
        }
        _ = fi18
    }
    v1 += (v1 << ((42) & 63))
    var v24 float32 = 1e10
    var v25 float64 = (0.1 * -2.25)
    fmt.Println(v24, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
