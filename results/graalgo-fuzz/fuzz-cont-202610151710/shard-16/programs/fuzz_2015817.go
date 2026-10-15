package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int8 }

func main() {
    var v1 int = 32767
    v2 := make(map[int]uint32)
    var v3 float32 = float32(1e10)
    v4 := make(map[int]string)
    v5 := []int{}
    delete(v4, v1)
    _ = len(v5)
    v3 -= 1.5
    v6 := GenS1{F0: -1, F1: v1, F2: -72}
    var v7 float64 = -2.25
    v8 := v2[v1]
    v3 = v3
    fmt.Println(v1)
    v2[v1] = v8
    var v9 int16 = -1
    v10 := GenS1{F0: v9, F1: (v1 >> ((452) & 63)), F2: 115}
    if len(v5) > 0 {
        v5[0] = (v1 ^ v1)
    }
    var v11 bool = false
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(v1)
        _ = fi12
    }
    v13 := GenS1{F0: (v9 << ((4294967295) & 63)), F1: v1, F2: -35}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
