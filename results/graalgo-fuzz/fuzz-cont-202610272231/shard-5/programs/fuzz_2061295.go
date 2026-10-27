package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 0.0
    var v3 uint64 = 4294967295
    var v4 float64 = v2
    v5 := make(map[uint32]int)
    for fi6 := 0; fi6 < 11; fi6++ {
        if false {
            v3 -= 18446744073709551615
        }
        switch fi6 {
        case 0:
            v7 := GenS1{F0: 42, F1: (fi6 >> ((1) & 63))}
            _ = v7
        case 162:
            fmt.Println(v4)
        }
        _ = fi6
    }
    v8 := []int{128, -2147483648, 32768, -1}
    v4 = v4
    v2 += v4
    var v9 uint = 303
    v3 = (v3 + 2)
    v10 := GenS1{F0: -1, F1: -1}
    if len(v8) > 0 {
        v8[0] = -9223372036854775808
    }
    fmt.Println(v4, v2)
    var v11 int32 = -556
    v4 += float64((0.0 * (-2.25 - -2.25)))
    v4 += 0.1
    var v12 int8 = 1
    switch "a" {
    case "":
        if false {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            for fi13 := 0; fi13 < 13; fi13++ {
                v2 -= (float64(0.1) - (-2.25 * v2))
                _ = fi13
            }
        } else {
            for fi14 := 0; fi14 < 16; fi14++ {
                delete(v1, fi14)
                for fi15 := 0; fi15 < 13; fi15++ {
                    var v16 int = -2147483648
                    v3 = 22
                    fmt.Println(fi14)
                    _ = fi15
                    _ = v16
                }
                _ = fi14
            }
            v12 -= (v12 | v12)
        }
    case "\n\t":
        var v17 uint8 = 49
        _ = v17
    default:
        var v18 int8 = v12
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
