package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int8 = -28
    var v2 uint16 = 951
    v3 := make(map[int]string)
    var v4 bool = false
    switch "a" {
    case "":
        v5 := []int{}
        _ = v5
    case "foo bar":
        fmt.Println(v1)
    }
    v2 += (^(v2))
    var v6 int32 = 792
    v7 := GenS1{F0: 0}
    var v8 int64 = 42
    v9 := v3[4294967295]
    fmt.Println(v2, v8)
    v10 := []int{896, 32768}
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
