package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v4 uint64 = 99
    v5 := v2[42]
    if (3.14159 != 0.0) {
        var v6 bool = true
        var v7 uint16 = 65535
        v8 := GenS1{F0: 22, F1: 13}
        _ = v6
        _ = v7
        _ = v8
    }
    v4 = v4
    var v9 uint16 = 0
    v4 -= 88
    var v10 float32 = (0.0 + ((1e10 * 3.14159) - ((3.14159 * float32(((3.14159 / float64(((float32(((3.14159 * (float64((1.5 * 0.1)) * (0.0 * (3.14159 / 1e10)))) + 0.1)) / 3.14159) - 3.14159))) + 1e10))) / 1.5)))
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 float32 = v10
        _ = fi11
        _ = v12
    }
    var v13 bool = (((3.14159 * ((v10 - 0.1) * v10)) < (v10 - 0.0)) || (756 != 2147483647))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
