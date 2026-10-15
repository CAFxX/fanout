package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int16, p1 int32, p2 int) string {
    p1 += (p1 | 42)
    var v5 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "\n\t"
}
func genFn1(p0 bool) uint8 {
    if (0.1 < (float32(0.0) * float32(((0.1 + -2.25) * 3.14159)))) {
        p0 = false
        if false {
            for fi6 := 0; fi6 < 11; fi6++ {
                v7 := []int{-2147483648, -1, fi6, fi6}
                p0 = (4294967295 >= 870)
                _ = fi6
                _ = v7
            }
            for fi8 := 0; fi8 < 3; fi8++ {
                if p0 {
                    var v9 uint64 = 4294967295
                    _ = v9
                }
                p0 = (128 != 0)
                _ = fi8
            }
        }
        if false {
            v10 := []int{127, -9223372036854775808, 504, 0}
            p0 = (!p0)
            _ = v10
        } else {
            if false {
                v11 := GenS1{F0: 0, F1: 280}
                v12 := []int{32768}
                _ = v11
                _ = v12
            }
        }
    } else {
        var v13 int64 = 42
        _ = v13
    }
    v14 := GenS1{F0: 668, F1: 591}
    var v15 uint8 = 128
    _ = p0
    _ = v14
    _ = v15
    return v15
}

func main() {
    var v1 uint = 2
    var v2 uint = v1
    var v3 int64 = 1
    var v4 int = 2147483647
    switch (-(v4)) {
    case 4294967295:
        fmt.Println(v4, v2)
    case 615:
        var v16 int16 = -509
        _ = v16
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        v2 += 255
        v18 := []int{fi17, fi17}
        var v19 int64 = (v3 >> (((v1 % ((975) | 1))) & 63))
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := []int{}
    var v21 uint64 = 2
    v3 += 1
    v2 -= 2
    var v22 string = "\n\t"
    for fi23 := 0; fi23 < 17; fi23++ {
        if len(v20) > 0 {
            v20[0] = v4
        }
        var v24 int64 = 42
        if (!false) {
            _ = len(v20)
        }
        _ = fi23
        _ = v24
    }
    v20 = append(v20, v4)
    var v25 int64 = -1
    if (676 != 660) {
        switch "" {
        case "\n\t":
            fmt.Println(v3, v1)
        case "foo bar":
            var v26 float32 = -2.25
            _ = v26
        case "a":
            var v27 int8 = -36
            _ = v27
        }
    }
    switch ("\n\t" + v22) {
    case "foo bar":
        if (-2147483648 < -951) {
            v28 := GenS2{F0: 12}
            _ = v28
        }
    case "\n\t":
        switch (v4 % ((32768) | 1)) {
        case 32768:
            fmt.Println(v3, v4)
        default:
            v4 = v4
        }
    case "a":
        fmt.Println(v1, v25)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
}
