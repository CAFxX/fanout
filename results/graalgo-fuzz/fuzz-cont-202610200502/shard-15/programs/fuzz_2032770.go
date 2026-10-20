package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 1
    v3 := make(map[int]uint32)
    var v4 float32 = -2.25
    v4 = 1.5
    v2 = v2
    for fi5 := 0; fi5 < 18; fi5++ {
        fmt.Println(v2, fi5)
        _ = fi5
    }
    var v6 uint = 935
    var v7 uint32 = 25
    var v8 uint32 = 0
    v9 := v1[(^(v7))]
    fmt.Println(v9)
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 float32 = float32(0.0)
        fmt.Println(v4, v9, v7)
        fmt.Println(v11, fi10, v7)
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 uint64 = (^(v2))
        v14 := []int{fi12, (fi12 ^ fi12)}
        v15 := GenS1{F0: 9223372036854775807, F1: 42, F2: 702}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    fmt.Println(v9, v4, v6)
    _ = len(v1)
    var v16 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
