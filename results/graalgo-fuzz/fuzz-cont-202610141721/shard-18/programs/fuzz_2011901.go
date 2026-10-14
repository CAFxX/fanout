package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 string; F1 int32; F2 int16 }
func genFn0(p0 int16, p1 uint16, p2 int8) uint64 {
    var v4 uint = 4294967295
    var v5 float64 = ((-2.25 - (((0.1 * (1e10 * 1e10)) + 1e10) + 1.5)) - ((float64(3.14159) + -2.25) - 1e10))
    v6 := GenS1{F0: 573, F1: 242, F2: (p0 != p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int = -591
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    v7 := GenS1{F0: 1, F1: v1, F2: true}
    var v8 int32 = 42
    v9 := GenS2{F0: v8}
    fmt.Println(v8, v8)
    var v10 float64 = 1.5
    v8 = v8
    if false {
        var v11 float32 = ((-2.25 * float32((v10 - (v10 - v10)))) - 1.5)
        _ = v11
    }
    var v12 int = v1
    var v13 bool = false
    fmt.Println(v1, v1, v1)
    for fi14 := 0; fi14 < 9; fi14++ {
        v10 -= 1.5
        var v15 uint16 = 581
        v16 := GenS3{F0: 4294967295, F1: 2}
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS4{F0: "foo bar", F1: (v8 ^ (^(v8))), F2: 741}
    fmt.Println(v10, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
