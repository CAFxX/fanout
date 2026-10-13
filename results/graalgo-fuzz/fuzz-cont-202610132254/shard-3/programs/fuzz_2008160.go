package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 83
    var v3 int = 1
    var v4 int32 = -1
    v5 := make(map[uint32]bool)
    var v6 bool = false
    v7 := GenS1{F0: 0}
    var v8 int32 = (-(v4))
    var v9 int8 = (v2 >> ((865) & 63))
    v2 -= (v2 ^ v2)
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 int16 = -32768
        var v12 int8 = v2
        var v13 int32 = (v8 >> ((18446744073709551615) & 63))
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v2 = -109
    v14 := GenS1{F0: 18446744073709551615}
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint64 = 4294967295
        v17 := []int{-1}
        var v18 float64 = (0.0 / (3.14159 - (0.0 - 0.1)))
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
