package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16 }
func genFn0() uint8 {
    v6 := GenS1{F0: -632, F1: 623}
    var v7 bool = (167 > 42)
    v7 = (!(127 < 42))
    switch -9223372036854775808 {
    case 599:
        v8 := []int{792}
        _ = v8
    case 752:
        switch ("hello" + "foo bar") {
        case "\n\t":
            var v9 float64 = 3.14159
            _ = v9
        case "":
            var v10 bool = (0 <= 54)
            _ = v10
        case "a":
            var v11 int8 = 1
            _ = v11
        default:
            v12 := GenS2{F0: 1, F1: 128}
            _ = v12
        }
    }
    v13 := []int{692, 209, -2147483648, 958}
    _ = v13
    _ = v6
    _ = v7
    return 150
}
func genFn1(p0 string, p1 int8, p2 int64) int64 {
    var v14 float64 = float64(float32(0.1))
    var v15 int64 = (p2 % ((p2) | 1))
    if (!((2147483647 == -1) && false)) {
        v16 := []int{}
        if (0 > 398) {
            fmt.Println(p1, v15)
            p0 = (p0 + (("foo bar" + "a") + ""))
            var v17 int = 2147483647
            _ = v17
        } else {
            if len(v16) > 0 {
                v16[0] = -2147483648
            }
            fmt.Println(p1, v15, p2)
        }
        _ = v16
    } else {
        for fi18 := 0; fi18 < 13; fi18++ {
            p1 += (p1 % ((int8(v15)) | 1))
            p1 = (p1 - (p1 >> ((469) & 63)))
            _ = fi18
        }
        var v19 int = 32767
        _ = v19
    }
    var v20 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v20
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[string]uint32)
    var v3 int8 = -73
    var v4 uint32 = 2147483648
    var v5 uint16 = (v1 ^ (v1 & v1))
    var v21 uint32 = v4
    delete(v2, "hello")
    v22 := GenS3{F0: 0, F1: -367}
    var v23 uint32 = v21
    v24 := GenS4{F0: 2}
    var v25 int64 = 0
    var v26 string = ("" + "a")
    for fi27 := 0; fi27 < 3; fi27++ {
        if false {
            var v28 int = (fi27 << ((42) & 63))
            _ = v28
        }
        v26 = v26
        _ = fi27
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        v21 += v21
        v30 := []int{(fi29 >> ((1) & 63)), (fi29 + (fi29 << ((51) & 63)))}
        _ = fi29
        _ = v30
    }
    var v31 uint8 = 131
    v4 += (v21 & (v23 + (v21 >> ((uint(v4)) & 63))))
    v32 := GenS3{F0: (v25 ^ genFn1("\n\t", (v3 % ((v3) | 1)), -9223372036854775808)), F1: 1000}
    for fi33 := 0; fi33 < 7; fi33++ {
        v34 := GenS5{F0: -791}
        v23 -= v23
        fmt.Println(v3, fi33, fi33)
        _ = fi33
        _ = v34
    }
    v21 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
