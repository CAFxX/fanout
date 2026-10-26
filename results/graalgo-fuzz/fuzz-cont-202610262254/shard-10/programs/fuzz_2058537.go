package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }

func main() {
    var v1 int8 = 127
    var v2 uint32 = 600
    v3 := make(map[uint32]string)
    var v4 uint16 = 861
    v5 := []int{265, 127}
    var v6 float32 = 3.14159
    var v7 int32 = 42
    var v8 uint32 = v2
    v9 := GenS1{F0: 2, F1: -4}
    var v10 uint8 = 153
    fmt.Println(v6)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
