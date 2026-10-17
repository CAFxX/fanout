package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int8, p1 uint64, p2 int32) int16 {
    var v6 int32 = (p2 & (p2 / (((p2 ^ p2)) | 1)))
    var v7 int16 = -798
    var v8 int16 = v7
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return v8
}
func genFn1(p0 int64, p1 uint, p2 int) uint64 {
    var v9 string = ("a" + ("\n\t" + ("" + ("foo bar" + "\n\t"))))
    p1 += 0
    switch 0 {
    case -234:
        var v10 uint = (p1 & p1)
        _ = v10
    case 4294967295:
        v11 := GenS1{F0: 2147483647, F1: 42, F2: p1}
        _ = v11
    default:
        v12 := []int{(-(p2)), p2, p2}
        _ = v12
    }
    fmt.Println(v9, p0)
    v13 := GenS2{F0: int32(p0), F1: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v9
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = -9223372036854775808
    var v3 int64 = (v2 | v2)
    var v4 uint64 = 1
    var v5 float32 = 1.5
    for fi14 := 0; fi14 < 3; fi14++ {
        v2 -= (v3 >> ((857) & 63))
        v3 -= v2
        _ = fi14
    }
    v15 := GenS1{F0: -926, F1: (v2 % (((v2 | 42)) | 1)), F2: 42}
    v16 := GenS1{F0: 42, F1: v2, F2: 351}
    for fi17 := 0; fi17 < 11; fi17++ {
        _ = len(v1)
        _ = fi17
    }
    var v18 uint32 = 2147483647
    v19 := []int{}
    fmt.Println(v5, v5, v18)
    if (140 != 217) {
        fmt.Println(v5, v4, v2)
        switch 0 {
        case -994:
            v20 := GenS2{F0: -1, F1: 855}
            _ = v20
        default:
            for fi21 := 0; fi21 < 10; fi21++ {
                v1[(fi21 - 9223372036854775807)] = v18
                v2 -= v3
                v22 := GenS3{F0: 128, F1: 449}
                _ = fi21
                _ = v22
            }
        }
    } else {
        fmt.Println(v5)
    }
    v18 = v18
    var v23 int8 = 1
    fmt.Println(v23, v18, v3)
    fmt.Println(v4, v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
