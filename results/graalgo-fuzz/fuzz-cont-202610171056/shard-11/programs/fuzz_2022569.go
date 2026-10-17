package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }

func main() {
    var v1 uint32 = 2
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    v4 := []uint32{v1}
    v5 := make(map[uint32]bool)
    v6 := make(map[int]uint32)
    var v7 int = 42
    fmt.Println(v7, v7)
    _ = len(v4)
    var v8 int8 = -128
    for fi9 := 0; fi9 < 11; fi9++ {
        if false {
            v4 = append(v4, v1)
            if len(v4) > 0 {
                v4[0] = v1
            }
            v8 = (v8 + v8)
        }
        v10 := GenS1{F0: int(v1), F1: 42}
        _ = fi9
        _ = v10
    }
    var v11 uint32 = v1
    fmt.Println(v11)
    var v12 int32 = 0
    var v13 int32 = v12
    v12 = v13
    var v14 uint32 = (v1 * (v11 + (v11 * (^(v11)))))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
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
}
