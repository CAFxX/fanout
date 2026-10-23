package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int16 }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]string)
    v3 := []int{128, 450, 128}
    var v4 int16 = 0
    v5 := make(map[uint32]bool)
    v4 -= 0
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint32 = 65535
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 9223372036854775807}
    v1 += 1e10
    v4 = (v4 ^ v4)
    var v9 int = 0
    var v10 int8 = 0
    v11 := GenS2{F0: v10, F1: (v4 >> ((255) & 63))}
    var v12 int = (v9 & v9)
    var v13 uint32 = 769
    if (false || (!true)) {
        _ = len(v2)
    } else {
        v5[2147483648] = true
        v10 -= v10
    }
    var v14 int32 = 1
    var v15 int = (v12 << ((18446744073709551615) & 63))
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 float32 = 1.5
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
