package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 float64 = (1.5 * (((float64(((0.1 + ((1e10 - 1e10) / 1.5)) + (3.14159 - 1.5))) + 0.0) + float64(0.1)) + 1e10))
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    var v4 float32 = 0.1
    v5 := make(map[int]bool)
    var v6 uint8 = 1
    v7 := GenS1{F0: 334}
    var v8 int32 = 2147483647
    v9 := v5[2147483647]
    switch "\n\t" {
    case "foo bar":
        var v10 float64 = (v1 + 1.5)
        _ = v10
    default:
        v11 := []int{-2147483648}
        _ = v11
    }
    fmt.Println(v1, v8)
    var v12 int64 = -195
    fmt.Println(v9, v6, v6)
    delete(v2, -1)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
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
