package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }

func main() {
    var v1 string = ("foo bar" + "hello")
    var v2 int16 = 42
    var v3 float32 = float32(0.0)
    var v4 uint = 42
    v3 += (v3 * v3)
    var v5 uint32 = 0
    var v6 uint64 = 42
    v5 = v5
    fmt.Println(v3, v1, v4)
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: -128, F1: v2}
        var v9 uint32 = (^(v5))
        var v10 int8 = int8(fi7)
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    if true {
        var v11 uint8 = 76
        v12 := []int{-2147483648}
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: 100, F1: (v2 & -32768)}
    v6 = uint64(v5)
    var v14 uint64 = (v6 ^ (^(v6)))
    var v15 int64 = -591
    var v16 uint = (v4 - uint(v14))
    var v17 int8 = -1
    v18 := []int{-989}
    v14 -= (v14 ^ v14)
    v15 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
