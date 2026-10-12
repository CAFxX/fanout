package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int64 }
type GenS2 struct { F0 string; F1 int16 }

func main() {
    var v1 float64 = 3.14159
    var v2 uint64 = 2
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 1
    v6 := make(map[string]bool)
    v1 += v1
    var v7 float64 = 0.0
    v8 := GenS1{F0: v2, F1: 42, F2: 42}
    v9 := []int{-9223372036854775808}
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 uint32 = 178
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: "foo bar", F1: 42}
    var v13 string = "a"
    if false {
        _ = len(v9)
        fmt.Println(v7)
    }
    var v14 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
