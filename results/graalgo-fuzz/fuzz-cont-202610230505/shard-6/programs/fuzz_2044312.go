package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 bool }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    v5 := []int{0}
    var v6 float32 = 0.1
    v7 := GenS1{F0: 42, F1: 127, F2: (!true)}
    var v8 int64 = 42
    v1 += v1
    v9 := v2[""]
    var v10 uint = 1
    v6 = v1
    if ((-364 >= -32768) || false) {
        var v11 bool = true
        fmt.Println(v8, v11, v1)
        _ = v11
    }
    var v12 bool = (!("foo bar" == "a"))
    v6 -= float32(0.0)
    _ = len(v5)
    fmt.Println(v12, v12, v12)
    fmt.Println(v1, v10, v6)
    fmt.Println(v1)
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 string = "hello"
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
