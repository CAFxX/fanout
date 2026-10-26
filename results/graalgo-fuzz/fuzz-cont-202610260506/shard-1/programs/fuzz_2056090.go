package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }

func main() {
    v1 := []int{}
    var v2 int16 = 1
    var v3 float64 = 0.0
    var v4 uint16 = 0
    v5 := make(map[uint32]int)
    var v6 float32 = -2.25
    var v7 int32 = 0
    var v8 float64 = 0.0
    if len(v1) > 0 {
        v1[0] = int(v7)
    }
    v3 = 3.14159
    v6 = 1.5
    for fi9 := 0; fi9 < 1; fi9++ {
        v8 = (v8 - 3.14159)
        v10 := GenS1{F0: -913, F1: v4}
        _ = fi9
        _ = v10
    }
    v3 += float64(v6)
    var v11 uint8 = 255
    var v12 int8 = 28
    for fi13 := 0; fi13 < 16; fi13++ {
        if (false && (v3 >= v8)) {
            v8 += v8
        }
        for fi14 := 0; fi14 < 20; fi14++ {
            if false {
                v15 := GenS1{F0: 1, F1: (v4 + 2)}
                _ = v15
            }
            v6 += (3.14159 / 3.14159)
            _ = fi14
        }
        var v16 float64 = (float64(v6) / 1e10)
        _ = fi13
        _ = v16
    }
    switch "a" {
    case "\n\t":
        v3 += (v3 - v3)
    }
    var v17 int16 = v2
    if true {
        var v18 int64 = -9223372036854775808
        _ = v18
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        _ = len(v1)
        _ = fi19
    }
    var v20 uint32 = 1
    v3 += (((((v3 - 0.0) + (v8 + v3)) * v3) / (3.14159 * v3)) / (float64(float32((1e10 / 3.14159))) / 0.1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
