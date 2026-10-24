package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]bool)
    var v3 int = 2147483647
    var v4 float64 = 1e10
    var v5 uint32 = 365
    for fi6 := 0; fi6 < 1; fi6++ {
        v5 = v5
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 float32 = (0.0 * 3.14159)
        v2[v5] = (!false)
        v9 := GenS1{F0: 2}
        _ = fi7
        _ = v8
        _ = v9
    }
    v3 -= 127
    var v10 float32 = ((1.5 * 3.14159) - 0.0)
    var v11 int64 = 42
    v12 := GenS1{F0: 42}
    fmt.Println(v11)
    fmt.Println(v11, v11)
    fmt.Println(v11, v10, v3)
    v3 -= v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
