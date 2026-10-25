package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int32 }
func genFn0(p0 bool, p1 uint64) string {
    for fi5 := 0; fi5 < 12; fi5++ {
        fmt.Println(p0, fi5)
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 float64 = float64((float32(1.5) / 1.0))
            v7 = v7
            fmt.Println(p0)
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    p0 = (!false)
    v8 := []int{-848, -9223372036854775808, 32768, 42}
    _ = p0
    _ = p1
    _ = v8
    return "hello"
}

func main() {
    v1 := []int{-1, -9223372036854775808, 2147483647}
    v2 := make(map[int]uint32)
    var v3 uint = 255
    v4 := []uint32{1, 4294967295}
    _ = len(v4)
    v3 = (v3 | 2)
    v3 = (v3 ^ (v3 >> (((v3 * (v3 & 490))) & 63)))
    switch 0 {
    case -9223372036854775808:
        v9 := GenS1{F0: v3}
        _ = v9
    }
    v10 := GenS1{F0: 42}
    delete(v2, 0)
    v3 -= (v3 + v3)
    v1 = append(v1, 738)
    var v11 uint32 = 465
    v12 := GenS2{F0: 42, F1: 676}
    var v13 int64 = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
