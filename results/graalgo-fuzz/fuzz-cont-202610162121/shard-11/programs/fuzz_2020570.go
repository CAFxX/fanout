package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint }
type GenS2 struct { F0 string }
func genFn0(p0 bool, p1 bool, p2 string) int64 {
    switch -709 {
    case 42:
        v6 := []int{9223372036854775807, 2147483647, 128}
        _ = v6
    case 0:
        v7 := GenS1{F0: 32767, F1: 529, F2: 18446744073709551615}
        _ = v7
    default:
        p1 = (186 != 1)
    }
    var v8 float32 = 0.0
    p0 = (false || false)
    var v9 float64 = 0.1
    p0 = true
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint32, p1 uint, p2 int16) int {
    var v10 int16 = (p2 - (p2 / (((p2 >> ((704) & 63))) | 1)))
    p0 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 32768
}

func main() {
    v1 := []int{}
    v2 := []int{-2147483648, 32768, 32767, 9223372036854775807}
    var v3 float64 = 1e10
    v4 := make(map[uint32]int)
    var v5 float64 = (float64(0.0) * v3)
    var v11 bool = true
    fmt.Println(v11, v11)
    var v12 int16 = 1
    var v13 uint8 = 255
    fmt.Println(v13, v11, v12)
    v13 = (^(v13))
    v14 := GenS2{F0: ""}
    var v15 int32 = -3
    v13 = (v13 + (v13 * (v13 << ((1) & 63))))
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 uint64 = 42
        if v11 {
            if true {
                _ = len(v4)
                v17 = v17
                switch genFn1(391, 65535, v12) {
                case 128:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 42:
                    _ = len(v4)
                case 127:
                    var v18 int32 = (v15 >> ((361) & 63))
                    _ = v18
                }
            }
        }
        _ = fi16
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
