package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := []uint32{2}
    var v2 int8 = 0
    v3 := make(map[string]bool)
    var v4 float64 = float64(1.5)
    v5 := GenS1{F0: 150, F1: 628}
    v2 = 100
    v4 = v4
    v4 += v4
    var v6 bool = (2 == 42)
    var v7 uint8 = 127
    var v8 int = 32767
    var v9 string = ("" + (("" + ("" + ("hello" + "\n\t"))) + "\n\t"))
    var v10 uint8 = 42
    switch (v9 + ((v9 + v9) + ("" + v9))) {
    case "":
        v4 = (1e10 * v4)
    default:
        var v11 uint64 = 4294967295
        _ = v11
    }
    fmt.Println(v8, v9, v7)
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int32 = 1
        for fi14 := 0; fi14 < 17; fi14++ {
            v8 = (fi14 & int(v10))
            v4 -= (3.14159 + float64(float32(v4)))
            _ = fi14
        }
        v15 := GenS1{F0: v7, F1: 1}
        _ = fi12
        _ = v13
        _ = v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
