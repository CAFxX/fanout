package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint = 169
    var v2 int8 = -128
    var v3 int = 2147483647
    var v4 int = v3
    v4 -= v3
    var v5 string = "a"
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int16 = 310
        _ = fi6
        _ = v7
    }
    if (true || (65535 < 526)) {
        var v8 int8 = v2
        v2 = -107
        _ = v8
    } else {
        var v9 int = (v3 % ((v3) | 1))
        v10 := GenS1{F0: 42, F1: 2}
        if (uint32(v1) <= 42) {
            switch v5 {
            case "\n\t":
                var v11 int = (v9 ^ v4)
                _ = v11
            case "foo bar":
                var v12 uint = v1
                _ = v12
            default:
                var v13 int16 = -32768
                _ = v13
            }
            switch "" {
            case "foo bar":
                fmt.Println(v5, v5)
            case "\n\t":
                var v14 int = v3
                _ = v14
            case "a":
                var v15 uint = (v1 | 433)
                _ = v15
            default:
                v2 -= v2
            }
        }
        _ = v10
        _ = v9
    }
    var v16 uint64 = 1
    fmt.Println(v5)
    if (18446744073709551615 > (v16 | 460)) {
        fmt.Println(v16, v16, v1)
        v17 := GenS2{F0: 42}
        switch v5 {
        case "hello":
            v18 := GenS2{F0: uint32(v1)}
            _ = v18
        }
        _ = v17
    } else {
        v3 = (-(v4))
        v19 := GenS1{F0: 2, F1: -15}
        if true {
            var v20 float32 = (1.5 + (0.1 / 3.14159))
            _ = v20
        } else {
            if false {
                fmt.Println(v3, v1, v5)
            }
            var v21 float64 = 3.14159
            _ = v21
        }
        _ = v19
    }
    fmt.Println(v4, v4, v2)
    for fi22 := 0; fi22 < 15; fi22++ {
        if (true || true) {
            v2 += v2
            v4 += (v3 ^ 32768)
        }
        _ = fi22
    }
    v1 = (v1 >> (((v1 - 42)) & 63))
    if false {
        v2 += (v2 ^ v2)
    }
    v23 := GenS1{F0: 2147483647, F1: v2}
    if ((!true) || (v4 == (v4 | v3))) {
        var v24 string = v5
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
