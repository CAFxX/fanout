package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int16; F2 string }
func genFn0(p0 int, p1 bool) uint16 {
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 int64 = int64(p0)
        _ = fi6
        _ = v7
    }
    var v8 int32 = 652
    fmt.Println(p0)
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: 32767}
        var v11 uint32 = 2
        _ = fi9
        _ = v10
        _ = v11
    }
    p1 = true
    _ = p0
    _ = p1
    _ = v8
    return 255
}
func genFn1(p0 string) int16 {
    p0 = p0
    if (42 > 220) {
        p0 = ((p0 + p0) + p0)
    } else {
        v12 := GenS1{F0: 623}
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 uint = 883
            fmt.Println(p0)
            fmt.Println(fi13, fi13, p0)
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 float64 = float64(float32(0.0))
            v17 := GenS1{F0: 183}
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = v12
    }
    if true {
        var v18 uint32 = 1
        for fi19 := 0; fi19 < 19; fi19++ {
            var v20 bool = false
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    if false {
        v21 := GenS1{F0: 1}
        var v22 uint16 = 239
        var v23 int64 = -727
        _ = v21
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: 664}
    _ = p0
    _ = v24
    return 462
}

func main() {
    var v1 float32 = 1e10
    var v2 uint32 = 0
    var v3 float64 = -2.25
    var v4 uint64 = 18446744073709551615
    v5 := []int{32767, 4294967295, 1, 460}
    var v25 float64 = v3
    v1 = 3.14159
    v26 := GenS2{F0: uint(v2), F1: -856, F2: ("" + "a")}
    v25 += v25
    var v27 uint8 = 220
    v28 := GenS1{F0: 756}
    for fi29 := 0; fi29 < 11; fi29++ {
        v4 -= v4
        _ = fi29
    }
    v1 = 0.0
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v2)
    v27 = (v27 + (v27 ^ (v27 ^ v27)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
