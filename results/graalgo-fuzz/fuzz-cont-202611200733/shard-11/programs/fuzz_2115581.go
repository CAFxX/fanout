package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 147
    var v3 float32 = ((-2.25 + 0.0) * (1.5 / (0.1 * 1.5)))
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 int32 = -203
        v2 = v2
        _ = fi4
        _ = v5
    }
    var v6 uint32 = 2
    fmt.Println(v3, v6, v6)
    var v7 int64 = -531
    var v8 int = 727
    v9 := GenS1{F0: -444}
    v8 -= v8
    v10 := v1["\n\t"]
    var v11 uint = 0
    var v12 float64 = -2.25
    v2 += uint16(v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
