package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int }
type GenS2 struct { F0 int64; F1 uint16; F2 uint }
func genFn0(p0 int8, p1 int) uint16 {
    v5 := GenS1{F0: true, F1: false, F2: p1}
    fmt.Println(p0, p1)
    p0 = 42
    _ = p0
    _ = p1
    _ = v5
    return 2
}
func genFn1(p0 int8) uint16 {
    p0 += p0
    if false {
        p0 += (p0 << ((0) & 63))
    } else {
        var v6 int8 = p0
        _ = v6
    }
    var v7 uint = 890
    v8 := []int{-634, -131}
    _ = p0
    _ = v7
    _ = v8
    return 647
}

func main() {
    var v1 uint32 = 716
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 128
    v4 := make(map[int]int)
    v2[65535] = v1
    v9 := GenS1{F0: true, F1: (-9223372036854775808 != 323), F2: -1}
    if (false && false) {
        fmt.Println(v3, v1)
    } else {
        var v10 int = -533
        _ = v10
    }
    var v11 uint64 = 0
    var v12 uint64 = (v11 - (v11 | (v11 ^ (v11 | v11))))
    var v13 string = ("\n\t" + "hello")
    for fi14 := 0; fi14 < 4; fi14++ {
        fmt.Println(v11, v12)
        _ = fi14
    }
    v15 := v4[-645]
    v16 := []int{(v15 - v15), 243, 32767}
    v11 -= 1
    for fi17 := 0; fi17 < 18; fi17++ {
        v3 = v3
        v18 := GenS1{F0: (((0.0 - (((1.5 * 0.0) / 1.0) * 1e10)) != -2.25) || (128 != v3)), F1: false, F2: fi17}
        var v19 float64 = 1.5
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS2{F0: 42, F1: uint16(v3), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
