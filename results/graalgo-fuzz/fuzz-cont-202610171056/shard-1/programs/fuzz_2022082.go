package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint64 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    var v4 float64 = 1.5
    v5 := make(map[string]string)
    v6 := make(map[uint32]uint32)
    var v7 int = 32768
    v1 -= (v1 >> ((65535) & 63))
    fmt.Println(v4, v7)
    for fi8 := 0; fi8 < 11; fi8++ {
        v4 += 1.5
        v9 := v6[410]
        fmt.Println(v4)
        _ = fi8
        _ = v9
    }
    if true {
        fmt.Println(v4, v7, v7)
        v10 := GenS1{F0: (v7 % (((v7 << ((42) & 63))) | 1)), F1: v1, F2: 4294967295}
        _ = v10
    } else {
        var v11 uint = 18446744073709551615
        _ = v11
    }
    v12 := GenS1{F0: (v7 << ((255) & 63)), F1: v1, F2: 1}
    for fi13 := 0; fi13 < 17; fi13++ {
        fmt.Println(v7, fi13)
        _ = fi13
    }
    var v14 float64 = ((0.0 / 3.14159) * 1e10)
    fmt.Println(v7)
    v15 := GenS1{F0: (v7 * (v7 % (((v7 % ((v7) | 1))) | 1))), F1: v1, F2: 18446744073709551615}
    v1 -= -845
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(v7)
        fmt.Println(v7, v1)
        v17 := GenS1{F0: v7, F1: (-(v1)), F2: 2}
        _ = fi16
        _ = v17
    }
    v4 = float64(float32(v4))
    v14 -= v4
    v18 := GenS2{F0: -1}
    var v19 int32 = 42
    v7 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
