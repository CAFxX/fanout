package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int16 }
type GenS2 struct { F0 int32; F1 uint16 }
func genFn0() bool {
    var v6 int32 = 2147483647
    fmt.Println(v6)
    v6 -= (v6 & v6)
    _ = v6
    return true
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{86, 608}
    var v3 uint32 = 432
    var v4 uint8 = 128
    var v5 int32 = 1
    var v7 int16 = 32767
    var v8 bool = (128 == (v4 / ((v4) | 1)))
    var v9 uint8 = (v4 - (v4 | v4))
    v10 := GenS1{F0: 428, F1: v9, F2: v7}
    v11 := GenS1{F0: 4294967295, F1: v9, F2: -1}
    switch -930 {
    case 849:
        v7 += 1
    case 364:
        v5 = (v5 << ((uint(v9)) & 63))
    case 1:
        v4 = v4
    }
    v12 := GenS1{F0: uint(v4), F1: (v4 | v9), F2: (v7 ^ v7)}
    var v13 uint32 = (v3 | v3)
    v1[-299] = 2
    _ = len(v2)
    if v8 {
        var v14 float64 = ((3.14159 + 1e10) - (0.1 * 0.1))
        if v8 {
            v7 -= int16(v5)
        } else {
            var v15 uint = 987
            for fi16 := 0; fi16 < 5; fi16++ {
                v17 := GenS1{F0: (v15 >> ((v15) & 63)), F1: (v9 / ((v9) | 1)), F2: (v7 / (((^(v7))) | 1))}
                if len(v2) > 0 {
                    v2[0] = v13
                }
                switch (fi16 | int(v4)) {
                case -1:
                    v4 = v9
                case 783:
                    var v18 uint = 625
                    _ = v18
                case 2147483647:
                    v19 := GenS2{F0: v5, F1: 415}
                    _ = v19
                default:
                    var v20 int32 = (-(v5))
                    _ = v20
                }
                _ = fi16
                _ = v17
            }
            v8 = (-9223372036854775808 < int64(v15))
            _ = v15
        }
        _ = len(v1)
        _ = v14
    } else {
        fmt.Println(v13)
        v7 = (v7 + (v7 % (((^(v7))) | 1)))
        v5 = v5
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
