package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int32) bool {
    var v6 int16 = -32768
    p0 = (p0 + (p0 ^ p0))
    _ = p0
    _ = v6
    return false
}

func main() {
    var v1 int16 = 461
    var v2 int64 = -1
    var v3 int16 = (v1 * 1000)
    var v4 int32 = -2147483648
    var v5 uint32 = 2147483647
    fmt.Println(v1)
    v7 := GenS1{F0: -9223372036854775808, F1: 42}
    v8 := GenS2{F0: true, F1: "a"}
    var v9 int = 1
    v10 := []int{-1, (v9 << ((2) & 63))}
    v11 := GenS3{F0: 0}
    v12 := GenS2{F0: ((2 <= 255) && true), F1: ""}
    var v13 uint16 = 1
    var v14 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
