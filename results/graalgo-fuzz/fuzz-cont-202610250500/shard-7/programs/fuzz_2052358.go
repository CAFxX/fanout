package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 bool }
type GenS2 struct { F0 int64 }
func genFn0(p0 int64) int64 {
    var v5 uint64 = 0
    v6 := []int{127}
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: 1, F1: 1, F2: true}
        v5 += v5
        _ = fi7
        _ = v8
    }
    var v9 uint = 18446744073709551615
    p0 += p0
    _ = p0
    _ = v5
    _ = v6
    _ = v9
    return p0
}
func genFn1(p0 int16, p1 int16) uint8 {
    p0 += 0
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := []int{fi10, (fi10 + (^(fi10))), -165, fi10}
        p1 += p0
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    return 233
}

func main() {
    var v1 int64 = -1
    v2 := make(map[int]bool)
    var v3 string = ""
    var v4 float64 = 0.1
    switch 2147483647 {
    case 128:
        v3 = v3
    }
    v1 = (v1 / ((v1) | 1))
    var v12 int32 = -1
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := GenS2{F0: v1}
        var v15 uint = 386
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        v17 := []int{}
        var v18 uint = 42
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 int32 = v12
        _ = fi19
        _ = v20
    }
    v3 = v3
    for fi21 := 0; fi21 < 9; fi21++ {
        v1 = (-(v1))
        var v22 int64 = v1
        _ = fi21
        _ = v22
    }
    v12 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
