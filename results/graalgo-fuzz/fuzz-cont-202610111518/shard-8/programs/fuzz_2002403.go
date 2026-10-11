package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }

func main() {
    var v1 bool = (-2147483648 != 0)
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 uint16 = 32767
    var v5 uint32 = 2147483647
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 bool = v1
        fmt.Println(v1, v4, v4)
        _ = fi6
        _ = v7
    }
    var v8 uint32 = (v5 + 0)
    var v9 uint16 = v4
    v10 := GenS1{F0: 42, F1: -128}
    var v11 float32 = 0.1
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 uint32 = v8
        fmt.Println(v11, v5, v1)
        _ = fi12
        _ = v13
    }
    fmt.Println(v1)
    for fi14 := 0; fi14 < 11; fi14++ {
        fmt.Println(fi14)
        fmt.Println(v11, v11)
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
