package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }
type GenS2 struct { F0 uint; F1 uint64 }

func main() {
    var v1 int16 = 32767
    var v2 int16 = (-(v1))
    v3 := make(map[int]string)
    var v4 uint8 = 70
    v5 := GenS1{F0: 18446744073709551615, F1: -177, F2: ""}
    v6 := GenS2{F0: 27, F1: 864}
    v1 = (-(v1))
    v1 = v2
    v2 -= v1
    var v7 string = "foo bar"
    var v8 uint8 = (v4 | v4)
    v9 := []int{-1, 32768, 9223372036854775807}
    if false {
        fmt.Println(v7)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
