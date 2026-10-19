package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint64; F2 uint }

func main() {
    v1 := make(map[uint32]int)
    var v2 string = ("\n\t" + "")
    var v3 int = 127
    v4 := make(map[string]string)
    var v5 bool = true
    v6 := []int{(v3 ^ (v3 / ((v3) | 1))), v3}
    v2 = ("a" + v2)
    var v7 int64 = -100
    v2 = v2
    v7 = -1
    var v8 string = "foo bar"
    if false {
        fmt.Println(v7, v8, v5)
    }
    v9 := GenS1{F0: v8}
    v10 := GenS2{F0: -32768, F1: 4294967295, F2: 42}
    delete(v1, 2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
