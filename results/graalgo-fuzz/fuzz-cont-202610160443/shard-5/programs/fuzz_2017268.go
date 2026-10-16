package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint; F2 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = (0.0 / float32(0.1))
    var v3 bool = true
    var v4 uint16 = 32767
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := GenS1{F0: v3, F1: 18446744073709551615, F2: 32767}
        v2 = v2
        _ = fi5
        _ = v6
    }
    var v7 uint64 = 2
    var v8 uint = 42
    var v9 uint16 = (v4 & 798)
    v7 = 18446744073709551615
    v10 := v1[0]
    v2 = (v2 * v2)
    v11 := GenS2{F0: 42, F1: (v8 / ((42) | 1)), F2: (v10 ^ int(v9))}
    fmt.Println(v10, v8, v7)
    v12 := GenS1{F0: true, F1: 117, F2: -1}
    var v13 int32 = 2147483647
    var v14 uint64 = v7
    v7 = v14
    v2 -= v2
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 uint16 = (^(v4))
        _ = fi15
        _ = v16
    }
    var v17 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
