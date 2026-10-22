package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 int64; F2 int32 }

func main() {
    var v1 float32 = 0.1
    var v2 uint32 = 927
    var v3 bool = true
    var v4 uint64 = 676
    var v5 int16 = 1
    var v6 int8 = 0
    v7 := GenS1{F0: 524, F1: "", F2: (v5 * v5)}
    v8 := []int{127, -2147483648, 42, 0}
    v5 = v5
    var v9 int = 0
    if true {
        var v10 uint8 = 127
        v2 += (v2 ^ 2)
        if v3 {
            fmt.Println(v1, v9, v10)
        }
        _ = v10
    }
    v4 -= (^(v4))
    var v11 string = "\n\t"
    v12 := GenS2{F0: (v11 + v11)}
    fmt.Println(v3)
    v13 := GenS3{F0: (v4 - v4), F1: 1, F2: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
