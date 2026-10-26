package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int8 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = ((float64(0.0) * (0.1 + 1e10)) >= 1.5)
    v3 := make(map[string]bool)
    var v4 uint64 = 2
    switch 32767 {
    case 856:
        v5 := GenS1{F0: 127, F1: -9223372036854775808, F2: 0}
        _ = v5
    default:
        v6 := GenS1{F0: 209, F1: 9223372036854775807, F2: 128}
        _ = v6
    }
    v4 = 10
    var v7 int32 = -2147483648
    v2 = (1000 == 1000)
    var v8 uint32 = 2147483648
    if ((^(v4)) != (v4 ^ v4)) {
        if v2 {
            v8 = v8
            var v9 float32 = (0.1 + (0.0 + -2.25))
            _ = v9
        } else {
            v4 -= v4
            v10 := []int{-1, -34, 42}
            _ = v10
        }
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 float64 = 0.1
            _ = fi11
            _ = v12
        }
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v7)
    case "foo bar":
        if v2 {
            v1[32768] = 1
        }
    case "a":
        var v13 uint = 42
        _ = v13
    default:
        v14 := v1[-9223372036854775808]
        _ = v14
    }
    var v15 uint64 = v4
    if v2 {
        for fi16 := 0; fi16 < 13; fi16++ {
            v4 += (v4 + v4)
            v2 = v2
            if false {
                var v17 float64 = ((1.5 * 0.1) - 1e10)
                switch "" {
                case "hello":
                    var v18 int32 = (-(v7))
                    _ = v18
                case "":
                    var v19 uint32 = v8
                    _ = v19
                case "\n\t":
                    v20 := GenS1{F0: 255, F1: -9223372036854775808, F2: 128}
                    _ = v20
                }
                _ = v17
            }
            _ = fi16
        }
        var v21 float32 = 3.14159
        var v22 int8 = 1
        _ = v21
        _ = v22
    } else {
        v4 = (v15 << ((uint(v4)) & 63))
        switch "hello" {
        case "":
            v23 := []int{}
            _ = v23
        case "a":
            v24 := GenS2{F0: -32768, F1: v8}
            _ = v24
        }
        v25 := GenS3{F0: 127, F1: 1}
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
