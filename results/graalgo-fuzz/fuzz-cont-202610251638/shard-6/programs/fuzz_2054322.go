package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 bool }
type GenS2 struct { F0 int8; F1 int64; F2 int16 }
type GenS3 struct { F0 int; F1 uint32; F2 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int }
func genFn0() int64 {
    var v4 float64 = ((((1.5 - (0.0 * 0.0)) / 1e10) - 1e10) * (0.1 - float64(3.14159)))
    var v5 uint64 = 713
    v6 := []int{-1, 9223372036854775807, 67}
    v4 -= ((1.5 * (v4 + v4)) - (v4 / 1.0))
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    var v3 uint64 = 1
    fmt.Println(v3)
    fmt.Println(v3)
    var v7 uint8 = 255
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            v3 = (v3 / ((0) | 1))
            switch fi9 {
            case 1:
                v7 = v7
            case -1:
                var v10 int16 = -1
                _ = v10
            case -9223372036854775808:
                v3 += 0
            default:
                switch 4294967295 {
                case -2147483648:
                    v11 := v1[4294967295]
                    _ = v11
                case 32767:
                    var v12 uint32 = 42
                    _ = v12
                case 0:
                    var v13 int64 = 0
                    _ = v13
                default:
                    v7 = (v7 << ((255) & 63))
                }
            }
            fmt.Println(fi8, v7)
            if (genFn0() > 42) {
                v14 := GenS1{F0: 1, F1: (774 <= -661), F2: (-2.25 > 1e10)}
                var v15 int = -2147483648
                v3 = 2
                var v16 string = ""
                var v17 uint32 = 2147483648
                v18 := GenS1{F0: 0, F1: false, F2: (uint(v7) <= uint(v7))}
                _ = v14
                _ = v15
                _ = v16
                _ = v17
                _ = v18
            }
            v19 := GenS2{F0: 42, F1: 42, F2: 32767}
            _ = fi9
            _ = v19
        }
        v20 := GenS3{F0: (fi8 >> ((1) & 63)), F1: 42, F2: (v7 <= 9)}
        fmt.Println(v3, v3)
        _ = fi8
        _ = v20
    }
    var v21 uint32 = 748
    if (true || (-2147483648 != 42)) {
        var v22 bool = ((127 != -806) || (!(false || (1 == 617))))
        var v23 uint16 = 358
        fmt.Println(v3, v3, v7)
        _ = v22
        _ = v23
    } else {
        var v24 uint = 9
        delete(v2, 537)
        _ = v24
    }
    v2[2147483647] = 127
    v25 := GenS4{F0: 1}
    fmt.Println(v7, v21)
    if (!((3.14159 / 1.5) >= (0.0 / 3.14159))) {
        var v26 uint32 = 939
        _ = v26
    }
    v3 -= (v3 << ((65535) & 63))
    v27 := GenS5{F0: 621}
    v28 := GenS4{F0: 1}
    v29 := []int{0, 398}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v7
}
