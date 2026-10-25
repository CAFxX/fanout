package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 42
    v3 := make(map[uint32]bool)
    var v4 bool = true
    v5 := GenS1{F0: 2}
    delete(v3, 4294967295)
    if (v4 || v4) {
        v2 = 2
    } else {
        v2 += (v2 * 4294967295)
        v6 := []int{654, 0, 128, 128}
        _ = v6
    }
    v4 = true
    v7 := GenS1{F0: 255}
    var v8 float64 = 3.14159
    var v9 uint = 0
    v2 += (v2 % ((v2) | 1))
    v2 = 969
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
