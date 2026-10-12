package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 uint; F2 string }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{-1, -68}
    v3 := make(map[string]bool)
    v4 := make(map[uint32]uint32)
    v5 := []uint32{}
    v6 := []int{}
    if true {
        _ = len(v2)
    } else {
        v7 := GenS1{F0: true, F1: 4294967295, F2: 255}
        _ = v7
    }
    var v8 int8 = 100
    var v9 uint = 761
    var v10 int = -2147483648
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch (v10 / ((v10) | 1)) {
    case -9223372036854775808:
        if true {
            switch "a" {
            case "hello":
                _ = len(v6)
            case "\n\t":
                var v11 int32 = 477
                _ = v11
            default:
                v8 -= v8
            }
            fmt.Println(v8)
            v12 := GenS2{F0: (v8 + (-(v8)))}
            _ = v12
        }
    case -724:
        switch "\n\t" {
        case "\n\t":
            var v13 uint32 = 65535
            _ = v13
        case "foo bar":
            switch "" {
            case "":
                var v14 float32 = 3.14159
                _ = v14
            case "hello":
                var v15 float64 = (float64((0.1 * 0.1)) - -2.25)
                _ = v15
            case "foo bar":
                v16 := GenS3{F0: 2, F1: v9, F2: "hello"}
                _ = v16
            default:
                v9 = 18446744073709551615
            }
        case "":
            var v17 uint8 = 231
            _ = v17
        }
    default:
        var v18 int16 = -832
        _ = v18
    }
    var v19 int64 = -1
    fmt.Println(v8, v19, v8)
    var v20 uint64 = 329
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
