package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 int = 127
    var v2 uint8 = 105
    var v3 float64 = -2.25
    v4 := GenS1{F0: true}
    v5 := []int{357, 128}
    var v6 string = ("a" + "a")
    var v7 string = v6
    var v8 uint16 = 268
    fmt.Println(v2, v8, v3)
    if len(v5) > 0 {
        v5[0] = (v1 | v1)
    }
    v6 = v6
    v5 = append(v5, 4294967295)
    v2 -= v2
    _ = len(v5)
    v8 += 32767
    v9 := GenS2{F0: v8, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
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
