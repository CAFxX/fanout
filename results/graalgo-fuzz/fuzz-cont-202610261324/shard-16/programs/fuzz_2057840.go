package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
func genFn0(p0 uint16, p1 int64) int32 {
    fmt.Println(p1)
    v4 := []int{}
    var v5 uint64 = 18446744073709551615
    var v6 int = -592
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 346
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 65535
    var v7 uint8 = 226
    for fi8 := 0; fi8 < 10; fi8++ {
        if true {
            if true {
                fmt.Println(v1, fi8, v7)
                var v9 int64 = int64(v7)
                _ = v9
            }
            if (-128 != int8(v3)) {
                var v10 int32 = 0
                fmt.Println(fi8, v7, v10)
                var v11 int64 = (v1 - 236)
                _ = v10
                _ = v11
            } else {
                if false {
                    var v12 int16 = 32767
                    switch "" {
                    case "\n\t":
                        v3 = (v3 ^ v3)
                    default:
                        var v13 int32 = -1
                        _ = v13
                    }
                    _ = v12
                }
                v3 = (^(v3))
                v14 := []int{fi8, fi8, fi8}
                _ = v14
            }
        } else {
            _ = len(v2)
            var v15 float64 = ((0.1 + 1.5) * 0.1)
            _ = v15
        }
        v7 = (v7 * (^(v7)))
        v16 := v2[42]
        _ = fi8
        _ = v16
    }
    v7 += 0
    var v17 int16 = -32768
    v18 := GenS1{F0: 293, F1: 2}
    v19 := GenS1{F0: 2147483647, F1: 226}
    v20 := []int{-1}
    v17 = v17
    fmt.Println(v17, v1, v17)
    v21 := GenS1{F0: -66, F1: 4294967295}
    fmt.Println(v7, v17, v1)
    v22 := GenS1{F0: 648, F1: 2}
    for fi23 := 0; fi23 < 8; fi23++ {
        v17 += (v17 ^ v17)
        v3 -= (v3 & v3)
        _ = fi23
    }
    v17 += (v17 / (((v17 / ((v17) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v7
}
