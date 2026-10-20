package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 string }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int16 = 32767
    var v2 int32 = 1
    var v3 uint64 = 544
    fmt.Println(v2)
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 bool = true
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 2, F1: 42, F2: ""}
    fmt.Println(v2, v2)
    v7 := GenS1{F0: 4294967295, F1: -9223372036854775808, F2: "a"}
    v8 := GenS2{F0: 42, F1: 127}
    v9 := GenS3{F0: 2}
    v2 += v2
    var v10 uint8 = 42
    v11 := []int{376, -592, 2147483647, 9223372036854775807}
    v2 += (v2 / ((-2147483648) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
