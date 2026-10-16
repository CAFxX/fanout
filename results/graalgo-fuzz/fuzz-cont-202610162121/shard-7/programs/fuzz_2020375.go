package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int8 }
func genFn0(p0 uint16) uint32 {
    var v7 float32 = 1.5
    var v8 uint32 = 197
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 int32 = -2147483648
        var v11 int = fi9
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(p0)
    v8 -= (v8 - (^(v8)))
    _ = p0
    _ = v7
    _ = v8
    return v8
}

func main() {
    v1 := []string{"a", "a"}
    v2 := []string{}
    v3 := make(map[string]uint32)
    var v4 uint64 = 4294967295
    var v5 uint8 = 255
    v6 := make(map[int]uint32)
    _ = len(v2)
    v5 = (v5 & v5)
    fmt.Println(v5, v4)
    var v12 int64 = 42
    v13 := GenS1{F0: -128, F1: -1, F2: 34}
    fmt.Println(v5, v5, v4)
    v5 += (v5 + (v5 / ((v5) | 1)))
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 bool = (-2147483648 >= -163)
        v12 = 203
        _ = fi14
        _ = v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
