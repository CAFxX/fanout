package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := []uint32{65535, 4294967295, 0}
    v2 := make(map[string]int)
    var v3 int64 = -1
    var v4 float32 = ((0.1 + (3.14159 + 0.0)) - 1.5)
    var v5 uint32 = 490
    v4 -= v4
    v6 := GenS1{F0: 100}
    var v7 uint8 = 128
    if false {
        switch -9223372036854775808 {
        case 42:
            var v8 float64 = (-2.25 + 1e10)
            _ = v8
        default:
            v9 := GenS1{F0: 127}
            _ = v9
        }
        var v10 uint64 = 652
        _ = v10
    } else {
        v4 = 1.5
        var v11 uint64 = 1
        _ = v11
    }
    v4 -= v4
    for fi12 := 0; fi12 < 3; fi12++ {
        v7 = 2
        _ = fi12
    }
    fmt.Println(v4)
    v3 = 0
    v4 += v4
    var v13 uint64 = 2
    var v14 uint32 = v5
    fmt.Println(v4, v14, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
