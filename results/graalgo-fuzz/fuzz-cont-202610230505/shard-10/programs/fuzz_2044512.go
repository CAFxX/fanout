package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint16 }

func main() {
    var v1 string = ""
    v2 := make(map[string]string)
    v3 := []int{127, 256}
    var v4 float64 = 1.5
    v5 := []uint32{2, 1}
    v6 := make(map[int]int)
    v4 = (((v4 * (v4 / (float64((-2.25 / (((((3.14159 - (1e10 - ((-2.25 / 1.5) + 1e10))) + 0.1) - 0.1) + -2.25) + -2.25))) * v4))) / float64(0.1)) - 0.1)
    v7 := GenS1{F0: 0, F1: 2, F2: 65535}
    if len(v5) > 0 {
        v5[0] = 42
    }
    var v8 uint16 = 65535
    var v9 bool = (0 != 2)
    v4 = float64(float32((v4 / v4)))
    if v9 {
        var v10 uint32 = 65535
        var v11 int = 127
        v1 = ("hello" + "\n\t")
        _ = v10
        _ = v11
    } else {
        var v12 bool = (126 != 42)
        var v13 string = v1
        _ = v12
        _ = v13
    }
    _ = len(v5)
    fmt.Println(v8, v9)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
