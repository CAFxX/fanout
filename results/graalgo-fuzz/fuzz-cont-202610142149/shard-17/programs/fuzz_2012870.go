package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int64 }
type GenS2 struct { F0 string; F1 int16; F2 uint64 }

func main() {
    var v1 int64 = 0
    var v2 float32 = 1.5
    var v3 uint16 = 0
    var v4 uint64 = 18446744073709551615
    fmt.Println(v1, v4, v3)
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: v1, F1: false, F2: v1}
        _ = fi5
        _ = v6
    }
    fmt.Println(v4)
    v7 := GenS1{F0: -118, F1: false, F2: (-(v1))}
    var v8 float32 = 3.14159
    var v9 uint8 = 228
    fmt.Println(v9, v3)
    switch 0 {
    case -9223372036854775808:
        v10 := GenS2{F0: ((("\n\t" + "a") + "hello") + "\n\t"), F1: 0, F2: v4}
        _ = v10
    }
    v11 := GenS1{F0: 776, F1: false, F2: (v1 | int64(v3))}
    v12 := []int{-9223372036854775808, -9223372036854775808}
    fmt.Println(v3, v3)
    v13 := GenS2{F0: "", F1: 1, F2: (v4 + (^(v4)))}
    var v14 int8 = 31
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
