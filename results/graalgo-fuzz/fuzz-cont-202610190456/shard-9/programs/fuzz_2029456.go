package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = float64(0.0)
    v3 := []uint32{2147483647, 42}
    v4 := make(map[string]int)
    v5 := make(map[int]bool)
    if (!true) {
        var v6 int64 = 406
        v7 := v4["foo bar"]
        var v8 int64 = (v6 | v6)
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: -25}
    var v10 bool = (!false)
    var v11 int32 = -1
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    v12 := GenS1{F0: 128}
    var v13 float32 = (0.1 + (0.0 / -2.25))
    if true {
        v14 := GenS2{F0: (0 >= 42), F1: 206}
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 int = (-(fi15))
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: false, F1: 133}
    var v18 float64 = (v2 / (v2 - float64(float32(v2))))
    v19 := GenS1{F0: 2147483647}
    var v20 int32 = (-(v11))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
