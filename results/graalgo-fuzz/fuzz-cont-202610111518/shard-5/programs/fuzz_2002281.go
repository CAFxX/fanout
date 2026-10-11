package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint16 }

func main() {
    v1 := []uint32{}
    var v2 float32 = (float32(1.5) + float32(-2.25))
    var v3 uint16 = 255
    var v4 int8 = -104
    var v5 float64 = float64(3.14159)
    var v6 int32 = 2147483647
    fmt.Println(v6, v2)
    _ = len(v1)
    if false {
        switch ("\n\t" + "hello") {
        case "foo bar":
            var v7 int64 = -93
            _ = v7
        case "\n\t":
            v4 = 0
        case "":
            switch (("" + "\n\t") + "\n\t") {
            case "a":
                v6 = v6
            case "hello":
                var v8 float64 = (1e10 * 1e10)
                _ = v8
            }
        }
        var v9 bool = (((v4 * (v4 ^ v4)) < (^(v4))) || false)
        _ = v9
    } else {
        var v10 uint8 = 42
        fmt.Println(v10, v2)
        _ = v10
    }
    if true {
        var v11 float64 = float64((0.0 + v2))
        v12 := GenS1{F0: 1, F1: 1}
        v6 = (-(v6))
        _ = v11
        _ = v12
    }
    fmt.Println(v4, v3, v3)
    var v13 float32 = float32((1e10 / v5))
    if false {
        v4 -= v4
        v2 += v2
    } else {
        v14 := GenS2{F0: 1}
        _ = len(v1)
        v13 += (1e10 * v13)
        _ = v14
    }
    v1 = append(v1, 42)
    if (0 != 246) {
        if (1 == -259) {
            _ = len(v1)
            var v15 uint = 4294967295
            _ = v15
        }
    } else {
        var v16 uint32 = 0
        v17 := GenS3{F0: -9223372036854775808, F1: v3}
        _ = v16
        _ = v17
    }
    var v18 uint8 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
