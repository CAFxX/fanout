package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int32 }

func main() {
    var v1 float64 = ((float64(float32(0.1)) + -2.25) + 0.0)
    var v2 uint32 = 4294967295
    var v3 uint8 = 127
    for fi4 := 0; fi4 < 20; fi4++ {
        if true {
            switch "\n\t" {
            case "":
                v5 := GenS1{F0: 1, F1: 127, F2: (^(v3))}
                _ = v5
            case "\n\t":
                for fi6 := 0; fi6 < 14; fi6++ {
                    v7 := []int{(fi6 >> ((448) & 63)), fi6, (-(fi4))}
                    _ = fi6
                    _ = v7
                }
            }
            fmt.Println(v3, v3)
        }
        if (fi4 > 2147483647) {
            v3 += v3
            v1 += v1
            switch "hello" {
            case "hello":
                v1 = 1e10
            default:
                for fi8 := 0; fi8 < 1; fi8++ {
                    var v9 uint16 = 2
                    _ = fi8
                    _ = v9
                }
            }
        }
        v3 = 0
        _ = fi4
    }
    var v10 int32 = 0
    v11 := GenS2{F0: 155, F1: (^(v10))}
    v12 := GenS2{F0: 2, F1: v10}
    if true {
        fmt.Println(v2, v10)
        var v13 int32 = (v10 >> ((18446744073709551615) & 63))
        v14 := GenS1{F0: 2147483647, F1: -128, F2: 255}
        _ = v13
        _ = v14
    }
    fmt.Println(v10, v2, v10)
    v15 := []int{9223372036854775807, 1, 9223372036854775807}
    v10 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
}
