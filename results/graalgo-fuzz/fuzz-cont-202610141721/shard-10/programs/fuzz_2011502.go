package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]int)
    var v3 uint = 255
    var v4 int8 = 0
    var v5 uint = v3
    delete(v2, "")
    v3 = (v5 % ((v3) | 1))
    v6 := []int{-1}
    var v7 uint8 = 189
    var v8 float32 = 1.5
    var v9 float64 = (1e10 * (3.14159 + 3.14159))
    v7 += 128
    var v10 float64 = v9
    if (!(!(778 <= 42))) {
        fmt.Println(v9)
        v6 = append(v6, 2147483647)
    }
    v11 := GenS1{F0: 42}
    var v12 uint8 = v7
    v13 := GenS1{F0: 0}
    v3 += v3
    var v14 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
