package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 32767
    var v3 float64 = (((-2.25 * float64(0.1)) - 1e10) * 0.1)
    var v4 float64 = ((float64((0.0 * (0.1 + 0.1))) + (0.0 + float64((((0.0 / 1.5) * (float32(0.0) / 1e10)) / -2.25)))) / v3)
    v5 := []int{32768, int(v2)}
    if len(v5) > 0 {
        v5[0] = 0
    }
    v4 = 0.0
    fmt.Println(v2, v2, v3)
    var v6 float64 = v3
    var v7 float32 = float32(v6)
    var v8 int64 = 0
    for fi9 := 0; fi9 < 20; fi9++ {
        fmt.Println(v4)
        _ = fi9
    }
    switch "" {
    case "":
        if true {
            var v10 uint16 = 355
            _ = v10
        } else {
            var v11 float32 = v7
            v11 -= 1e10
            _ = v11
        }
    case "\n\t":
        v12 := GenS1{F0: 2147483647, F1: 1}
        _ = v12
    }
    var v13 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
