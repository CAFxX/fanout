package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint64 = v1
    var v3 int = 252
    var v4 float32 = 0.1
    v5 := make(map[string]bool)
    var v6 int = v3
    var v7 float32 = (3.14159 * (v4 / v4))
    v1 += (v2 | (v1 >> ((2) & 63)))
    var v8 int64 = 1
    v7 = v7
    var v9 int = (v3 * v3)
    v4 += (3.14159 - (v7 + v4))
    if (true && true) {
        if false {
            v10 := GenS1{F0: 42}
            _ = v10
        }
        var v11 uint64 = 42
        _ = v11
    }
    var v12 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
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
