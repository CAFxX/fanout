package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int16, p1 int32) uint {
    p1 = (p1 % ((p1) | 1))
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := []int{int(p0), -2147483648}
        _ = fi7
        _ = v8
    }
    if ((p0 + (p0 & (p0 * p0))) >= (p0 - p0)) {
        p1 += p1
    }
    p0 -= (p0 + (p0 << ((255) & 63)))
    if (true && false) {
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 float64 = 1.5
            v11 := GenS1{F0: 100}
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 uint = 1
        v13 := GenS1{F0: 1}
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    return 73
}

func main() {
    v1 := []uint32{642}
    var v2 int16 = 1000
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    v5 := []int{42}
    v6 := make(map[uint32]string)
    var v14 float64 = (3.14159 - 1.5)
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS1{F0: -22}
        var v17 string = "a"
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint16 = 2
    v18 -= (v18 * v18)
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    switch 416 {
    case -1:
        var v19 int16 = -903
        _ = v19
    }
    v18 += v18
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = len(v6)
    var v20 int64 = 614
    var v21 uint8 = 2
    v22 := GenS2{F0: 0, F1: 18446744073709551615}
    var v23 int32 = -547
    var v24 int32 = int32(v2)
    if ((v2 << ((255) & 63)) > 768) {
        v25 := GenS1{F0: 100}
        _ = v25
    }
    fmt.Println(v14)
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 int = (fi26 >> ((genFn0(v2, v24)) & 63))
        fmt.Println(v14, v21)
        _ = fi26
        _ = v27
    }
    if true {
        v28 := GenS3{F0: v23}
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v18
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
