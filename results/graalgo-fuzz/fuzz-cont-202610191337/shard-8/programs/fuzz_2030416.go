package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 bool, p1 int16, p2 uint) uint32 {
    v4 := GenS1{F0: 1, F1: 0}
    v5 := GenS2{F0: -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2147483647
}

func main() {
    var v1 string = "foo bar"
    var v2 int64 = -259
    var v3 int32 = 2147483647
    var v6 uint32 = 4294967295
    v1 = "\n\t"
    v2 += (v2 | (v2 * v2))
    var v7 uint64 = 1
    fmt.Println(v6, v1, v3)
    fmt.Println(v2, v2, v7)
    var v8 float32 = 1e10
    v9 := []int{-2147483648}
    v1 = "foo bar"
    v9 = append(v9, -9223372036854775808)
    v10 := GenS2{F0: 541}
    v1 = (v1 + (v1 + ((("a" + v1) + "hello") + "\n\t")))
    var v11 uint = 65535
    if false {
        v11 -= uint(v6)
        v11 = (v11 & v11)
    }
    v2 -= v2
    v9 = append(v9, 32768)
    var v12 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
