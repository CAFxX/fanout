package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []string{"\n\t", "a"}
    v2 := []uint32{2}
    var v3 uint32 = 1
    v4 := []uint32{(^(v3))}
    fmt.Println(v3)
    var v5 int8 = 42
    var v6 float64 = (((1.5 / (3.14159 + -2.25)) / 1e10) / 1.0)
    v7 := GenS1{F0: (!false)}
    v3 += v3
    var v8 int16 = 1000
    var v9 int = 127
    if false {
        var v10 float32 = -2.25
        v6 = 1e10
        _ = v10
    } else {
        v9 -= v9
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
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
