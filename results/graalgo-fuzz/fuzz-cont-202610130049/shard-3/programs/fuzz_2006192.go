package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int8; F1 uint }

func main() {
    var v1 uint = 255
    var v2 uint16 = 2
    var v3 int64 = -1
    var v4 float32 = 0.0
    var v5 int16 = 1
    v6 := []uint32{2147483648}
    v1 = v1
    var v7 bool = true
    switch "hello" {
    case "a":
        if (v3 == 0) {
            for fi8 := 0; fi8 < 6; fi8++ {
                v9 := GenS1{F0: (fi8 & (fi8 % ((-467) | 1))), F1: ((4294967295 > (^(fi8))) || (!((v2 | v2) == (v2 << ((v1) & 63)))))}
                _ = fi8
                _ = v9
            }
            v1 = v1
            var v10 uint = (v1 / ((1) | 1))
            _ = v10
        }
    case "\n\t":
        fmt.Println(v1, v1)
    default:
        fmt.Println(v5)
    }
    if v7 {
        if (v1 < (v1 - v1)) {
            switch 42 {
            case 127:
                var v11 int = 0
                _ = v11
            }
            if v7 {
                var v12 int8 = -128
                _ = v12
            } else {
                var v13 float64 = 1e10
                v14 := GenS1{F0: 1, F1: v7}
                fmt.Println(v5, v4, v5)
                _ = v13
                _ = v14
            }
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v2 -= (v2 | (v2 / ((1) | 1)))
    } else {
        fmt.Println(v7, v7)
    }
    var v15 uint64 = 18446744073709551615
    if (127 <= 30) {
        var v16 int8 = -128
        var v17 float32 = 0.1
        v15 = 707
        _ = v16
        _ = v17
    }
    var v18 float64 = 3.14159
    switch 2147483647 {
    case 2147483647:
        var v19 bool = (3.14159 > (v18 / (float64((v4 + 3.14159)) + ((1e10 + v18) + 1e10))))
        _ = v19
    case 32767:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        if v7 {
            if (!v7) {
                v20 := GenS2{F0: 127, F1: 0}
                fmt.Println(v3, v7)
                _ = v20
            }
            v6 = append(v6, 225)
            switch "foo bar" {
            case "\n\t":
                v21 := GenS2{F0: -128, F1: (v1 << ((v1) & 63))}
                _ = v21
            case "foo bar":
                v4 = v4
            case "":
                for fi22 := 0; fi22 < 17; fi22++ {
                    var v23 uint64 = (^(v15))
                    v4 -= (v4 - 0.0)
                    v23 -= (^(v23))
                    _ = fi22
                    _ = v23
                }
            }
        }
    }
    var v24 uint = v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
