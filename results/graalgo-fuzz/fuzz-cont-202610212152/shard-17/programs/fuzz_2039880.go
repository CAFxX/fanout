package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int16; F1 int64; F2 uint }
func genFn0(p0 int) uint {
    if (2 != 255) {
        fmt.Println(p0)
    }
    var v6 float64 = 1e10
    p0 += p0
    var v7 int16 = 0
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := []int{}
        fmt.Println(fi8, v7)
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int32) uint8 {
    switch "" {
    case "foo bar":
        var v10 int = 32768
        _ = v10
    case "hello":
        var v11 int32 = p0
        _ = v11
    default:
        p0 = 42
    }
    v12 := GenS1{F0: 2}
    var v13 uint32 = 457
    var v14 int8 = 3
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            switch ("" + "") {
            case "":
                fmt.Println(p0, p0, v14)
            case "\n\t":
                v17 := GenS1{F0: uint16(v13)}
                _ = v17
            case "foo bar":
                var v18 uint16 = 2
                _ = v18
            }
            _ = fi16
        }
        v19 := []int{(^(fi15)), -9223372036854775808, fi15, 128}
        var v20 uint64 = 903
        _ = fi15
        _ = v19
        _ = v20
    }
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 218
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"\n\t", "foo bar", ("hello" + "a")}
    v3 := []string{((("" + "\n\t") + "\n\t") + ("foo bar" + (((("foo bar" + "\n\t") + "hello") + ("hello" + "")) + ""))), ("\n\t" + ("hello" + "\n\t")), "\n\t"}
    var v4 int8 = -128
    var v5 bool = ("hello" != "\n\t")
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    if (0 != 1000) {
        v4 += v4
        v4 = v4
    } else {
        fmt.Println(v5)
    }
    switch 127 {
    case 32768:
        fmt.Println(v4)
    }
    v4 = -128
    _ = len(v3)
    var v21 int = 2147483647
    v22 := GenS2{F0: 1000, F1: 2}
    v21 -= v21
    v4 -= 0
    var v23 float32 = (0.0 + (float32((0.1 + (float64((0.0 * 1.5)) * -2.25))) / 1.0))
    v3 = append(v3, "a")
    _ = len(v1)
    v5 = (1 != 0)
    v5 = (-128 > v4)
    for fi24 := 0; fi24 < 12; fi24++ {
        if (((float64(v23) - 3.14159) + 1e10) == 0.1) {
            v4 += (v4 - v4)
        }
        v25 := GenS2{F0: 1000, F1: 4294967295}
        _ = fi24
        _ = v25
    }
    v26 := GenS3{F0: -276, F1: 58, F2: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
