package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = (float32(3.14159) / -2.25)
    var v3 int = -476
    var v4 float32 = 1e10
    var v5 uint32 = 0
    v6 := make(map[uint32]string)
    _ = len(v6)
    v7 := GenS1{F0: 65535, F1: 118}
    var v8 int16 = 0
    var v9 uint64 = 2
    if (true || false) {
        v10 := GenS1{F0: uint(v9), F1: 42}
        v11 := GenS2{F0: (v9 % ((4294967295) | 1))}
        _ = v10
        _ = v11
    }
    fmt.Println(v2)
    var v12 string = "a"
    if false {
        v13 := GenS2{F0: v9}
        v14 := v6[v5]
        var v15 uint = 2
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 float32 = (v4 / float32(-2.25))
    v17 := []int{v3, v3}
    switch v12 {
    case "":
        v3 -= (v3 << ((255) & 63))
    }
    var v18 uint32 = v5
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := v6[(v5 ^ 2147483648)]
        fmt.Println(v4)
        v9 -= 1
        _ = fi19
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
