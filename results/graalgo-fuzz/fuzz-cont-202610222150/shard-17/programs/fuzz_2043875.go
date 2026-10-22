package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int16; F1 uint64; F2 string }
func genFn0(p0 uint, p1 string, p2 uint8) uint {
    for fi6 := 0; fi6 < 17; fi6++ {
        p1 = p1
        var v7 int64 = 42
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 42, F1: 42}
    v9 := GenS2{F0: -32768, F1: 0, F2: p1}
    var v10 bool = true
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return p0
}
func genFn1(p0 uint32, p1 int64) bool {
    var v11 uint64 = 1
    v12 := []int{253, 42, 32767}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return true
}

func main() {
    var v1 float64 = (1e10 - (0.0 - 0.0))
    v2 := make(map[string]string)
    var v3 int8 = -1
    var v4 int8 = v3
    v5 := []uint32{2147483648, 2147483647, 334}
    v3 -= (v3 ^ v3)
    var v13 int32 = 42
    v13 += (v13 % ((v13) | 1))
    var v14 int8 = (v3 % ((v4) | 1))
    if true {
        var v15 float32 = 1e10
        var v16 int = -615
        _ = v15
        _ = v16
    } else {
        v17 := GenS1{F0: 0, F1: v3}
        fmt.Println(v3)
        var v18 int32 = (^(v13))
        _ = v17
        _ = v18
    }
    v2[("\n\t" + "\n\t")] = "hello"
    v4 = 127
    var v19 int16 = -1
    var v20 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
