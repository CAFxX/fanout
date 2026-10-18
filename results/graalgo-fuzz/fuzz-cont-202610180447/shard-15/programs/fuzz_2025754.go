package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int64; F1 int16 }

func main() {
    var v1 uint8 = 127
    var v2 int8 = 127
    v3 := []uint32{317, 42, 2147483647}
    var v4 int64 = -1
    v5 := make(map[uint32]bool)
    var v6 bool = (42 <= 18446744073709551615)
    switch "foo bar" {
    case "a":
        var v7 uint = 0
        _ = v7
    case "":
        var v8 int = 0
        _ = v8
    }
    var v9 int64 = v4
    fmt.Println(v9)
    var v10 uint8 = v1
    if v6 {
        v11 := GenS1{F0: v10, F1: 205}
        var v12 bool = v6
        _ = v11
        _ = v12
    }
    v2 = (v2 + (v2 / ((v2) | 1)))
    v13 := GenS1{F0: v10, F1: 65535}
    fmt.Println(v2)
    var v14 float64 = 0.1
    fmt.Println(v4, v10, v9)
    var v15 uint16 = 1
    if v6 {
        var v16 float32 = float32(((((1.5 - (v14 + float64((-2.25 + 1.5)))) * (-2.25 + 1.5)) * v14) / v14))
        v17 := GenS2{F0: (v15 * (v15 ^ v15)), F1: 258}
        v4 = v4
        _ = v16
        _ = v17
    }
    v18 := GenS3{F0: (-(v4)), F1: 1}
    switch 0 {
    case 0:
        var v19 bool = (166 != 42)
        _ = v19
    case -1:
        if ("foo bar" >= (("\n\t" + "\n\t") + "a")) {
            var v20 uint8 = v1
            _ = v20
        } else {
            var v21 uint16 = (v15 % ((v15) | 1))
            v21 += v15
            _ = v21
        }
    case 128:
        var v22 int64 = v9
        _ = v22
    default:
        v23 := GenS2{F0: (v15 ^ v15), F1: 9223372036854775807}
        _ = v23
    }
    if (42 <= 42) {
        v1 -= v1
    } else {
        for fi24 := 0; fi24 < 1; fi24++ {
            v25 := v5[274]
            _ = fi24
            _ = v25
        }
        var v26 float32 = -2.25
        _ = v26
    }
    if true {
        switch "\n\t" {
        case "hello":
            var v27 uint32 = 118
            _ = v27
        }
        switch "" {
        case "\n\t":
            v4 = 380
        case "foo bar":
            for fi28 := 0; fi28 < 15; fi28++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi28
            }
        case "a":
            v29 := GenS2{F0: v15, F1: 1}
            _ = v29
        }
        for fi30 := 0; fi30 < 16; fi30++ {
            switch "a" {
            case "\n\t":
                _ = len(v3)
            }
            v31 := GenS2{F0: uint16(v1), F1: (fi30 / ((-806) | 1))}
            _ = fi30
            _ = v31
        }
    } else {
        v10 = (^(v10))
    }
    v32 := GenS3{F0: -9223372036854775808, F1: 1000}
    var v33 uint8 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
