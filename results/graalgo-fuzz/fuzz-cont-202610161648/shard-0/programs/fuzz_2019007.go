package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint16 = 730
    var v2 uint = 502
    var v3 uint = v2
    var v4 float32 = 0.0
    if false {
        var v5 uint64 = 18446744073709551615
        _ = v5
    }
    var v6 int = -9223372036854775808
    var v7 bool = false
    v8 := []int{v6, v6, v6}
    v4 -= v4
    v9 := GenS1{F0: (v1 * (^(v1)))}
    v4 = (v4 / v4)
    fmt.Println(v3, v2)
    for fi10 := 0; fi10 < 10; fi10++ {
        v7 = (2 <= 0)
        _ = fi10
    }
    v11 := GenS1{F0: (v1 & v1)}
    v12 := GenS1{F0: (v1 / ((v1) | 1))}
    var v13 bool = true
    fmt.Println(v3, v6)
    v6 += v6
    v14 := GenS1{F0: (v1 | v1)}
    v1 = v1
    if len(v8) > 0 {
        v8[0] = 257
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
