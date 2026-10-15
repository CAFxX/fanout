package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int16; F1 int64; F2 int8 }
func genFn0(p0 uint) int64 {
    fmt.Println(p0)
    v5 := []int{32768}
    v5 = append(v5, 0)
    _ = p0
    _ = v5
    return 42
}
func genFn1(p0 bool, p1 int16, p2 int64) uint16 {
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: (fi6 >> ((4294967295) & 63)), F1: 255}
        v8 := GenS1{F0: fi6, F1: 2}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: p1, F1: p2, F2: 100}
    for fi10 := 0; fi10 < 9; fi10++ {
        fmt.Println(p1, p0)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 237
}

func main() {
    v1 := []uint32{65535, 677, 0, 42}
    v2 := []uint32{42}
    var v3 int = 0
    var v4 uint16 = 255
    var v11 int8 = 42
    fmt.Println(v3, v3)
    var v12 uint8 = 42
    if true {
        for fi13 := 0; fi13 < 10; fi13++ {
            v2 = append(v2, 1)
            _ = fi13
        }
        v14 := GenS2{F0: 42, F1: -333, F2: (v11 + (v11 * (v11 | v11)))}
        v4 = v4
        _ = v14
    } else {
        var v15 uint16 = (^(v4))
        v11 = v11
        var v16 uint32 = 1
        _ = v15
        _ = v16
    }
    if false {
        v17 := GenS2{F0: 42, F1: 42, F2: v11}
        _ = v17
    } else {
        switch (v3 + 32767) {
        case -552:
            for fi18 := 0; fi18 < 9; fi18++ {
                switch (v3 * (v3 - 756)) {
                case 4294967295:
                    v11 += (-(v11))
                case 127:
                    var v19 bool = false
                    _ = v19
                case 128:
                    var v20 float32 = 0.1
                    _ = v20
                default:
                    v1 = append(v1, 862)
                }
                fmt.Println(v3)
                _ = fi18
            }
        }
    }
    if true {
        var v21 int16 = 42
        _ = v21
    }
    v3 += v3
    _ = len(v2)
    for fi22 := 0; fi22 < 13; fi22++ {
        fmt.Println(v11, v4)
        fmt.Println(v12, fi22, fi22)
        _ = fi22
    }
    v11 -= (v11 >> ((418) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
