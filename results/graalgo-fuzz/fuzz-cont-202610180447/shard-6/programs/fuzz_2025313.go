package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint8, p1 int32, p2 int64) uint8 {
    v7 := []int{}
    if true {
        p2 -= (p2 << ((749) & 63))
        for fi8 := 0; fi8 < 12; fi8++ {
            p0 = (p0 >> ((0) & 63))
            p0 -= p0
            _ = fi8
        }
    } else {
        p1 += (p1 / (((p1 - p1)) | 1))
        var v9 float64 = (0.0 - ((((1e10 - float64(0.1)) * -2.25) + (0.1 / float64((0.0 + (0.1 / 1e10))))) * 1.5))
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p0 - (p0 * p0))
}
func genFn1(p0 int32, p1 uint, p2 int) int16 {
    fmt.Println(p0)
    p1 += (^(p1))
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    v3 := []string{"foo bar", "foo bar", ("a" + ""), "foo bar"}
    v4 := make(map[uint32]bool)
    v5 := []int{-684, 177, 42}
    var v6 uint = 727
    for fi10 := 0; fi10 < 8; fi10++ {
        v5 = append(v5, (^(fi10)))
        v6 = 4294967295
        _ = fi10
    }
    var v11 uint = 65535
    fmt.Println(v11, v6)
    v12 := GenS1{F0: "foo bar", F1: 2147483647, F2: 2147483647}
    var v13 uint = (v6 + v6)
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := GenS2{F0: 127}
        fmt.Println(fi14, v13)
        _ = fi14
        _ = v15
    }
    var v16 float64 = float64((1e10 - ((3.14159 / 1.0) / -2.25)))
    v17 := GenS2{F0: 110}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
