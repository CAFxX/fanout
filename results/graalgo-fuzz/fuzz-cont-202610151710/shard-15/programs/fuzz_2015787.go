package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []int{9223372036854775807, 32767, -782}
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    var v4 int = -667
    var v5 bool = false
    var v6 int64 = 483
    var v7 int8 = 78
    v8 := GenS1{F0: 148, F1: 332}
    var v9 bool = (1 != 0)
    _ = len(v1)
    fmt.Println(v6, v7)
    v1 = append(v1, (-(v4)))
    var v10 float64 = (1.5 / 1e10)
    fmt.Println(v10, v5, v4)
    var v11 float32 = 3.14159
    var v12 uint = 65535
    delete(v3, 566)
    v4 -= v4
    v11 = v11
    v13 := GenS2{F0: 1}
    v14 := GenS2{F0: -571}
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
