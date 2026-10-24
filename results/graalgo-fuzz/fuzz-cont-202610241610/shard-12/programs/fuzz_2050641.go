package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]int)
    var v3 int8 = 42
    var v4 int = 404
    v5 := make(map[string]bool)
    var v6 int = -9223372036854775808
    v6 = v6
    var v7 float32 = 1e10
    _ = len(v5)
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint = 2
        _ = fi8
        _ = v9
    }
    fmt.Println(v7)
    var v10 string = "hello"
    fmt.Println(v10, v3)
    v6 = (v6 + v4)
    var v11 bool = (v4 < v6)
    v7 = 1e10
    delete(v5, "")
    v12 := GenS1{F0: 98, F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
