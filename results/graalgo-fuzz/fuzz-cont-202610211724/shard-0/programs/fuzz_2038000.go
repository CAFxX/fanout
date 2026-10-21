package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int64 }

func main() {
    var v1 int16 = 0
    var v2 float32 = 3.14159
    v3 := make(map[string]string)
    v4 := v3["a"]
    switch v4 {
    case "hello":
        fmt.Println(v2)
    default:
        v5 := GenS1{F0: 373}
        _ = v5
    }
    var v6 int = 0
    delete(v3, v4)
    var v7 uint8 = 1
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := []int{fi8, (fi8 >> ((1) & 63)), v6, (v6 ^ v6)}
        var v10 int = v6
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(v4, v4, v2)
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 uint64 = 1
        _ = fi11
        _ = v12
    }
    var v13 uint16 = 93
    v14 := []int{}
    for fi15 := 0; fi15 < 7; fi15++ {
        v2 -= ((1e10 / v2) - ((v2 / v2) * v2))
        var v16 uint16 = v13
        _ = fi15
        _ = v16
    }
    var v17 int32 = 0
    v6 += v6
    v7 += (^(v7))
    var v18 uint32 = 0
    var v19 int32 = (v17 | 42)
    v20 := GenS2{F0: 2, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
