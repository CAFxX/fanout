package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int16 }
func genFn0(p0 uint16) uint {
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := []int{9223372036854775807}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 0, F1: 767, F2: 32767}
    _ = p0
    _ = v8
    return 42
}

func main() {
    v1 := []uint32{4294967295, 344, 2147483647, 435}
    var v2 int64 = 0
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    var v5 bool = false
    fmt.Println(v5)
    v5 = v5
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := v3[2]
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        switch "hello" {
        case "\n\t":
            var v13 float64 = 1.5
            _ = v13
        }
        _ = fi12
    }
    var v14 uint64 = 0
    v14 += v14
    fmt.Println(v2, v14)
    var v15 uint16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
