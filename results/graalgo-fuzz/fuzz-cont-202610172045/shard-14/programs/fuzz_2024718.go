package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := []string{"foo bar", "a"}
    v3 := []int{}
    var v4 bool = false
    v5 := []int{-2147483648, 42}
    var v6 int32 = 883
    var v7 float64 = (3.14159 / 1.5)
    fmt.Println(v4)
    var v8 float32 = ((1.5 * 3.14159) - 3.14159)
    v9 := GenS1{F0: v6}
    v8 -= 0.1
    fmt.Println(v4, v7, v6)
    v6 = v6
    v4 = ((172 >= 4294967295) && v4)
    v7 = ((v7 / ((v7 + (v7 / v7)) - (((0.1 - v7) / 1.5) / 0.1))) * v7)
    v7 += 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
