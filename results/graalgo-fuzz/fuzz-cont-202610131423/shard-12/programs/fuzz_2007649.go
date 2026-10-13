package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint32 }
func genFn0(p0 uint64) uint64 {
    var v6 int = 4294967295
    v6 = -273
    if (false && (297 < -892)) {
        v6 -= (v6 + v6)
    }
    v6 = (^(v6))
    var v7 bool = true
    _ = p0
    _ = v6
    _ = v7
    return 457
}

func main() {
    var v1 uint8 = 42
    var v2 int8 = -128
    var v3 string = "foo bar"
    v4 := make(map[string]uint32)
    var v5 float64 = ((3.14159 * 3.14159) / 3.14159)
    v8 := []int{-9223372036854775808, 32767, -711, -9223372036854775808}
    v9 := GenS1{F0: -9223372036854775808}
    switch v3 {
    case "a":
        _ = len(v8)
    case "":
        var v10 float64 = float64(3.14159)
        _ = v10
    }
    v2 -= v2
    v2 += v2
    var v11 uint16 = 429
    v12 := GenS2{F0: v11, F1: (v11 | v11), F2: 0}
    v13 := GenS2{F0: v11, F1: v11, F2: 679}
    v14 := GenS1{F0: -9223372036854775808}
    v15 := GenS1{F0: -45}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
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
