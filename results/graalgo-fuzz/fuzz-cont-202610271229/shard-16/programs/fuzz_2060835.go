package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]int)
    var v3 float64 = (1.5 + 3.14159)
    v4 := make(map[string]string)
    var v5 float32 = float32(v3)
    var v6 uint64 = 1
    fmt.Println(v5, v5, v6)
    var v7 uint16 = 2
    var v8 int64 = -9223372036854775808
    var v9 int8 = 127
    v10 := GenS1{F0: 2, F1: true}
    var v11 int8 = (v9 - v9)
    fmt.Println(v1, v6)
    var v12 int = 9223372036854775807
    v7 -= v7
    var v13 string = "foo bar"
    var v14 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
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
