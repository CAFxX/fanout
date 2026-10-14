package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }

func main() {
    var v1 int32 = 1
    v2 := []uint32{}
    v3 := make(map[uint32]bool)
    var v4 uint32 = 1
    if false {
        v2 = append(v2, v4)
    } else {
        v1 -= v1
        v5 := GenS1{F0: 42}
        _ = v5
    }
    var v6 uint32 = v4
    v7 := GenS2{F0: v6}
    var v8 uint8 = 183
    delete(v3, 2147483647)
    v9 := GenS3{F0: "a"}
    fmt.Println(v6, v6, v1)
    v10 := GenS1{F0: 340}
    if len(v2) > 0 {
        v2[0] = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
