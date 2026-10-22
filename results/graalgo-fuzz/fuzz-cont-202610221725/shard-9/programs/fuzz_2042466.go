package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
func genFn0(p0 int16, p1 string, p2 uint16) uint32 {
    var v7 float32 = float32(-2.25)
    var v8 uint16 = (p2 >> ((255) & 63))
    v9 := GenS1{F0: (p0 >= (p0 | -539)), F1: 25}
    v10 := GenS1{F0: false, F1: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}
func genFn1() uint32 {
    var v11 uint8 = 193
    for fi12 := 0; fi12 < 16; fi12++ {
        v11 = v11
        fmt.Println(fi12, fi12)
        _ = fi12
    }
    _ = v11
    return 649
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]uint32)
    var v3 int16 = 1000
    v4 := make(map[string]int)
    v5 := []uint32{}
    var v6 int8 = 42
    v13 := v4[("\n\t" + "hello")]
    var v14 bool = v1
    var v15 bool = (!v1)
    delete(v2, 0)
    var v16 uint8 = 128
    for fi17 := 0; fi17 < 14; fi17++ {
        v3 -= v3
        v18 := v2[560]
        var v19 float64 = (0.1 + 3.14159)
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 int16 = (v3 & (v3 - 1000))
    var v21 int8 = (v6 / (((^(v6))) | 1))
    for fi22 := 0; fi22 < 3; fi22++ {
        var v23 int = (-(v13))
        _ = fi22
        _ = v23
    }
    delete(v2, 1)
    fmt.Println(v1, v15)
    v24 := GenS1{F0: ((2 == 0) && true), F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
