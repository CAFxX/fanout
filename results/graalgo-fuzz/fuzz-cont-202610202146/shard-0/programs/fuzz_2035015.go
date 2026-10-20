package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := []uint32{2147483647, 566}
    var v2 int8 = 67
    v3 := make(map[int]string)
    var v4 int64 = 468
    var v5 float32 = 0.1
    var v6 uint16 = 32767
    v7 := v3[int(v4)]
    v8 := GenS1{F0: -307}
    var v9 float32 = v5
    var v10 float32 = (0.0 - float32((1e10 - 1e10)))
    v11 := GenS1{F0: 2147483647}
    if false {
        var v12 float64 = float64(float32(((-2.25 - 1.5) + 0.0)))
        var v13 string = "\n\t"
        var v14 uint32 = 2
        _ = v12
        _ = v13
        _ = v14
    }
    if false {
        var v15 bool = true
        fmt.Println(v5, v4, v6)
        _ = v15
    }
    v5 = (v5 + (v10 / v10))
    if len(v1) > 0 {
        v1[0] = 690
    }
    switch "\n\t" {
    case "hello":
        fmt.Println(v7, v4)
    case "":
        v16 := GenS1{F0: -765}
        _ = v16
    case "a":
        v17 := GenS1{F0: -1}
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
