package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
func genFn0(p0 uint64, p1 uint16) int8 {
    var v5 uint16 = p1
    p0 += (p0 * p0)
    var v6 int32 = 1
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: fi7, F1: 0}
        v9 := []int{32767}
        var v10 int = fi7
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 100
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = (float32((1.5 * 0.0)) + 1.5)
    v3 := make(map[int]int)
    var v4 uint8 = 2
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 int8 = 100
        _ = fi11
        _ = v12
    }
    _ = len(v1)
    var v13 float32 = (v2 * v2)
    if true {
        v14 := []int{-9223372036854775808, 103, 127}
        v15 := GenS1{F0: 2147483647, F1: 705}
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 93, F1: 29}
    _ = len(v3)
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 uint32 = 65535
        var v19 uint64 = 0
        var v20 int16 = 1
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    v2 -= float32((1e10 - -2.25))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
