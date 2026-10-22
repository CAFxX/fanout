package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint, p1 int16, p2 int64) int16 {
    fmt.Println(p2)
    if false {
        var v7 uint = 18446744073709551615
        p2 = (p2 >> (((^(v7))) & 63))
        _ = v7
    }
    p2 += (p2 % ((p2) | 1))
    _ = p0
    _ = p1
    _ = p2
    return p1
}
func genFn1(p0 int8, p1 uint64, p2 uint32) int {
    v8 := GenS1{F0: p0, F1: 0}
    if (false || (3.14159 > 0.1)) {
        fmt.Println(p2)
        p1 = 42
    }
    var v9 bool = true
    var v10 float64 = (0.0 - 1.5)
    var v11 uint16 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 466
    v3 := make(map[int]int)
    var v4 int = 827
    var v5 int64 = 26
    var v6 uint16 = 65535
    v12 := GenS1{F0: 0, F1: v6}
    if false {
        if false {
            v2 = (v2 % (((v2 + v2)) | 1))
        }
    }
    var v13 uint = 0
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 int16 = genFn0(v13, (-(v2)), v5)
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: -100, F1: (v6 << ((v13) & 63))}
    var v17 bool = (!(1.5 <= ((1e10 / 1e10) - 0.1)))
    switch "a" {
    case "foo bar":
        v18 := []int{}
        _ = v18
    default:
        var v19 uint16 = (v6 >> ((4294967295) & 63))
        _ = v19
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        for fi21 := 0; fi21 < 12; fi21++ {
            if v17 {
                fmt.Println(v6, fi21, v13)
                v22 := GenS1{F0: 1, F1: 32767}
                _ = v22
            } else {
                v23 := GenS1{F0: 15, F1: v6}
                var v24 uint32 = uint32(v6)
                v6 = 255
                _ = v23
                _ = v24
            }
            v25 := []int{(fi20 ^ fi21), v4}
            v26 := GenS2{F0: -128}
            v4 += genFn1(102, 0, 2)
            _ = len(v25)
            var v27 int16 = -32768
            v13 = 380
            v28 := v1[genFn1(66, 42, 1)]
            v29 := GenS3{F0: (v5 << (((^(v13))) & 63))}
            fmt.Println(v27, v5, v4)
            _ = fi21
            _ = v25
            _ = v26
            _ = v27
            _ = v28
            _ = v29
        }
        _ = fi20
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        v6 += (v6 | 0)
        _ = fi30
    }
    var v31 string = ("" + "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
