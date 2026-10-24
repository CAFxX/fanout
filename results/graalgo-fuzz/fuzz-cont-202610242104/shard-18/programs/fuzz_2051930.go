package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint; F1 uint; F2 int }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int, p1 uint16, p2 int32) int {
    fmt.Println(p1)
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    return (p0 << ((4294967295) & 63))
}
func genFn1(p0 int8) uint {
    if (true && false) {
        v7 := GenS1{F0: -9223372036854775808, F1: 18446744073709551615, F2: -2147483648}
        fmt.Println(p0)
        _ = v7
    }
    if (!false) {
        p0 -= (p0 >> ((65535) & 63))
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 uint64 = 2
            v10 := []int{(fi8 - (fi8 % (((fi8 * -1)) | 1))), (fi8 << ((924) & 63)), genFn0((fi8 * (-(fi8))), 65535, -14), 652}
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        switch "hello" {
        case "\n\t":
            var v11 float32 = 1.5
            _ = v11
        default:
            if false {
                p0 -= p0
                v12 := []int{genFn0(9223372036854775807, 255, 0), 128, 32768, -1}
                _ = v12
            }
        }
    }
    var v13 bool = true
    _ = p0
    _ = v13
    return 65535
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 2
    v3 := []int{127, 32767, -1}
    v4 := make(map[string]int)
    var v5 uint16 = 538
    var v6 uint = 973
    if (!false) {
        if false {
            fmt.Println(v5, v2)
        }
        _ = len(v4)
    }
    fmt.Println(v2)
    fmt.Println(v2, v6)
    v5 = (v5 & v5)
    for fi14 := 0; fi14 < 15; fi14++ {
        if (!(!false)) {
            for fi15 := 0; fi15 < 8; fi15++ {
                switch "foo bar" {
                case "":
                    v1["hello"] = (-(fi15))
                case "foo bar":
                    v5 = (^(v5))
                default:
                    v16 := GenS2{F0: v6, F1: 636, F2: fi15}
                    _ = v16
                }
                _ = fi15
            }
            v2 = v2
            v3 = append(v3, fi14)
        }
        _ = fi14
    }
    if true {
        var v17 int16 = 0
        var v18 uint16 = (v5 & v5)
        fmt.Println(v6)
        _ = v17
        _ = v18
    }
    v6 = (v6 ^ 18446744073709551615)
    v5 += 65535
    var v19 int32 = -2147483648
    var v20 uint8 = 1
    var v21 float64 = -2.25
    v22 := GenS3{F0: 2, F1: 474}
    var v23 int = 1
    fmt.Println(v5, v20, v23)
    v3 = append(v3, (-(v23)))
    v24 := GenS4{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
