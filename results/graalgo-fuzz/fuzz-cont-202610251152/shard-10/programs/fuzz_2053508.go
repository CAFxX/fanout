package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int32; F1 bool }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 2
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]string)
    v5 := make(map[int]int)
    v6 := []int{-235, -9223372036854775808, 127, 1}
    v7 := v1[0]
    v2 = (v2 << ((42) & 63))
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            v7 = fi8
            _ = fi9
        }
        _ = fi8
    }
    v10 := GenS1{F0: 409, F1: v2, F2: 925}
    for fi11 := 0; fi11 < 13; fi11++ {
        v7 = v7
        v2 -= v2
        _ = fi11
    }
    v12 := GenS2{F0: 821, F1: (255 == 65535)}
    v2 = (^(v2))
    v7 -= v7
    if (2147483647 != -2147483648) {
        v7 = v7
    }
    if true {
        var v13 float32 = (-2.25 / (0.1 - (1e10 * ((1e10 * 3.14159) / 1.0))))
        switch v7 {
        case 9223372036854775807:
            v2 = v2
        default:
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 int64 = 0
                var v16 uint32 = 449
                v16 = (^(v16))
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        _ = v13
    } else {
        v7 += (v7 << ((667) & 63))
        if (!(true || (1e10 < float64(1e10)))) {
            for fi17 := 0; fi17 < 5; fi17++ {
                v18 := GenS2{F0: -2147483648, F1: false}
                if (-1 <= -32768) {
                    fmt.Println(fi17, fi17, fi17)
                    fmt.Println(fi17, v2)
                } else {
                    v7 = fi17
                    fmt.Println(v7)
                    var v19 int32 = 0
                    _ = v19
                }
                _ = fi17
                _ = v18
            }
            var v20 uint32 = 42
            _ = v20
        } else {
            for fi21 := 0; fi21 < 13; fi21++ {
                var v22 uint = 0
                _ = fi21
                _ = v22
            }
        }
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        v2 -= 268
        _ = fi23
    }
    v24 := GenS2{F0: -2147483648, F1: (!(-9223372036854775808 <= -9223372036854775808))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
