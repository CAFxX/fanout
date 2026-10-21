package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint }
type GenS2 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int32; F1 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 255
    var v3 bool = (0 >= -308)
    var v4 float64 = (0.1 - 0.0)
    v5 := make(map[string]int)
    v6 := make(map[uint32]string)
    fmt.Println(v2, v4, v2)
    v3 = (379 >= 960)
    var v7 float32 = 1.5
    fmt.Println(v2)
    var v8 float32 = 0.1
    v9 := []int{9223372036854775807, -619}
    v2 -= 190
    v10 := GenS1{F0: v2, F1: "a", F2: 2}
    fmt.Println(v2, v3, v8)
    var v11 int64 = -1
    v12 := GenS2{F0: -2147483648, F1: 42, F2: 2}
    if (("foo bar" != "hello") || (v3 && (92 == 255))) {
        v13 := GenS1{F0: (^(v2)), F1: "hello", F2: uint(v2)}
        fmt.Println(v11, v2)
        _ = v13
    }
    var v14 uint32 = 2147483648
    if (!(!v3)) {
        for fi15 := 0; fi15 < 13; fi15++ {
            switch "foo bar" {
            case "a":
                if len(v9) > 0 {
                    v9[0] = 663
                }
            default:
                fmt.Println(v3)
            }
            if len(v9) > 0 {
                v9[0] = (fi15 & fi15)
            }
            _ = fi15
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            v17 := GenS2{F0: 0, F1: 1, F2: v14}
            _ = fi16
            _ = v17
        }
        v14 = 0
    } else {
        var v18 int32 = 2147483647
        var v19 float32 = v7
        switch "a" {
        case "":
            v2 += (^(v2))
        case "a":
            for fi20 := 0; fi20 < 3; fi20++ {
                v21 := GenS3{F0: v18, F1: (v18 + (v18 >> ((2) & 63)))}
                _ = len(v9)
                _ = fi20
                _ = v21
            }
        case "\n\t":
            v22 := v1[("" + "a")]
            _ = v22
        }
        _ = v18
        _ = v19
    }
    v8 += -2.25
    v7 = (0.1 / v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
