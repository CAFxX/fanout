package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 string }

func main() {
    var v1 int8 = 1
    v2 := make(map[string]uint32)
    var v3 string = "\n\t"
    v4 := []string{"hello", ("foo bar" + "")}
    var v5 float32 = ((((1.5 * -2.25) * 0.1) / 1.0) * (0.0 + 1.5))
    v6 := GenS1{F0: (^(v1)), F1: 2, F2: v3}
    var v7 bool = true
    var v8 int32 = 0
    var v9 int = 127
    var v10 int64 = 0
    v9 -= (^(v9))
    v1 = v1
    v3 = "hello"
    v4 = append(v4, ((v3 + v3) + ((("" + v3) + "foo bar") + v3)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
