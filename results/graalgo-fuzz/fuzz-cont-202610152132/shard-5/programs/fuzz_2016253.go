package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int64 = 994
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]uint32)
    var v4 int8 = 42
    v1 += 9223372036854775807
    var v5 int16 = 0
    v1 -= v1
    v5 -= v5
    v6 := GenS1{F0: 42}
    v5 -= 542
    v4 = (v4 % (((v4 | (v4 / ((v4) | 1)))) | 1))
    v7 := GenS2{F0: 42}
    v1 += (v2 ^ v2)
    v8 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
