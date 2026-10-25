package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int32 = -954
    var v2 uint32 = 92
    v3 := make(map[int]bool)
    var v4 int64 = -574
    v5 := GenS1{F0: 2}
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := []int{}
        v2 -= v2
        _ = fi6
        _ = v7
    }
    v8 := []int{492, 757, 455, 127}
    var v9 uint32 = v2
    v10 := GenS2{F0: -97}
    fmt.Println(v1, v4)
    var v11 string = "hello"
    v12 := GenS1{F0: 645}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
