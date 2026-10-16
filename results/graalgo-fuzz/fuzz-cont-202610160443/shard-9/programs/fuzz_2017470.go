package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []int{-933}
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    var v5 int64 = 644
    if false {
        v5 = 0
    }
    if (0 != 438) {
        fmt.Println(v5)
    } else {
        v6 := GenS1{F0: true}
        v5 += -601
        _ = v6
    }
    v5 -= (v5 << ((370) & 63))
    var v7 int32 = -2147483648
    v7 += v7
    v1 = append(v1, 405)
    fmt.Println(v7)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
