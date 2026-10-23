package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int8 }

func main() {
    var v1 bool = false
    var v2 uint64 = 440
    var v3 int8 = -1
    v4 := make(map[uint32]int)
    v5 := make(map[int]bool)
    var v6 int8 = v3
    var v7 int16 = -733
    for fi8 := 0; fi8 < 2; fi8++ {
        v1 = false
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int = (-(fi8))
            v11 := v4[1]
            v12 := GenS1{F0: (-(v7)), F1: (-128 != v3), F2: v3}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    switch 2147483647 {
    case 32767:
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 float64 = 1e10
            _ = fi13
            _ = v14
        }
    case 9223372036854775807:
        v15 := GenS1{F0: (v7 & -32768), F1: true, F2: -18}
        _ = v15
    case 42:
        _ = len(v5)
    }
    var v16 uint = 1
    v2 -= v2
    fmt.Println(v3, v3, v1)
    v3 = int8(v7)
    v6 = v6
    v1 = (1 < 42)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
