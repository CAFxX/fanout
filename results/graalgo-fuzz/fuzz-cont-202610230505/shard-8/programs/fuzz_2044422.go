package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int; F1 int8 }
func genFn0(p0 uint16, p1 int8) string {
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 int16 = -217
        p1 = 100
        v7 := []int{(^(fi5)), (fi5 >> ((255) & 63))}
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        switch fi8 {
        case -94:
            fmt.Println(fi8)
        }
        if ((p0 / ((p0) | 1)) < 65535) {
            v9 := []int{(fi8 >> ((4294967295) & 63)), (fi8 >> ((4294967295) & 63))}
            v10 := GenS1{F0: (p0 != p0), F1: uint64(p0), F2: 38}
            var v11 int = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi8
    }
    var v12 int16 = int16(p0)
    v13 := GenS1{F0: true, F1: 18446744073709551615, F2: 2}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return "foo bar"
}

func main() {
    var v1 uint16 = 2
    var v2 float64 = (float64((0.1 + (0.0 * 0.0))) * float64(1.5))
    var v3 int = -9223372036854775808
    var v4 int = v3
    v14 := GenS2{F0: (v4 ^ (-(v4))), F1: 1}
    v15 := []int{}
    fmt.Println(v4, v2, v1)
    fmt.Println(v4)
    var v16 uint = 2
    v2 += v2
    v1 = (v1 % ((uint16(v16)) | 1))
    v1 += (v1 / (((v1 - (v1 + (v1 + (^(v1)))))) | 1))
    fmt.Println(v3)
    var v17 int16 = 317
    v1 += v1
    if (!(!true)) {
        switch ("hello" + ("" + "a")) {
        case "":
            var v18 uint = v16
            _ = v18
        case "hello":
            v16 -= 326
        case "foo bar":
            switch (-(v3)) {
            case 951:
                fmt.Println(v4)
            default:
                var v19 uint8 = 128
                _ = v19
            }
        default:
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 uint16 = 32767
                _ = fi20
                _ = v21
            }
        }
        v17 -= 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
