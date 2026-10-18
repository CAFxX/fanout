package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint }
func genFn0(p0 int8, p1 int32) uint16 {
    p0 -= (-(p0))
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := []int{fi7}
        fmt.Println(fi7, p0)
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 int64, p1 int8) uint16 {
    p0 += p0
    p0 = p0
    if ((p1 + p1) >= (p1 & p1)) {
        p1 -= p1
        if false {
            v9 := GenS1{F0: 42}
            if (2 < 840) {
                p0 += p0
                v10 := GenS1{F0: 1}
                v11 := []int{128, -5, 42, 0}
                _ = v10
                _ = v11
            }
            v12 := []int{}
            _ = v12
            _ = v9
        }
        fmt.Println(p1, p0)
    }
    v13 := []int{}
    _ = p0
    _ = p1
    _ = v13
    return 387
}

func main() {
    var v1 float32 = -2.25
    var v2 int8 = 0
    var v3 int16 = 1000
    v4 := make(map[string]int)
    v5 := make(map[uint32]uint32)
    var v6 int16 = 42
    fmt.Println(v1, v3, v1)
    var v14 uint16 = 32767
    for fi15 := 0; fi15 < 16; fi15++ {
        v2 = (v2 - v2)
        for fi16 := 0; fi16 < 13; fi16++ {
            v14 -= v14
            _ = fi16
        }
        _ = fi15
    }
    fmt.Println(v6, v1)
    var v17 int8 = v2
    v18 := GenS2{F0: (v6 << ((68) & 63)), F1: 255}
    v19 := []int{953}
    if len(v19) > 0 {
        v19[0] = 1
    }
    v3 -= (v6 + 284)
    v4["hello"] = 128
    if (v17 == 127) {
        v14 -= 656
        v19 = append(v19, 1)
    } else {
        if len(v19) > 0 {
            v19[0] = -462
        }
    }
    _ = len(v19)
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    if false {
        v14 = (v14 ^ (v14 << ((0) & 63)))
        v20 := GenS3{F0: 2, F1: 2}
        var v21 int8 = (-(v17))
        _ = v20
        _ = v21
    }
    var v22 uint16 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
