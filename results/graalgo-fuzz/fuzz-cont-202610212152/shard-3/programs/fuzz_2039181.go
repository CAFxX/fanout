package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 string; F2 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 bool) int8 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int8 = -128
        v6 -= (v6 >> ((255) & 63))
        p0 = (p0 || (!(282 <= 145)))
        _ = fi5
        _ = v6
    }
    _ = p0
    return -123
}

func main() {
    var v1 uint = 37
    var v2 float64 = ((0.1 / 1.5) + 1e10)
    v3 := make(map[uint32]int)
    var v4 bool = false
    v2 -= v2
    v7 := GenS1{F0: 1}
    var v8 uint32 = 42
    var v9 uint64 = 42
    v10 := GenS2{F0: v4, F1: "a", F2: v8}
    delete(v3, v8)
    v11 := GenS3{F0: 127}
    for fi12 := 0; fi12 < 2; fi12++ {
        fmt.Println(v4, fi12, v4)
        delete(v3, (^(v8)))
        _ = fi12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
