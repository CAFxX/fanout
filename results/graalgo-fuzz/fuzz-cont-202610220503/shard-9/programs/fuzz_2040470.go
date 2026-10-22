package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int64 = 42
    var v2 float32 = 1e10
    v3 := make(map[uint32]bool)
    var v4 float32 = 3.14159
    var v5 uint16 = 255
    v6 := GenS1{F0: 1}
    v7 := []int{0, -1}
    if false {
        fmt.Println(v5, v5)
        var v8 uint = 65535
        v5 = v5
        _ = v8
    }
    if len(v7) > 0 {
        v7[0] = -61
    }
    _ = len(v3)
    _ = len(v3)
    v9 := GenS2{F0: 1}
    var v10 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
