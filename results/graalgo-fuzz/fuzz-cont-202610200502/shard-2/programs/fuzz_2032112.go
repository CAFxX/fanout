package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 string = ("" + "\n\t")
    var v2 int16 = 0
    v3 := make(map[int]bool)
    v3[128] = true
    v4 := GenS1{F0: 42}
    v1 = ((v1 + v1) + v1)
    var v5 uint = 2
    var v6 int8 = 42
    v5 = v5
    var v7 float32 = 1.5
    var v8 string = v1
    v7 += ((-2.25 - (-2.25 * (float32((-2.25 - 3.14159)) / float32(1e10)))) / ((1.5 + 0.1) + (v7 * v7)))
    v1 = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
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
}
