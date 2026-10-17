package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint8 }
func genFn0() uint64 {
    var v4 uint = 18446744073709551615
    var v5 int64 = 9223372036854775807
    var v6 uint8 = 196
    var v7 int = 196
    var v8 int16 = 699
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 649
}

func main() {
    var v1 int16 = 1000
    var v2 uint8 = 2
    var v3 float32 = 0.0
    v3 = (1e10 - v3)
    var v9 uint16 = 2
    switch "\n\t" {
    case "hello":
        v9 = (v9 >> ((491) & 63))
    default:
        var v10 int8 = 18
        _ = v10
    }
    v11 := []int{}
    var v12 int16 = v1
    for fi13 := 0; fi13 < 6; fi13++ {
        v1 = (v1 - 42)
        var v14 float64 = -2.25
        v1 = (^(v12))
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: (v2 % ((v2) | 1)), F1: 255}
    v11 = append(v11, 4294967295)
    v16 := GenS1{F0: (v2 / ((v2) | 1)), F1: 1}
    v3 += (v3 / (v3 * 1.5))
    fmt.Println(v12, v3)
    v3 += 1e10
    v12 -= -724
    v17 := GenS1{F0: (^(v2)), F1: 476}
    v2 -= v2
    for fi18 := 0; fi18 < 2; fi18++ {
        if (!(!false)) {
            for fi19 := 0; fi19 < 2; fi19++ {
                var v20 int = 286
                if true {
                    if ((911 != 279) || (v3 <= v3)) {
                        v12 = v12
                        _ = len(v11)
                    }
                    var v21 float32 = (3.14159 + (float32((1e10 * 0.1)) / v3))
                    _ = v21
                }
                _ = fi19
                _ = v20
            }
            for fi22 := 0; fi22 < 13; fi22++ {
                v23 := GenS2{F0: 4294967295, F1: 42}
                _ = len(v11)
                _ = fi22
                _ = v23
            }
        } else {
            var v24 int = 4294967295
            v25 := GenS3{F0: v2}
            _ = v24
            _ = v25
        }
        v2 += v2
        fmt.Println(fi18, v2, v9)
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
