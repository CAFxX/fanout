package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int8 }
func genFn0(p0 int8, p1 uint) int8 {
    var v6 uint8 = 127
    v7 := GenS1{F0: 9223372036854775807, F1: p1}
    for fi8 := 0; fi8 < 19; fi8++ {
        fmt.Println(fi8, fi8)
        p0 = p0
        fmt.Println(p0, v6, p0)
        _ = fi8
    }
    if true {
        v9 := GenS1{F0: 127, F1: 65535}
        var v10 uint16 = 32767
        v10 -= v10
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 127
}

func main() {
    var v1 float32 = (0.1 * (0.1 * 1e10))
    v2 := make(map[string]bool)
    var v3 int8 = 1
    v4 := make(map[uint32]bool)
    var v5 uint32 = 65535
    switch "" {
    case "hello":
        var v11 uint32 = 1
        _ = v11
    case "":
        v12 := []int{42}
        _ = v12
    case "foo bar":
        var v13 int32 = 0
        _ = v13
    default:
        if (-1 <= -233) {
            if (!true) {
                if (4294967295 >= 2) {
                    var v14 float64 = (1e10 / 1.0)
                    switch "hello" {
                    case "":
                        v14 += v14
                    case "hello":
                        v2["hello"] = (1 != 1)
                    }
                    v3 += int8(v5)
                    _ = v14
                }
                var v15 uint8 = 2
                fmt.Println(v15, v5)
                _ = v15
            }
        }
    }
    if (308 == 0) {
        fmt.Println(v1, v3)
        v16 := GenS1{F0: 9223372036854775807, F1: 255}
        _ = v16
    } else {
        fmt.Println(v3)
    }
    var v17 uint = 2
    v18 := GenS2{F0: genFn0((v3 & v3), 609)}
    v19 := []int{374, 527, 32767, 0}
    var v20 uint = v17
    if len(v19) > 0 {
        v19[0] = 4294967295
    }
    switch -139 {
    case 4294967295:
        v21 := v4[v5]
        _ = v21
    case 32767:
        if len(v19) > 0 {
            v19[0] = -2147483648
        }
    case -1:
        if (32767 < 1) {
            v20 = (v17 * (v20 ^ v17))
        }
    }
    var v22 uint8 = 128
    var v23 bool = false
    v5 -= uint32(v20)
    if v23 {
        _ = len(v4)
        for fi24 := 0; fi24 < 9; fi24++ {
            for fi25 := 0; fi25 < 4; fi25++ {
                fmt.Println(v20, v20, v5)
                _ = fi25
            }
            var v26 int = (-(fi24))
            _ = fi24
            _ = v26
        }
    } else {
        var v27 uint16 = 32767
        fmt.Println(v23, v27)
        if false {
            v28 := GenS2{F0: (v3 >> (((^(v17))) & 63))}
            _ = v28
        } else {
            var v29 int = 763
            _ = v29
        }
        _ = v27
    }
    delete(v4, (v5 * 559))
    if (!(v23 || (v20 != v17))) {
        v3 -= v3
        var v30 int64 = -327
        v3 -= genFn0((v3 % ((1) | 1)), 255)
        _ = v30
    } else {
        if (!true) {
            for fi31 := 0; fi31 < 15; fi31++ {
                var v32 int8 = 127
                v1 += (float32(float64(v1)) / v1)
                _ = fi31
                _ = v32
            }
            fmt.Println(v17, v1, v20)
            v3 = (-(v3))
        } else {
            v17 = (v17 >> (((v20 / ((1) | 1))) & 63))
            fmt.Println(v1)
            _ = len(v2)
        }
        _ = len(v4)
        v22 = v22
    }
    var v33 float64 = (-2.25 / 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
