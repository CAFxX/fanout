package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 bool }
type GenS2 struct { F0 int }

func main() {
    var v1 float32 = 3.14159
    v2 := []int{2147483647, 1}
    v3 := []int{-628, -1}
    v4 := make(map[string]uint32)
    v5 := make(map[string]bool)
    v1 += (float32((0.0 * (1e10 - 0.1))) - 0.1)
    delete(v4, "hello")
    v1 += 1e10
    v1 = 1e10
    _ = len(v5)
    var v6 float32 = (v1 + (v1 + (-2.25 + ((((v1 / v1) / (0.0 / ((float32(3.14159) + (3.14159 / 1e10)) + v1))) * v1) - float32((-2.25 * -2.25))))))
    v7 := GenS1{F0: -143, F1: false, F2: (!false)}
    switch 1 {
    case 42:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v6 += v6
    switch 127 {
    case 128:
        var v8 int32 = 395
        _ = v8
    }
    v6 = 1e10
    v9 := GenS2{F0: 797}
    v2 = append(v2, 344)
    v1 = (1.5 * float32(3.14159))
    for fi10 := 0; fi10 < 5; fi10++ {
        fmt.Println(v6, fi10, fi10)
        _ = fi10
    }
    _ = len(v3)
    fmt.Println(v6, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
