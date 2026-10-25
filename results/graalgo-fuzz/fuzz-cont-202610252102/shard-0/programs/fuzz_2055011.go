package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
func genFn0(p0 string, p1 uint) uint32 {
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := []int{fi8, (fi7 & (fi8 ^ fi7)), 802, (^(fi7))}
            _ = fi8
            _ = v9
        }
        v10 := []int{fi7, (fi7 ^ (fi7 - (-(fi7)))), fi7, 32767}
        var v11 int = 0
        _ = fi7
        _ = v10
        _ = v11
    }
    v12 := []int{875, -1, 4294967295}
    v13 := GenS1{F0: 466, F1: 127}
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 string = ""
        var v16 float64 = 0.0
        fmt.Println(p0, p0)
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 159
}
func genFn1(p0 uint, p1 string) uint64 {
    p1 = p1
    p1 = ""
    var v17 string = (p1 + "foo bar")
    v18 := []int{32768, 136, 900, -790}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    var v3 int8 = 1
    v4 := []uint32{}
    var v5 bool = false
    v6 := make(map[string]int)
    _ = len(v4)
    var v19 int64 = 1
    v5 = v5
    var v20 uint64 = 2
    if true {
        var v21 uint = 2
        _ = v21
    }
    v22 := GenS1{F0: -208, F1: 1}
    _ = len(v1)
    var v23 int = 9223372036854775807
    v24 := GenS1{F0: 0, F1: 255}
    v19 += v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
