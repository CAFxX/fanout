package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint16 }

func main() {
    v1 := []int{}
    var v2 bool = false
    var v3 float64 = 1e10
    var v4 int64 = 1
    var v5 bool = false
    var v6 uint16 = 255
    var v7 int16 = 173
    v8 := GenS1{F0: false, F1: "foo bar", F2: v6}
    var v9 int16 = -1
    v5 = (-2.25 <= (v3 - (0.1 - v3)))
    var v10 int = 32768
    var v11 float32 = float32((v3 / v3))
    fmt.Println(v7, v3, v5)
    v12 := GenS1{F0: (v2 || (v2 && v2)), F1: "a", F2: (v6 - 255)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
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
