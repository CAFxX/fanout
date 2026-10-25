package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 bool = (!true)
    v2 := make(map[uint32]int)
    var v3 bool = v1
    var v4 float32 = ((0.0 + 0.1) - 0.0)
    v4 = 0.1
    var v5 bool = v3
    v5 = (-1 >= -1)
    v6 := GenS1{F0: 0}
    var v7 float64 = (3.14159 * (3.14159 - 1e10))
    v5 = v1
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int16 = 32767
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: -1}
    var v11 uint = 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
