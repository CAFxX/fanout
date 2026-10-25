package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 string }

func main() {
    var v1 float64 = float64((((1.5 / 0.1) + -2.25) + 1e10))
    var v2 int16 = 32767
    var v3 uint8 = 0
    v4 := []int{-319, 804}
    v5 := GenS1{F0: -895, F1: 65535, F2: (("a" + "hello") + ((("\n\t" + "hello") + ("a" + "hello")) + "foo bar"))}
    v2 = (-(v2))
    var v6 int8 = 1
    v7 := GenS1{F0: 1, F1: 2147483648, F2: "foo bar"}
    var v8 float32 = 0.0
    v2 = 1000
    fmt.Println(v8, v1, v2)
    var v9 uint8 = (v3 & (v3 / ((128) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
