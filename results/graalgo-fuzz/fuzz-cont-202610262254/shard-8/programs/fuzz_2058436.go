package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint16 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 1
    v3 := make(map[string]uint32)
    var v4 uint16 = 2
    var v5 uint8 = 17
    v6 := make(map[string]bool)
    var v7 int = -9223372036854775808
    var v8 uint32 = 2
    var v9 float64 = 0.0
    v10 := GenS1{F0: "a", F1: 627, F2: (v2 - (v4 << ((4294967295) & 63)))}
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := []int{fi11}
        _ = fi11
        _ = v12
    }
    fmt.Println(v8)
    v7 -= (v7 + v7)
    v13 := GenS2{F0: ""}
    v9 = float64(float32(0.0))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
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
