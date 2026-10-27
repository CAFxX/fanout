package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int8 }
func genFn0(p0 uint16, p1 uint64, p2 int64) int32 {
    fmt.Println(p0, p0, p1)
    var v7 uint32 = 785
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: v7, F1: 2, F2: 116}
        var v10 bool = false
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return 1
}
func genFn1(p0 int) int8 {
    var v12 uint8 = 1
    var v13 uint8 = (v12 / ((v12) | 1))
    v14 := GenS1{F0: 540, F1: v12, F2: 63}
    if true {
        switch p0 {
        case 2147483647:
            var v15 uint32 = 2
            _ = v15
        }
        var v16 int32 = 2147483647
        _ = v16
    } else {
        v17 := []int{(p0 + (p0 & p0)), (p0 + p0)}
        var v18 bool = false
        _ = v17
        _ = v18
    }
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return -1
}

func main() {
    v1 := []uint32{65535, 4294967295, 2, 42}
    v2 := make(map[uint32]bool)
    var v3 string = ""
    v4 := make(map[uint32]uint32)
    var v5 int16 = 1000
    v6 := []int{1}
    var v19 float64 = 0.1
    fmt.Println(v3, v5, v3)
    var v20 uint64 = 203
    v21 := GenS1{F0: 4294967295, F1: 255, F2: -68}
    v22 := GenS1{F0: 0, F1: 42, F2: -128}
    var v23 int32 = 821
    _ = len(v1)
    var v24 int = -9223372036854775808
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
