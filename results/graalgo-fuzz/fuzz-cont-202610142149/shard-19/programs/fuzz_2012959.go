package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int8 }
type GenS2 struct { F0 int16; F1 uint64 }
func genFn0(p0 int64, p1 int16, p2 uint8) uint8 {
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := []int{fi5, fi5, -1}
        _ = fi5
        _ = v6
    }
    p1 = p1
    var v7 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 127
}
func genFn1(p0 bool, p1 int, p2 uint8) int16 {
    var v8 float64 = ((float64(float32(0.0)) - ((float64(0.0) + 1e10) * float64(0.1))) + 1.5)
    v9 := GenS1{F0: -890, F1: p1, F2: 42}
    p0 = (v8 >= v8)
    var v10 float32 = ((0.0 + 3.14159) / (-2.25 * 1e10))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []string{}
    var v2 int8 = 42
    var v3 uint16 = 32767
    v4 := []uint32{558, 464}
    v11 := GenS2{F0: -32768, F1: 0}
    v2 -= v2
    if true {
        var v12 bool = (("\n\t" + (("" + "foo bar") + "hello")) < "\n\t")
        fmt.Println(v2)
        _ = v12
    }
    v2 = (v2 / ((v2) | 1))
    v3 = v3
    var v13 uint32 = 65535
    fmt.Println(v3, v3)
    v13 = 42
    v14 := GenS1{F0: 1000, F1: -1, F2: (v2 / ((v2) | 1))}
    _ = len(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
