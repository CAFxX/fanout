package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int32, p1 uint8) string {
    fmt.Println(p1)
    switch -1 {
    case -976:
        v5 := GenS1{F0: "", F1: (p1 | p1), F2: -1}
        _ = v5
    case -678:
        if ((0.1 - (0.1 - 1e10)) < (1.5 + 0.1)) {
            v6 := []int{32768, 127, 145, 32768}
            fmt.Println(p1, p0)
            _ = v6
        }
    case 2147483647:
        var v7 int16 = -32768
        _ = v7
    }
    _ = p0
    _ = p1
    return "hello"
}
func genFn1() uint8 {
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 uint32 = 2147483647
        var v10 int = fi8
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 float64 = float64((1.5 * 0.0))
    if (42 > 0) {
        fmt.Println(v11)
    } else {
        v12 := []int{-1}
        _ = v12
    }
    if (127 != 42) {
        switch 32768 {
        case 204:
            v13 := []int{128, 87}
            _ = v13
        default:
            if (838 != 0) {
                v14 := []int{}
                v11 += (1e10 - ((v11 + (0.0 / v11)) + 0.1))
                _ = v14
            }
        }
        var v15 int8 = 127
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 int64 = -9223372036854775808
            if false {
                v17 -= v17
                v17 += -593
            } else {
                var v18 uint64 = 1
                switch "foo bar" {
                case "hello":
                    v19 := GenS2{F0: 127}
                    _ = v19
                }
                _ = v18
            }
            fmt.Println(v17, fi16)
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        v20 := []int{32767}
        _ = v20
    }
    _ = v11
    return 0
}

func main() {
    v1 := []int{618, -232, -782, 241}
    var v2 uint = 65535
    var v3 uint64 = 311
    var v4 int64 = 0
    v2 = 18446744073709551615
    _ = len(v1)
    v3 += v3
    switch "hello" {
    case "foo bar":
        v4 = (v4 << ((18446744073709551615) & 63))
    case "":
        var v21 int8 = -72
        _ = v21
    case "hello":
        v1 = append(v1, 0)
    }
    if (!true) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        fmt.Println(v4, v3)
        switch ("foo bar" + "hello") {
        case "hello":
            var v23 int8 = -128
            _ = v23
        case "":
            v24 := GenS2{F0: 1}
            _ = v24
        case "\n\t":
            var v25 int32 = 2147483647
            _ = v25
        default:
            if true {
                _ = len(v1)
            }
        }
        _ = fi22
    }
    v1 = append(v1, 128)
    var v26 float32 = 3.14159
    var v27 uint32 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (-2.25 != v26) {
        var v28 uint64 = 573
        if (255 == 2) {
            v3 -= uint64(v27)
        } else {
            var v29 uint32 = v27
            _ = v29
        }
        _ = v28
    } else {
        if len(v1) > 0 {
            v1[0] = 32767
        }
    }
    fmt.Println(v27, v4, v26)
    v27 += 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
