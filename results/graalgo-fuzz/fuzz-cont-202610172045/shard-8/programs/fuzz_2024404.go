package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint8; F2 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 bool = true
    var v3 int16 = -32768
    v4 := make(map[uint32]string)
    switch "" {
    case "":
        _ = len(v1)
    }
    v3 += v3
    if ((v3 < (-(v3))) || v2) {
        v2 = false
        v2 = v2
    }
    v5 := []int{32768, 2147483647}
    v6 := GenS1{F0: (-2147483648 >= -687)}
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int8 = -12
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: 42, F1: 70, F2: 193}
    if (v2 || ((v3 / ((v3) | 1)) > v3)) {
        var v10 int32 = -1
        var v11 uint8 = 127
        v11 -= 92
        _ = v10
        _ = v11
    }
    var v12 uint32 = 42
    for fi13 := 0; fi13 < 19; fi13++ {
        fmt.Println(v3)
        _ = fi13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
