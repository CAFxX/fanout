package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 int64, p1 int) string {
    p0 = (p0 >> ((504) & 63))
    var v5 bool = (42 != 210)
    for fi6 := 0; fi6 < 17; fi6++ {
        v5 = (fi6 > 310)
        v7 := GenS1{F0: 128}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return ""
}
func genFn1(p0 uint16, p1 uint64, p2 int32) int64 {
    fmt.Println(p2, p0)
    p2 -= -864
    for fi8 := 0; fi8 < 20; fi8++ {
        p2 = 1
        _ = fi8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := []int{}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}

func main() {
    v1 := []int{32768, -37}
    v2 := make(map[int]bool)
    var v3 uint16 = 839
    var v4 float64 = (0.1 * float64(3.14159))
    var v11 uint32 = 1
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 uint = 42
        for fi14 := 0; fi14 < 19; fi14++ {
            v15 := GenS2{F0: (128 <= 105)}
            var v16 uint = v13
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v4, v3)
    switch 4294967295 {
    case -2147483648:
        var v17 uint16 = 255
        _ = v17
    }
    var v18 float32 = -2.25
    fmt.Println(v18, v18)
    var v19 int64 = 1
    v4 = (v4 - v4)
    v20 := GenS1{F0: uint8(v3)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
