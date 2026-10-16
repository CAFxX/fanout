package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 float32 = 0.0
    var v2 uint = 42
    var v3 float64 = (-2.25 * 0.0)
    var v4 string = "\n\t"
    var v5 int16 = 0
    v4 = v4
    fmt.Println(v4)
    v6 := GenS1{F0: 1}
    var v7 int = 1
    fmt.Println(v2)
    for fi8 := 0; fi8 < 20; fi8++ {
        v5 = (v5 ^ v5)
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    v7 += v7
    var v10 bool = true
    v11 := []int{(v7 + v7), v7, -312, 128}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
