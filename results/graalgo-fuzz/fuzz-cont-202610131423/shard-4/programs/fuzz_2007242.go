package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int; F2 int16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int, p1 uint) int8 {
    if (0 == 32767) {
        var v4 int64 = -321
        _ = v4
    }
    switch -1 {
    case 42:
        p1 -= p1
    default:
        v5 := []int{(p0 - (p0 - (p0 * (p0 ^ -9223372036854775808)))), (p0 % (((p0 & p0)) | 1))}
        _ = v5
    }
    _ = p0
    _ = p1
    return -128
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]uint32)
    var v3 bool = (4294967295 != 255)
    v3 = (!true)
    v6 := []int{32767, 32768}
    if false {
        var v7 int8 = 127
        var v8 int = 424
        _ = v7
        _ = v8
    }
    var v9 uint16 = 42
    v10 := GenS1{F0: 18446744073709551615}
    fmt.Println(v1)
    var v11 float64 = (((1e10 / 1.0) * 0.1) * (0.0 * -2.25))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if v3 {
        var v12 uint16 = v9
        var v13 uint16 = (^(v12))
        v13 += 65535
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: 4294967295}
        _ = v14
    }
    v15 := GenS2{F0: 628, F1: 32767, F2: 1000}
    for fi16 := 0; fi16 < 2; fi16++ {
        v6 = append(v6, 2147483647)
        v17 := GenS3{F0: 32767}
        _ = fi16
        _ = v17
    }
    fmt.Println(v9, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
