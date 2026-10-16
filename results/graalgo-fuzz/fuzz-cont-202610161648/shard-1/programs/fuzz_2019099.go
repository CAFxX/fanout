package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }

func main() {
    var v1 bool = (675 >= 42)
    var v2 int32 = 2147483647
    v3 := make(map[string]string)
    var v4 float32 = (3.14159 - 1e10)
    v5 := make(map[int]bool)
    var v6 string = ("" + "hello")
    var v7 string = ((v6 + v6) + "\n\t")
    var v8 int16 = 0
    fmt.Println(v8, v1, v4)
    v9 := GenS1{F0: 528, F1: v8}
    if v1 {
        v7 = (v7 + "foo bar")
        v10 := v3[((v6 + (v7 + "")) + v7)]
        var v11 int64 = 581
        _ = v10
        _ = v11
    } else {
        var v12 float64 = -2.25
        v13 := GenS1{F0: 467, F1: (v8 | v8)}
        v1 = (v2 <= v2)
        _ = v12
        _ = v13
    }
    var v14 int64 = -9223372036854775808
    fmt.Println(v14, v6, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
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
