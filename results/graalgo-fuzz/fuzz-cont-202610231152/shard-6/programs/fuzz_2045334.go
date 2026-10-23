package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint64; F2 uint32 }
func genFn0(p0 uint32) int8 {
    var v4 bool = false
    v5 := GenS1{F0: -415}
    v6 := []int{-943, 9223372036854775807, 32767}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 127
}

func main() {
    v1 := []int{1, -1, 32768}
    v2 := make(map[string]uint32)
    var v3 int16 = 0
    fmt.Println(v3)
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 float32 = ((0.1 - 0.1) + 1e10)
        _ = fi7
        _ = v8
    }
    v3 = 1
    v1 = append(v1, 32768)
    fmt.Println(v3)
    var v9 int8 = 94
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := GenS1{F0: fi10}
        _ = fi10
        _ = v11
    }
    _ = len(v1)
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 float32 = ((1e10 - 3.14159) * 3.14159)
        var v14 uint64 = 1
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 string = "a"
    v9 = (v9 % ((v9) | 1))
    v16 := GenS2{F0: v9, F1: 238, F2: 2147483648}
    fmt.Println(v15, v3, v15)
    var v17 int16 = (^(v3))
    v2[v15] = 988
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
