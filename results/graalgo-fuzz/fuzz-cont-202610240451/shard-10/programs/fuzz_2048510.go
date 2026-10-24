package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint16; F2 int64 }
func genFn0(p0 uint8, p1 bool) int64 {
    p0 = p0
    if (255 == 255) {
        v6 := GenS1{F0: -99, F1: 255}
        v7 := GenS1{F0: -1, F1: 4294967295}
        if p1 {
            for fi8 := 0; fi8 < 13; fi8++ {
                p1 = true
                _ = fi8
            }
        }
        _ = v6
        _ = v7
    } else {
        for fi9 := 0; fi9 < 12; fi9++ {
            if false {
                var v10 int16 = -32768
                var v11 uint16 = 750
                _ = v10
                _ = v11
            } else {
                p0 = 109
                v12 := []int{-1}
                _ = v12
            }
            v13 := GenS2{F0: 1, F1: 2147483648}
            _ = fi9
            _ = v13
        }
        fmt.Println(p1, p1)
        var v14 int8 = 0
        _ = v14
    }
    fmt.Println(p1, p0)
    if ((float32((3.14159 - float64(0.0))) > 3.14159) && p1) {
        var v15 int16 = -1
        switch "\n\t" {
        case "\n\t":
            v15 = (v15 / (((v15 + 1)) | 1))
        case "a":
            fmt.Println(p0, p1)
        }
        _ = v15
    } else {
        p0 = p0
        var v16 bool = (p1 || (1 != uint32(p0)))
        _ = v16
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int64 = 0
    var v3 uint64 = 1
    var v4 int8 = -125
    v5 := []uint32{(v1 % ((65535) | 1)), v1, (v1 << ((0) & 63)), v1}
    v17 := GenS2{F0: 255, F1: 4294967295}
    v2 = -1
    var v18 int = -11
    for fi19 := 0; fi19 < 9; fi19++ {
        fmt.Println(v2)
        var v20 uint8 = 128
        if true {
            _ = len(v5)
            var v21 uint = 4294967295
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    if len(v5) > 0 {
        v5[0] = v1
    }
    v22 := GenS3{F0: 9223372036854775807, F1: 32767, F2: v2}
    switch int(v4) {
    case -9223372036854775808:
        fmt.Println(v18, v4, v2)
    case 128:
        if len(v5) > 0 {
            v5[0] = v1
        }
    case 42:
        for fi23 := 0; fi23 < 6; fi23++ {
            fmt.Println(v3)
            v24 := GenS2{F0: 42, F1: 4294967295}
            v1 += v1
            _ = fi23
            _ = v24
        }
    }
    v5 = append(v5, 0)
    fmt.Println(v18, v18)
    switch v18 {
    case -441:
        var v25 uint32 = 42
        _ = v25
    default:
        if false {
            if false {
                v1 -= 42
                fmt.Println(v18)
            }
        }
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
