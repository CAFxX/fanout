package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint16 }

func main() {
    var v1 string = ""
    v2 := []int{}
    v3 := make(map[uint32]uint32)
    v4 := []uint32{}
    var v5 uint64 = 298
    var v6 int16 = -1
    switch -809 {
    case -2147483648:
        v2 = append(v2, -2147483648)
    }
    _ = len(v2)
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int = (fi7 - fi7)
        v5 = 1
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(v6, v1, v1)
        v6 = 42
        _ = fi9
    }
    _ = len(v4)
    var v10 float64 = 3.14159
    v10 += 1.5
    v11 := GenS1{F0: -53, F1: (v5 ^ 4294967295), F2: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
