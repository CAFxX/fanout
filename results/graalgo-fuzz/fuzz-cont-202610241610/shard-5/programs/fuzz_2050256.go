package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint32 = 2147483647
    var v2 uint32 = 107
    var v3 bool = (((0.0 + -2.25) / 1.0) > 1.5)
    v4 := make(map[uint32]int)
    var v5 int64 = -313
    var v6 int = int(v2)
    fmt.Println(v1)
    v1 = 1
    v7 := GenS1{F0: 322}
    _ = len(v4)
    var v8 uint16 = 600
    v9 := GenS1{F0: (v8 | v8)}
    v8 = (v8 / (((^(v8))) | 1))
    var v10 float32 = (-2.25 * 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
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
