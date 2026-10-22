package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int64 }

func main() {
    var v1 float32 = float32(1e10)
    var v2 int = -1
    v3 := make(map[int]string)
    var v4 int32 = 838
    var v5 int64 = 400
    switch ("\n\t" + "\n\t") {
    case "hello":
        var v6 uint = 18446744073709551615
        _ = v6
    case "a":
        fmt.Println(v5, v1, v5)
    default:
        var v7 uint = 0
        _ = v7
    }
    v5 = 42
    v4 = v4
    v8 := GenS1{F0: "a", F1: 1, F2: -1}
    fmt.Println(v1, v2, v5)
    v9 := []int{128}
    var v10 uint = 2
    v11 := GenS1{F0: "a", F1: 0, F2: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
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
