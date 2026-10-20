package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
func genFn0() uint16 {
    var v7 uint = 974
    v8 := []int{182, -703}
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(fi9)
        fmt.Println(v7, fi9)
        _ = fi9
    }
    _ = v7
    _ = v8
    return 875
}
func genFn1(p0 int32, p1 int) uint8 {
    for fi10 := 0; fi10 < 18; fi10++ {
        p0 = (p0 << ((870) & 63))
        var v11 bool = true
        _ = fi10
        _ = v11
    }
    p1 += p1
    var v12 float64 = ((0.0 * -2.25) + ((0.0 + 0.0) / -2.25))
    var v13 uint8 = 255
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return v13
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{32767}
    var v3 uint16 = 239
    v4 := make(map[string]int)
    v5 := []int{128, 2147483647}
    var v6 uint8 = 127
    var v14 int64 = -1
    if false {
        var v15 float32 = 0.1
        for fi16 := 0; fi16 < 20; fi16++ {
            if (42 >= 4294967295) {
                fmt.Println(fi16, v3)
            }
            if len(v5) > 0 {
                v5[0] = (fi16 / (((fi16 + fi16)) | 1))
            }
            _ = fi16
        }
        v15 -= v15
        _ = v15
    }
    var v17 uint8 = v6
    var v18 float32 = ((0.0 / ((3.14159 - (0.1 + 0.0)) - 1.5)) / 1e10)
    v19 := GenS1{F0: 2147483648, F1: 855}
    v20 := GenS1{F0: 4294967295, F1: 787}
    var v21 int64 = (v14 - 175)
    if len(v2) > 0 {
        v2[0] = -1
    }
    var v22 int32 = -1
    var v23 int64 = (v21 | v21)
    v3 = genFn0()
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
