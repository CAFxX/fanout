package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
func genFn0(p0 int, p1 uint16, p2 uint64) uint {
    for fi4 := 0; fi4 < 12; fi4++ {
        p2 -= 42
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 uint16 = p1
            fmt.Println(v6)
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    p1 = p1
    p1 += p1
    p1 -= (p1 / ((p1) | 1))
    _ = p0
    _ = p1
    _ = p2
    return 734
}

func main() {
    v1 := []uint32{1, 1}
    v2 := make(map[string]string)
    v3 := make(map[uint32]string)
    v7 := v3[554]
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int64 = -9223372036854775808
        fmt.Println(fi8, v9, v9)
        _ = fi8
        _ = v9
    }
    v7 = "a"
    for fi10 := 0; fi10 < 7; fi10++ {
        v1 = append(v1, 2)
        if (18446744073709551615 >= 1) {
            var v11 uint8 = 42
            var v12 int8 = 76
            _ = v11
            _ = v12
        }
        if (2 < 1) {
            var v13 uint64 = 382
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 float64 = float64(((1.5 * float32(-2.25)) - 0.0))
                _ = fi14
                _ = v15
            }
            _ = v13
        } else {
            if false {
                var v16 int64 = 9223372036854775807
                var v17 int64 = 1
                fmt.Println(v7, v16, v16)
                _ = v16
                _ = v17
            } else {
                var v18 int = (fi10 - fi10)
                var v19 bool = false
                var v20 float32 = float32(-2.25)
                _ = v18
                _ = v19
                _ = v20
            }
        }
        _ = fi10
    }
    var v21 bool = false
    switch ("hello" + v7) {
    case "\n\t":
        var v22 bool = ((1.5 * 3.14159) != 3.14159)
        _ = v22
    default:
        v7 = "a"
    }
    switch v7 {
    case "":
        v21 = v21
    case "foo bar":
        v23 := GenS1{F0: "foo bar", F1: v21}
        _ = v23
    }
    var v24 int16 = 32767
    switch "foo bar" {
    case "":
        if v21 {
            fmt.Println(v24, v21, v24)
            _ = len(v1)
        }
    default:
        v21 = v21
    }
    v7 = v7
    var v25 uint8 = 127
    var v26 int16 = int16(v25)
    var v27 int8 = -128
    var v28 uint = 255
    var v29 float32 = 0.0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v7
}
