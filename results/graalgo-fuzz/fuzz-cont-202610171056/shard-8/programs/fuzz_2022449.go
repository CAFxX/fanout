package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int; F2 int64 }

func main() {
    var v1 uint8 = 0
    v2 := []uint32{531}
    var v3 int8 = -128
    var v4 float64 = ((0.1 + (1.5 + 1.5)) / (3.14159 * -2.25))
    var v5 int = -447
    var v6 uint8 = (v1 & (v1 | 255))
    switch v5 {
    case 42:
        v5 -= (v5 & v5)
    case 4294967295:
        _ = len(v2)
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: (v6 >> ((1) & 63))}
        var v9 int16 = -36
        var v10 string = "a"
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    switch "hello" {
    case "foo bar":
        var v11 uint64 = 966
        _ = v11
    case "hello":
        fmt.Println(v6, v4, v6)
    default:
        fmt.Println(v1)
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            switch "\n\t" {
            case "a":
                v14 := GenS1{F0: (v6 & 2)}
                _ = v14
            case "\n\t":
                var v15 int8 = v3
                _ = v15
            }
            fmt.Println(fi12, v3)
            _ = fi13
        }
        _ = fi12
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 int64 = -9223372036854775808
        v18 := GenS2{F0: 32767, F1: (fi16 ^ -9223372036854775808), F2: v17}
        _ = fi16
        _ = v17
        _ = v18
    }
    if true {
        var v19 string = "hello"
        _ = v19
    } else {
        fmt.Println(v1, v1, v5)
        if len(v2) > 0 {
            v2[0] = 0
        }
        var v20 int8 = v3
        _ = v20
    }
    var v21 uint32 = 42
    var v22 uint8 = 42
    v23 := GenS2{F0: -853, F1: (v5 * v5), F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
