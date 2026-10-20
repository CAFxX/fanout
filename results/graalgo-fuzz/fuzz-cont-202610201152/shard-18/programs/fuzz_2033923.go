package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 string = "hello"
    var v2 int32 = 2147483647
    var v3 int32 = (v2 ^ (v2 * v2))
    var v4 float64 = 1e10
    var v5 int8 = 100
    switch v1 {
    case "hello":
        v4 = 1.5
    }
    var v6 uint8 = 210
    v7 := GenS1{F0: 4294967295, F1: 810}
    v8 := GenS2{F0: 0}
    var v9 uint64 = 4294967295
    v10 := GenS3{F0: 1000}
    fmt.Println(v2, v9)
    fmt.Println(v3, v9, v1)
    var v11 float32 = (1e10 / (1.5 + float32(1e10)))
    fmt.Println(v3, v3)
    for fi12 := 0; fi12 < 2; fi12++ {
        v4 = (v4 - 1e10)
        _ = fi12
    }
    v6 -= v6
    var v13 int32 = v3
    var v14 int16 = 665
    v1 = v1
    v9 = 1
    v15 := []int{}
    v3 = (v3 << ((18446744073709551615) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
