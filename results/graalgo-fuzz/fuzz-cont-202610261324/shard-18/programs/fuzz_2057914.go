package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint8; F1 int32 }
func genFn0() uint {
    var v4 float64 = 1e10
    var v5 uint32 = 42
    v6 := GenS1{F0: 0}
    v7 := []int{1}
    var v8 uint32 = (v5 * v5)
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 bool, p1 uint8, p2 int16) int16 {
    fmt.Println(p0)
    var v9 int8 = -128
    if p0 {
        v10 := GenS1{F0: -1}
        _ = v10
    }
    v11 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v9
    return p2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 301
    var v3 uint8 = 0
    fmt.Println(v2, v2)
    for fi12 := 0; fi12 < 19; fi12++ {
        if (!true) {
            var v13 bool = false
            _ = len(v1)
            v14 := GenS2{F0: -521, F1: uint16(v2), F2: 965}
            _ = v13
            _ = v14
        }
        _ = fi12
    }
    var v15 int = -1
    fmt.Println(v2, v2)
    v3 += uint8(v2)
    v16 := GenS1{F0: (v15 | (v15 ^ v15))}
    v3 += (v3 + v3)
    var v17 float64 = -2.25
    var v18 float64 = v17
    var v19 int = (-(v15))
    v20 := GenS1{F0: int(v2)}
    var v21 float64 = 0.1
    for fi22 := 0; fi22 < 19; fi22++ {
        v23 := []int{(v15 + 127), 42}
        var v24 bool = (1.5 == 1e10)
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        var v26 string = "foo bar"
        delete(v1, (^(v2)))
        _ = fi25
        _ = v26
    }
    v27 := GenS2{F0: -966, F1: 0, F2: 967}
    if (148 != 0) {
        v28 := GenS3{F0: v3, F1: 1}
        v19 = (-(v15))
        v2 += (v2 | (v2 - v2))
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
}
