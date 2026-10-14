package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int }
func genFn0(p0 uint16) uint {
    p0 -= p0
    fmt.Println(p0)
    if true {
        var v7 int16 = 655
        _ = v7
    }
    v8 := GenS1{F0: (p0 << ((522) & 63)), F1: 1, F2: 768}
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := GenS1{F0: (p0 | (p0 % (((p0 & (p0 << ((uint(p0)) & 63)))) | 1))), F1: 1, F2: (fi9 & (fi9 << ((18446744073709551615) & 63)))}
        if true {
            var v11 uint32 = uint32(p0)
            _ = v11
        }
        p0 -= (p0 + 672)
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v8
    return 0
}
func genFn1(p0 uint64, p1 int) uint64 {
    fmt.Println(p1, p1)
    if (((-2.25 + -2.25) + 1e10) > (-2.25 / 1.0)) {
        if (true && true) {
            for fi12 := 0; fi12 < 2; fi12++ {
                fmt.Println(fi12, fi12)
                _ = fi12
            }
        }
    }
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 float32 = float32(1e10)
    v2 := make(map[uint32]int)
    var v3 uint8 = 127
    var v4 float64 = 0.0
    v5 := make(map[uint32]int)
    v6 := []uint32{278, 794}
    v1 -= v1
    _ = len(v6)
    if true {
        fmt.Println(v1)
    } else {
        for fi13 := 0; fi13 < 18; fi13++ {
            _ = len(v6)
            var v14 uint32 = 2147483648
            _ = fi13
            _ = v14
        }
        v6 = append(v6, 42)
        fmt.Println(v3)
    }
    var v15 int16 = -32768
    v16 := GenS1{F0: 255, F1: -63, F2: 42}
    var v17 int8 = -37
    var v18 uint32 = 2
    var v19 int = 32767
    if false {
        v4 = (v4 * 0.1)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v20 int32 = 695
        _ = v20
    }
    var v21 int = (v19 | 32767)
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 int16 = v15
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
