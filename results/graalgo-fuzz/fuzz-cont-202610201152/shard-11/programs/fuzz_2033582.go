package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint = 59
    var v2 int64 = 1
    var v3 float32 = ((0.0 - ((3.14159 - 1.5) * ((1.5 + -2.25) + float32((3.14159 - 0.1))))) / 0.1)
    var v4 float64 = float64((float32(1e10) + v3))
    var v5 uint32 = 65535
    v6 := []uint32{(v5 & 0), (v5 / ((65535) | 1))}
    if (!(!true)) {
        v2 += v2
    }
    v5 = (v5 * v5)
    v7 := GenS1{F0: -51}
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(v5)
        v2 = -1
        _ = fi8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        v3 = 3.14159
        _ = fi9
    }
    var v10 uint16 = 65535
    fmt.Println(v1, v3, v1)
    var v11 int16 = 0
    var v12 int64 = (-(v2))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
