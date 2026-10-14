package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int = -585
    var v2 int16 = 1
    var v3 string = "foo bar"
    var v4 float32 = (0.1 - 3.14159)
    v5 := GenS1{F0: 42}
    fmt.Println(v2, v3)
    var v6 int64 = 42
    v3 = "a"
    v7 := []int{}
    v4 = v4
    var v8 float32 = v4
    v7 = append(v7, v1)
    for fi9 := 0; fi9 < 15; fi9++ {
        v4 -= (v4 + (0.1 - 1e10))
        fmt.Println(v1, v1)
        if false {
            fmt.Println(v2)
            v6 += v6
        }
        _ = fi9
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
