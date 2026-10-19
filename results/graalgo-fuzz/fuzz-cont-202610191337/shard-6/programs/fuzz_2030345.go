package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }

func main() {
    var v1 float32 = ((((0.0 + float32(0.0)) + 1e10) - ((1.5 * 1.5) / 1.0)) - 0.1)
    var v2 string = "\n\t"
    var v3 bool = true
    v4 := GenS1{F0: true}
    var v5 int32 = -2147483648
    v2 = "foo bar"
    v6 := GenS2{F0: -750}
    v3 = v3
    v7 := []int{9223372036854775807, -2147483648, 4294967295, 2147483647}
    var v8 string = v2
    v9 := GenS2{F0: 0}
    v10 := GenS1{F0: v3}
    var v11 int8 = -1
    v8 = "hello"
    switch v8 {
    case "\n\t":
        v8 = v2
    default:
        v12 := GenS1{F0: (-1 > 1)}
        _ = v12
    }
    var v13 int32 = -779
    fmt.Println(v3)
    v11 -= 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
