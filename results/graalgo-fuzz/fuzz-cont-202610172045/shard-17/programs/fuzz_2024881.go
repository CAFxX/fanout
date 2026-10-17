package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    var v3 int32 = 42
    v4 := []string{("\n\t" + ("hello" + ("foo bar" + ("" + "a")))), "a", "\n\t"}
    v5 := make(map[string]bool)
    v6 := []uint32{65535, 2147483647, 4294967295, 0}
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    var v7 bool = false
    v8 := GenS1{F0: 2147483648, F1: false, F2: 42}
    if v7 {
        var v9 int8 = 127
        v7 = (((-1 != v3) && (!(v7 || v7))) || (v7 || (42 > 128)))
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 bool = v7
        _ = fi10
        _ = v11
    }
    v4 = append(v4, "a")
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            v7 = ((false && true) || v7)
            fmt.Println(v7, v7)
            _ = fi13
        }
        _ = fi12
    }
    switch "a" {
    case "a":
        v7 = v7
    default:
        v3 = (v3 & (v3 | (v3 << ((255) & 63))))
    }
    v3 = v3
    var v14 float32 = float32(float64((-2.25 - 3.14159)))
    fmt.Println(v7)
    _ = len(v6)
    for fi15 := 0; fi15 < 20; fi15++ {
        fmt.Println(fi15, fi15, v7)
        switch -9223372036854775808 {
        case 127:
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 uint16 = 255
                var v18 bool = (!v7)
                v3 = v3
                _ = fi16
                _ = v17
                _ = v18
            }
        case 0:
            if true {
                var v19 int64 = -962
                fmt.Println(v19, v19, fi15)
                _ = v19
            } else {
                v7 = (v3 != (v3 ^ v3))
                var v20 int32 = (v3 + (v3 * v3))
                _ = v20
            }
        case -9223372036854775808:
            v14 += v14
        default:
            fmt.Println(v14, v14)
        }
        switch "" {
        case "hello":
            for fi21 := 0; fi21 < 3; fi21++ {
                delete(v1, 2147483647)
                var v22 float64 = (0.0 / 1.0)
                _ = fi21
                _ = v22
            }
        default:
            v14 = v14
        }
        _ = fi15
    }
    v14 = v14
    if (0 != 39) {
        v7 = ((0 > 2) && (v7 && v7))
    }
    var v23 int32 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
