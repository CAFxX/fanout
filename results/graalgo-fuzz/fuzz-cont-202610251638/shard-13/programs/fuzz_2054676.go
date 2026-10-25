package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 bool) uint64 {
    fmt.Println(p0)
    p0 = (1 != -94)
    v5 := GenS1{F0: 784}
    var v6 int8 = 127
    _ = p0
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint32) uint64 {
    var v7 uint32 = p0
    var v8 uint8 = 255
    switch "foo bar" {
    case "hello":
        v9 := GenS2{F0: p0}
        _ = v9
    default:
        var v10 int16 = 707
        _ = v10
    }
    var v11 uint16 = 2
    _ = p0
    _ = v11
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint32 = 1
    var v2 int32 = -2147483648
    v3 := make(map[int]int)
    var v4 uint = 18446744073709551615
    if (127 >= -1) {
        fmt.Println(v4)
        v12 := GenS1{F0: (v4 | (v4 + (v4 | (v4 - 255))))}
        var v13 float64 = ((3.14159 * 1.5) * (-2.25 / 3.14159))
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        v3[0] = fi14
        _ = fi14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        if false {
            _ = len(v3)
            for fi16 := 0; fi16 < 1; fi16++ {
                if (9223372036854775807 <= 42) {
                    fmt.Println(fi15)
                    v17 := GenS1{F0: v4}
                    var v18 uint8 = 237
                    _ = v17
                    _ = v18
                }
                switch "\n\t" {
                case "\n\t":
                    v3[fi16] = fi16
                default:
                    v2 -= 2147483647
                }
                v2 = (v2 | (v2 * v2))
                _ = fi16
            }
        } else {
            var v19 int8 = -50
            _ = v19
        }
        switch 32767 {
        case 826:
            fmt.Println(v4, v2)
        }
        v2 = v2
        _ = fi15
    }
    fmt.Println(v1)
    v1 -= 0
    v1 = v1
    v20 := GenS2{F0: 2147483647}
    v2 += v2
    v3[9223372036854775807] = 9223372036854775807
    v21 := []int{}
    switch "foo bar" {
    case "\n\t":
        var v22 int32 = -80
        _ = v22
    case "foo bar":
        _ = len(v21)
    default:
        var v23 int8 = 0
        _ = v23
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        fmt.Println(v2)
        v25 := GenS3{F0: (v2 + v2)}
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
