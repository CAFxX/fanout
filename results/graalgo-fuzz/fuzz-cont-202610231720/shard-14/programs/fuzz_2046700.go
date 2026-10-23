package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int; F2 int16 }
func genFn0(p0 uint64) string {
    v5 := GenS1{F0: 1}
    var v6 int32 = 42
    _ = p0
    _ = v5
    _ = v6
    return "foo bar"
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 128
    var v3 uint64 = 903
    v4 := make(map[uint32]bool)
    v7 := GenS1{F0: 1}
    var v8 int = -416
    v9 := GenS2{F0: 42, F1: (v8 ^ v8), F2: 42}
    v3 = 855
    v8 -= v8
    fmt.Println(v3, v2, v3)
    fmt.Println(v8, v2, v8)
    if len(v1) > 0 {
        v1[0] = 42
    }
    v3 = (v3 ^ 4294967295)
    var v10 string = "\n\t"
    var v11 uint8 = v2
    _ = len(v4)
    v1 = append(v1, 0)
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 int8 = -128
        _ = fi12
        _ = v13
    }
    delete(v4, 42)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
