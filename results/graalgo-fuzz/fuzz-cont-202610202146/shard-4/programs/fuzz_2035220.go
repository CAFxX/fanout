package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int8 }

func main() {
    var v1 uint8 = 195
    var v2 int8 = -1
    var v3 int16 = 1000
    var v4 int8 = 42
    v5 := make(map[uint32]string)
    v2 = (v4 * v4)
    _ = len(v5)
    v3 -= int16(v1)
    var v6 int16 = (v3 ^ -1)
    v2 = (^(v2))
    var v7 float32 = 0.1
    v8 := GenS1{F0: "a", F1: (!(!false)), F2: (^(v2))}
    fmt.Println(v4, v2, v3)
    v9 := GenS1{F0: "\n\t", F1: false, F2: (v2 * (^(v4)))}
    var v10 uint = 42
    v11 := []int{4294967295, 127, 1}
    var v12 int8 = 100
    v6 += (v6 / ((v6) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
