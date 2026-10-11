package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int = -519
    var v2 float64 = 0.1
    var v3 string = ("\n\t" + "a")
    fmt.Println(v1, v3, v2)
    var v4 int8 = 100
    fmt.Println(v3, v4, v2)
    v5 := []int{-48, v1, (v1 << ((0) & 63))}
    switch (v1 * 1) {
    case 4294967295:
        v3 = v3
    case 293:
        fmt.Println(v4, v4)
    default:
        var v6 uint16 = 42
        _ = v6
    }
    _ = len(v5)
    for fi7 := 0; fi7 < 6; fi7++ {
        v3 = v3
        var v8 uint8 = 2
        _ = fi7
        _ = v8
    }
    if (1 != 42) {
        fmt.Println(v3, v4)
        var v9 float64 = (1e10 - (v2 + float64(1.5)))
        v10 := GenS1{F0: 128}
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 128}
    v12 := GenS1{F0: 128}
    v2 -= 1.5
    _ = len(v5)
    var v13 float32 = (float32(v2) * 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
