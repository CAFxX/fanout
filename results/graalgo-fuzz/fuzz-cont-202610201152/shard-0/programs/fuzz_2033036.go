package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint; F2 int16 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint8; F1 int }

func main() {
    var v1 int16 = 1
    var v2 uint = 759
    var v3 int = 9223372036854775807
    var v4 bool = ((770 != 283) || true)
    var v5 uint = (v2 + (v2 - (v2 >> ((v2) & 63))))
    var v6 uint = (v2 * 4294967295)
    v7 := GenS1{F0: 2147483647, F1: 2147483648, F2: 240}
    var v8 int64 = -1
    if v4 {
        var v9 float64 = 3.14159
        v10 := GenS2{F0: v4, F1: 42, F2: (^(v1))}
        var v11 int8 = 1
        _ = v10
        _ = v11
        _ = v9
    }
    switch v3 {
    case 42:
        v5 = 4294967295
    case 32768:
        switch "foo bar" {
        case "hello":
            switch (v3 << ((v6) & 63)) {
            case 0:
                fmt.Println(v2)
            case 2147483647:
                v3 = -1
            case 127:
                v12 := GenS3{F0: v3, F1: 874}
                _ = v12
            }
        default:
            v6 += 2
        }
    default:
        var v13 bool = (v4 || v4)
        _ = v13
    }
    if (!true) {
        v14 := []int{-505, (v3 | 4294967295)}
        var v15 uint = v6
        fmt.Println(v15, v5, v5)
        _ = v14
        _ = v15
    } else {
        v16 := GenS4{F0: 2, F1: v3}
        v8 -= v8
        if (-1 == -132) {
            v17 := []int{}
            _ = v17
        } else {
            v3 = v3
            if v4 {
                v18 := []int{(v3 ^ (v3 + int(v5))), v3}
                for fi19 := 0; fi19 < 12; fi19++ {
                    v20 := GenS3{F0: 42, F1: 505}
                    var v21 int32 = -2147483648
                    var v22 uint32 = 148
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                _ = v18
            }
            v23 := []int{(v3 ^ v3), 9223372036854775807}
            _ = v23
        }
        _ = v16
    }
    v24 := []int{v3, (v3 / ((v3) | 1))}
    var v25 int64 = v8
    fmt.Println(v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
