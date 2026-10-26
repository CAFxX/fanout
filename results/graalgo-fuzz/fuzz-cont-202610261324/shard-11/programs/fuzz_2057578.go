package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int16; F1 uint64; F2 uint }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []uint32{65535, 1, 759}
    var v2 string = ((((("\n\t" + "") + "") + "\n\t") + "foo bar") + "hello")
    var v3 bool = ((141 == 1) && (2147483647 < 4294967295))
    v4 := []int{127, 1, 2147483647}
    var v5 int32 = 0
    var v6 int64 = 1
    v6 += (v6 >> ((2) & 63))
    var v7 uint8 = 1
    v8 := GenS1{F0: v2, F1: 9223372036854775807}
    switch 42 {
    case 32768:
        var v9 int8 = -1
        _ = v9
    case 2147483647:
        if v3 {
            v4 = append(v4, -1)
            fmt.Println(v3, v2, v7)
            v10 := GenS2{F0: 0, F1: 42, F2: 4294967295}
            _ = v10
        }
    case 0:
        v11 := GenS3{F0: v5}
        _ = v11
    default:
        _ = len(v1)
    }
    v5 = (v5 % (((v5 | 1)) | 1))
    v7 = (v7 ^ (v7 - 239))
    var v12 float32 = -2.25
    var v13 float32 = float32(float64(v12))
    var v14 int8 = 101
    v6 += (v6 << ((4294967295) & 63))
    for fi15 := 0; fi15 < 17; fi15++ {
        if len(v4) > 0 {
            v4[0] = fi15
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            switch fi16 {
            case -722:
                var v17 uint8 = v7
                _ = v17
            case 42:
                v13 = (0.1 * 1e10)
            }
            fmt.Println(fi16, fi16)
            if (!(!true)) {
                v18 := GenS2{F0: 1000, F1: 346, F2: 65535}
                v7 = (^(v7))
                var v19 uint64 = 4294967295
                _ = v18
                _ = v19
            }
            _ = fi16
        }
        _ = fi15
    }
    v14 -= (v14 / ((v14) | 1))
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    switch v2 {
    case "":
        for fi20 := 0; fi20 < 3; fi20++ {
            switch fi20 {
            case 0:
                var v21 int64 = v6
                _ = v21
            case 42:
                _ = len(v1)
            case -1:
                var v22 bool = true
                _ = v22
            }
            v6 += (^(v6))
            _ = fi20
        }
    default:
        v5 = 91
    }
    v6 = 650
    v5 = 784
    for fi23 := 0; fi23 < 3; fi23++ {
        for fi24 := 0; fi24 < 16; fi24++ {
            v13 += v13
            _ = fi24
        }
        v25 := GenS2{F0: 1000, F1: 18446744073709551615, F2: 344}
        _ = fi23
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
